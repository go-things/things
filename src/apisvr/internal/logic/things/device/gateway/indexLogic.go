package gateway

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/apisvr/internal/logic/things/device"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.DeviceGateWayIndexReq) (resp *types.DeviceGateWayIndexResp, err error) {
	dmReq := &dm.DeviceGatewayIndexReq{
		GatewayDeviceName: req.GateWayDeviceName,
		GatewayProductID:  req.GateWayProductID,
		Page:              logic.ToDmPageRpc(req.Page),
	}
	dmResp, err := l.svcCtx.DeviceM.DeviceGatewayIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceInfo req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.DeviceInfo, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		pi := device.DeviceInfoToApi(v)
		pis = append(pis, pi)
	}
	return &types.DeviceGateWayIndexResp{
		Total: dmResp.Total,
		List:  pis,
	}, nil
	return
}
