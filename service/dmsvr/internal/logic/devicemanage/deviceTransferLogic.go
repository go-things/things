package devicemanagelogic

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/stores"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"gorm.io/gorm"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceTransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceTransferLogic {
	return &DeviceTransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	DeviceTransferToUser    = 1
	DeviceTransferToProject = 2
)

func Transfer() {

}

func (l *DeviceTransferLogic) DeviceTransfer(in *dm.DeviceTransferReq) (*dm.Empty, error) {
	uc := ctxs.GetUserCtx(l.ctx)
	if in.SrcProjectID != 0 {
		uc.ProjectID = in.SrcProjectID
		if !uc.IsAdmin {
			if uc.ProjectAuth == nil || uc.ProjectAuth[in.ProjectID] == nil || uc.ProjectAuth[in.ProjectID].AuthType != def.AuthAdmin {
				return nil, errors.Permissions.AddMsg("只有项目管理员才能创建区域")
			}
		}
	}
	diDB := relationDB.NewDeviceInfoRepo(l.ctx)
	var dis []*relationDB.DmDeviceInfo
	if in.Device != nil {
		di, err := diDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{
			ProductID:   in.Device.ProductID,
			DeviceNames: []string{in.Device.DeviceName},
		})
		if err != nil {
			return nil, err
		}
		dis = append(dis, di)
	}
	if len(in.Devices) != 0 {
		di, err := diDB.FindByFilter(l.ctx, relationDB.DeviceFilter{
			Cores: utils.CopySlice[devices.Core](in.Devices),
		}, nil)
		if err != nil {
			return nil, err
		}
		dis = append(dis, di...)
	}
	if len(dis) == 0 {
		return &dm.Empty{}, nil
	}
	for _, di := range dis {
		pi, err := l.svcCtx.ProjectCache.GetData(l.ctx, int64(di.ProjectID))
		if err != nil {
			return nil, err
		}
		if pi.AdminUserID != pi.AdminUserID {
			return nil, errors.Permissions
		}
	}
	var (
		ProjectID  stores.ProjectID
		AreaID     stores.AreaID = def.NotClassified
		AreaIDPath string        = def.NotClassifiedPath
		UserID     int64
	)

	switch in.TransferTo {
	case DeviceTransferToUser:
		dp, err := l.svcCtx.DataM.DataProjectIndex(ctxs.WithAllProject(l.ctx), &sys.DataProjectIndexReq{
			Page: &sys.PageInfo{
				Page: 1,
				Size: 1,
				Orders: []*sys.PageInfo_OrderBy{{
					Field: "createdTime", //第一个一定是默认的
					Sort:  stores.OrderAsc,
				}},
			},
			TargetID:   in.UserID,
			TargetType: "user",
		})
		if err != nil {
			return nil, err
		}
		if len(dp.List) == 0 {
			return nil, errors.NotFind.AddMsg("用户未找到")
		}
		ProjectID = stores.ProjectID(dp.List[0].ProjectID)
		UserID = in.UserID
	case DeviceTransferToProject:
		ProjectID = stores.ProjectID(in.ProjectID)
		if in.AreaID != 0 {
			ai, err := l.svcCtx.AreaCache.GetData(l.ctx, in.AreaID)
			if err != nil {
				return nil, err
			}
			if ai.ProjectID != in.ProjectID {
				return nil, errors.Parameter.AddMsg("项目不对")
			}
			AreaID = stores.AreaID(ai.AreaID)
			AreaIDPath = ai.AreaIDPath
		}
		pi, err := l.svcCtx.ProjectCache.GetData(l.ctx, in.ProjectID)
		if err != nil {
			return nil, err
		}
		UserID = pi.AdminUserID
	}
	if in.IsCleanData == def.True {
		for _, di := range dis {
			err := DeleteDeviceTimeData(l.ctx, l.svcCtx, di.ProductID, di.DeviceName)
			if err != nil {
				return nil, err
			}
		}

	}
	var devs = utils.CopySlice[devices.Core](dis)
	err := stores.GetTenantConn(l.ctx).Transaction(func(tx *gorm.DB) error {
		err := relationDB.NewUserDeviceShareRepo(tx).DeleteByFilter(l.ctx, relationDB.UserDeviceShareFilter{
			Devices: devs,
		})
		if err != nil {
			return err
		}
		if in.IsCleanData == def.True {
			err = relationDB.NewDeviceProfileRepo(tx).DeleteByFilter(ctxs.WithRoot(l.ctx),
				relationDB.DeviceProfileFilter{Devices: devs})
			if err != nil {
				return err
			}
		}
		err = relationDB.NewDeviceInfoRepo(tx).UpdateWithField(ctxs.WithAllProject(l.ctx), relationDB.DeviceFilter{Cores: devs}, map[string]any{
			"project_id":   ProjectID,
			"user_id":      UserID,
			"area_id":      AreaID,
			"area_id_path": AreaIDPath,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, di := range devs {
		err = l.svcCtx.DeviceCache.SetData(l.ctx, *di, nil)
		if err != nil {
			l.Error(err)
		}
	}

	return &dm.Empty{}, nil
}
