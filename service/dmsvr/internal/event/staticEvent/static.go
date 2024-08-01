package staticEvent

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/stores"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/sdk/service/protocol"
	"github.com/i-Things/things/service/dmsvr/internal/logic"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"sync/atomic"
	"time"
)

type StaticHandle struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
	logx.Logger
}

func NewStaticHandle(ctx context.Context, svcCtx *svc.ServiceContext) *StaticHandle {
	return &StaticHandle{
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}

func (l *StaticHandle) Handle() error { //产品品类设备数量统计
	w := sync.WaitGroup{}
	w.Add(3)
	utils.Go(l.ctx, func() {
		defer w.Done()
		err := l.ProductCategoryStatic()
		if err != nil {
			l.Error(err)
		}
	})
	//utils.Go(l.ctx, func() {
	//	err := l.AreaDeviceStatic()
	//	if err != nil {
	//		l.Error(err)
	//	}
	//})
	utils.Go(l.ctx, func() {
		defer w.Done()
		err := l.DeviceExp()
		if err != nil {
			l.Error(err)
		}
	})
	utils.Go(l.ctx, func() {
		defer w.Done()
		err := l.DeviceOnlineFix()
		if err != nil {
			l.Error(err)
		}
	})
	w.Wait()
	return nil
}
func (l *StaticHandle) AreaDeviceStatic() error { //区域下的设备数量统计
	ret, err := l.svcCtx.AreaM.AreaInfoIndex(l.ctx, &sys.AreaInfoIndexReq{})
	if err != nil {
		return err
	}
	var areaPaths []string
	for _, v := range ret.List {
		areaPaths = append(areaPaths, v.AreaIDPath)
	}
	err = logic.FillAreaDeviceCount(l.ctx, l.svcCtx, areaPaths...)
	return err
}

var count atomic.Int64

func (l *StaticHandle) DeviceOnlineFix() error { //设备在线修复
	nc := count.Add(1)
	if nc/2 == 1 { //1小时处理一次
		return nil
	}
	devs, err := relationDB.NewDeviceInfoRepo(l.ctx).FindCoreByFilter(l.ctx, relationDB.DeviceFilter{IsOnline: def.True}, nil)
	if err != nil {
		return err
	}
	devMap, err := protocol.GetActivityDevices(l.ctx)
	if err != nil {
		l.Error(err)
		return err
	}
	var needOnline []devices.Core
	for _, d := range devs {
		if _, ok := devMap[d]; ok {
			continue
		}
		//如果线上没有,但是这里有,需要进行处理
		needOnline = append(needOnline, d)
	}
	if len(needOnline) > 0 {
		l.Infof("DeviceOnlineFix.UpdatesDeviceActivity devs:%v", utils.Fmt(needOnline))
		protocol.UpdatesDeviceActivity(l.ctx, needOnline)
	}
	return nil
}

func (l *StaticHandle) DeviceExp() error { //设备过期处理
	{ //有效期到了之后不启用
		err := relationDB.NewDeviceInfoRepo(l.ctx).UpdateWithField(l.ctx,
			relationDB.DeviceFilter{ExpTime: stores.CmpAnd(stores.CmpLte(time.Now()), stores.CmpIsNull(false))},
			map[string]any{"status": def.DeviceStatusArrearage})
		if err != nil {
			l.Error(err)
		}
	}
	{ //清除设置了过期时间且过期了的分享
		err := relationDB.NewUserDeviceShareRepo(l.ctx).DeleteByFilter(l.ctx, relationDB.UserDeviceShareFilter{
			ExpTime: stores.CmpAnd(stores.CmpLte(time.Now()), stores.CmpIsNull(false)),
		})
		if err != nil {
			l.Error(err)
		}
	}
	return nil
}

func (l *StaticHandle) ProductCategoryStatic() error { //产品品类设备数量统计
	pcDB := relationDB.NewProductCategoryRepo(l.ctx)
	pcs, err := pcDB.FindByFilter(l.ctx, relationDB.ProductCategoryFilter{}, nil)
	if err != nil {
		return err
	}
	for _, pc := range pcs {
		ids := utils.GetIDPath(pc.IDPath)
		total, err := relationDB.NewDeviceInfoRepo(l.ctx).CountByFilter(l.ctx, relationDB.DeviceFilter{ProductCategoryIDs: ids})
		if err != nil {
			l.Error(err)
			continue
		}
		pc.DeviceCount = total
		err = pcDB.Update(l.ctx, pc)
		if err != nil {
			l.Error(err)
			continue
		}
	}
	return nil
}
