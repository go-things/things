package rulelogic

import (
	"context"
	"fmt"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/oss"
	"github.com/i-Things/things/service/dmsvr/dmExport"
	"github.com/i-Things/things/service/udsvr/internal/domain/scene"
	"github.com/i-Things/things/service/udsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/udsvr/internal/svc"
	"github.com/i-Things/things/service/udsvr/pb/ud"

	"github.com/zeromicro/go-zero/core/logx"
)

type SceneInfoCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSceneInfoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SceneInfoCreateLogic {
	return &SceneInfoCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 场景
func (l *SceneInfoCreateLogic) SceneInfoCreate(in *ud.SceneInfo) (*ud.WithID, error) {
	do := ToSceneInfoDo(in)
	if do.Status == 0 {
		do.Status = def.True
	}
	if do.AreaID == 0 {
		return nil, errors.Parameter.AddMsg("areaID必填")
	}
	if do.Tag == "deviceTiming" { //单设备定时
		uc := ctxs.GetUserCtx(l.ctx)
		err := dmExport.AccessPerm(l.ctx, l.svcCtx.DeviceCache, l.svcCtx.UserShareCache, def.AuthReadWrite, devices.Core{
			ProductID:  do.ProductID,
			DeviceName: do.DeviceName,
		}, "deviceTiming")
		if err != nil {
			return nil, err
		}
		di, err := l.svcCtx.DeviceCache.GetData(l.ctx, devices.Core{
			ProductID:  in.ProductID,
			DeviceName: in.DeviceName,
		})
		if err != nil {
			return nil, err
		}
		if uc.ProjectID != di.ProjectID {
			uc.ProjectID = di.ProjectID
		}
	}

	uc := ctxs.GetUserCtx(l.ctx)
	if do.ProjectID == 0 {
		do.ProjectID = ctxs.GetUserCtxNoNil(l.ctx).ProjectID
	}
	if !uc.IsAdmin && do.ProjectID <= def.RootNode {
		return nil, errors.Parameter.AddMsg("只有管理员可以创建全局场景联动")
	}
	err := do.Validate(NewSceneCheckRepo(l.ctx, l.svcCtx, do))
	if err != nil {
		return nil, err
	}
	po := ToSceneInfoPo(do)
	err = relationDB.NewSceneInfoRepo(l.ctx).Insert(l.ctx, po)
	if err != nil {
		return nil, err
	}
	if in.HeadImg != "" && in.IsUpdateHeadImg { //如果填了参数且不等于原来的,说明修改头像,需要处理
		nwePath := oss.GenFilePath(l.ctx, l.svcCtx.Config.Name, oss.BusinessScene, oss.SceneHeadIng, fmt.Sprintf("%d/%s", po.ID, oss.GetFileNameWithPath(in.HeadImg)))
		path, err := l.svcCtx.OssClient.PrivateBucket().CopyFromTempBucket(in.HeadImg, nwePath)
		if err != nil {
			return nil, errors.System.AddDetail(err)
		}
		po.HeadImg = path
		err = relationDB.NewSceneInfoRepo(l.ctx).UpdateHeadImg(l.ctx, po)
		if err != nil {
			l.Error(err)
		}
	}
	if len(do.If.Triggers) == 0 && do.Type == scene.SceneTypeAuto && do.Status == def.True { //立即执行一次
		_, err = NewSceneManuallyTriggerLogic(l.ctx, l.svcCtx).SceneManuallyTrigger(&ud.WithID{Id: po.ID})
		if err != nil {
			return nil, err
		}
	}
	return &ud.WithID{Id: po.ID}, nil
}
