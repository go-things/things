package gateway

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic/things"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMultiDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiDeleteLogic {
	return &MultiDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiDeleteLogic) MultiDelete(req *types.DeviceGateWayMultiDeleteReq) error {
	_, err := l.svcCtx.DeviceM.DeviceGatewayMultiDelete(l.ctx,
		&dm.DeviceGatewayMultiDeleteReq{
			GatewayProductID:  req.GateWayProductID,
			GatewayDeviceName: req.GateWayDeviceName,
			List:              things.ToDmDeviceCoresPb(req.List)})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.MultiDelete MultiDelete req=%v err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
