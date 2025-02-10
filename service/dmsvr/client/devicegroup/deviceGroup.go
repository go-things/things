// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package devicegroup

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AbnormalLogIndexReq               = dm.AbnormalLogIndexReq
	AbnormalLogIndexResp              = dm.AbnormalLogIndexResp
	AbnormalLogInfo                   = dm.AbnormalLogInfo
	ActionRespReq                     = dm.ActionRespReq
	ActionSendReq                     = dm.ActionSendReq
	ActionSendResp                    = dm.ActionSendResp
	CommonSchemaCreateReq             = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq              = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp             = dm.CommonSchemaIndexResp
	CommonSchemaInfo                  = dm.CommonSchemaInfo
	CommonSchemaUpdateReq             = dm.CommonSchemaUpdateReq
	CompareInt64                      = dm.CompareInt64
	CompareString                     = dm.CompareString
	CustomTopic                       = dm.CustomTopic
	DeviceBindTokenInfo               = dm.DeviceBindTokenInfo
	DeviceBindTokenReadReq            = dm.DeviceBindTokenReadReq
	DeviceCore                        = dm.DeviceCore
	DeviceCountInfo                   = dm.DeviceCountInfo
	DeviceCountReq                    = dm.DeviceCountReq
	DeviceCountResp                   = dm.DeviceCountResp
	DeviceError                       = dm.DeviceError
	DeviceGatewayBindDevice           = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq             = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp            = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq       = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiSaveReq         = dm.DeviceGatewayMultiSaveReq
	DeviceGatewaySign                 = dm.DeviceGatewaySign
	DeviceGroupMultiSaveReq           = dm.DeviceGroupMultiSaveReq
	DeviceInfo                        = dm.DeviceInfo
	DeviceInfoBindReq                 = dm.DeviceInfoBindReq
	DeviceInfoCanBindReq              = dm.DeviceInfoCanBindReq
	DeviceInfoCount                   = dm.DeviceInfoCount
	DeviceInfoCountReq                = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq               = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp               = dm.DeviceInfoIndexResp
	DeviceInfoMultiBindReq            = dm.DeviceInfoMultiBindReq
	DeviceInfoMultiBindResp           = dm.DeviceInfoMultiBindResp
	DeviceInfoMultiUpdateReq          = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                 = dm.DeviceInfoReadReq
	DeviceInfoUnbindReq               = dm.DeviceInfoUnbindReq
	DeviceModuleVersion               = dm.DeviceModuleVersion
	DeviceModuleVersionIndexReq       = dm.DeviceModuleVersionIndexReq
	DeviceModuleVersionIndexResp      = dm.DeviceModuleVersionIndexResp
	DeviceModuleVersionReadReq        = dm.DeviceModuleVersionReadReq
	DeviceMoveReq                     = dm.DeviceMoveReq
	DeviceOnlineMultiFix              = dm.DeviceOnlineMultiFix
	DeviceOnlineMultiFixReq           = dm.DeviceOnlineMultiFixReq
	DeviceProfile                     = dm.DeviceProfile
	DeviceProfileIndexReq             = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp            = dm.DeviceProfileIndexResp
	DeviceProfileReadReq              = dm.DeviceProfileReadReq
	DeviceResetReq                    = dm.DeviceResetReq
	DeviceSchema                      = dm.DeviceSchema
	DeviceSchemaIndexReq              = dm.DeviceSchemaIndexReq
	DeviceSchemaIndexResp             = dm.DeviceSchemaIndexResp
	DeviceSchemaMultiCreateReq        = dm.DeviceSchemaMultiCreateReq
	DeviceSchemaMultiDeleteReq        = dm.DeviceSchemaMultiDeleteReq
	DeviceSchemaTslReadReq            = dm.DeviceSchemaTslReadReq
	DeviceSchemaTslReadResp           = dm.DeviceSchemaTslReadResp
	DeviceShareInfo                   = dm.DeviceShareInfo
	DeviceTransferReq                 = dm.DeviceTransferReq
	DeviceTypeCountReq                = dm.DeviceTypeCountReq
	DeviceTypeCountResp               = dm.DeviceTypeCountResp
	EdgeSendReq                       = dm.EdgeSendReq
	EdgeSendResp                      = dm.EdgeSendResp
	Empty                             = dm.Empty
	EventLogIndexReq                  = dm.EventLogIndexReq
	EventLogIndexResp                 = dm.EventLogIndexResp
	EventLogInfo                      = dm.EventLogInfo
	Firmware                          = dm.Firmware
	FirmwareFile                      = dm.FirmwareFile
	FirmwareInfo                      = dm.FirmwareInfo
	FirmwareInfoDeleteReq             = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp            = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq              = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp             = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq               = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp              = dm.FirmwareInfoReadResp
	FirmwareResp                      = dm.FirmwareResp
	GatewayCanBindIndexReq            = dm.GatewayCanBindIndexReq
	GatewayCanBindIndexResp           = dm.GatewayCanBindIndexResp
	GatewayGetFoundReq                = dm.GatewayGetFoundReq
	GatewayNotifyBindSendReq          = dm.GatewayNotifyBindSendReq
	GroupCore                         = dm.GroupCore
	GroupDeviceMultiDeleteReq         = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq           = dm.GroupDeviceMultiSaveReq
	GroupInfo                         = dm.GroupInfo
	GroupInfoCreateReq                = dm.GroupInfoCreateReq
	GroupInfoIndexReq                 = dm.GroupInfoIndexReq
	GroupInfoIndexResp                = dm.GroupInfoIndexResp
	GroupInfoMultiCreateReq           = dm.GroupInfoMultiCreateReq
	GroupInfoReadReq                  = dm.GroupInfoReadReq
	GroupInfoUpdateReq                = dm.GroupInfoUpdateReq
	HubLogIndexReq                    = dm.HubLogIndexReq
	HubLogIndexResp                   = dm.HubLogIndexResp
	HubLogInfo                        = dm.HubLogInfo
	IDPath                            = dm.IDPath
	IDPathWithUpdate                  = dm.IDPathWithUpdate
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceConfirmReq       = dm.OtaFirmwareDeviceConfirmReq
	OtaFirmwareDeviceIndexReq         = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp        = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo             = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq         = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                   = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq           = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp          = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo               = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp               = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                   = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq          = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq           = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp          = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq          = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq            = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp           = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq            = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                 = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                  = dm.OtaJobStaticInfo
	OtaModuleInfo                     = dm.OtaModuleInfo
	OtaModuleInfoIndexReq             = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp            = dm.OtaModuleInfoIndexResp
	PageInfo                          = dm.PageInfo
	PageInfo_OrderBy                  = dm.PageInfo_OrderBy
	Point                             = dm.Point
	ProductCategory                   = dm.ProductCategory
	ProductCategoryIndexReq           = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp          = dm.ProductCategoryIndexResp
	ProductCategorySchemaIndexReq     = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp    = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiSaveReq = dm.ProductCategorySchemaMultiSaveReq
	ProductCustom                     = dm.ProductCustom
	ProductCustomReadReq              = dm.ProductCustomReadReq
	ProductCustomUi                   = dm.ProductCustomUi
	ProductInfo                       = dm.ProductInfo
	ProductInfoDeleteReq              = dm.ProductInfoDeleteReq
	ProductInfoIndexReq               = dm.ProductInfoIndexReq
	ProductInfoIndexResp              = dm.ProductInfoIndexResp
	ProductInfoReadReq                = dm.ProductInfoReadReq
	ProductInitReq                    = dm.ProductInitReq
	ProductRemoteConfig               = dm.ProductRemoteConfig
	ProductSchemaCreateReq            = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq            = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq             = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp            = dm.ProductSchemaIndexResp
	ProductSchemaInfo                 = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq       = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq         = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq           = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp          = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq            = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq       = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp      = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg            = dm.PropertyControlSendMsg
	PropertyControlSendReq            = dm.PropertyControlSendReq
	PropertyControlSendResp           = dm.PropertyControlSendResp
	PropertyGetReportMultiSendReq     = dm.PropertyGetReportMultiSendReq
	PropertyGetReportMultiSendResp    = dm.PropertyGetReportMultiSendResp
	PropertyGetReportSendMsg          = dm.PropertyGetReportSendMsg
	PropertyGetReportSendReq          = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp         = dm.PropertyGetReportSendResp
	PropertyLogIndexReq               = dm.PropertyLogIndexReq
	PropertyLogIndexResp              = dm.PropertyLogIndexResp
	PropertyLogInfo                   = dm.PropertyLogInfo
	PropertyLogLatestIndexReq         = dm.PropertyLogLatestIndexReq
	ProtocolConfigField               = dm.ProtocolConfigField
	ProtocolConfigInfo                = dm.ProtocolConfigInfo
	ProtocolInfo                      = dm.ProtocolInfo
	ProtocolInfoIndexReq              = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp             = dm.ProtocolInfoIndexResp
	ProtocolService                   = dm.ProtocolService
	ProtocolServiceIndexReq           = dm.ProtocolServiceIndexReq
	ProtocolServiceIndexResp          = dm.ProtocolServiceIndexResp
	RemoteConfigCreateReq             = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq              = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp             = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq           = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp          = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq            = dm.RemoteConfigPushAllReq
	RespReadReq                       = dm.RespReadReq
	RootCheckReq                      = dm.RootCheckReq
	SdkLogIndexReq                    = dm.SdkLogIndexReq
	SdkLogIndexResp                   = dm.SdkLogIndexResp
	SdkLogInfo                        = dm.SdkLogInfo
	SendLogIndexReq                   = dm.SendLogIndexReq
	SendLogIndexResp                  = dm.SendLogIndexResp
	SendLogInfo                       = dm.SendLogInfo
	SendMsgReq                        = dm.SendMsgReq
	SendMsgResp                       = dm.SendMsgResp
	SendOption                        = dm.SendOption
	ShadowIndex                       = dm.ShadowIndex
	ShadowIndexResp                   = dm.ShadowIndexResp
	SharePerm                         = dm.SharePerm
	StatusLogIndexReq                 = dm.StatusLogIndexReq
	StatusLogIndexResp                = dm.StatusLogIndexResp
	StatusLogInfo                     = dm.StatusLogInfo
	TimeRange                         = dm.TimeRange
	UserDeviceCollectSave             = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq           = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp          = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo               = dm.UserDeviceShareInfo
	UserDeviceShareMultiAcceptReq     = dm.UserDeviceShareMultiAcceptReq
	UserDeviceShareMultiDeleteReq     = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareMultiInfo          = dm.UserDeviceShareMultiInfo
	UserDeviceShareMultiToken         = dm.UserDeviceShareMultiToken
	UserDeviceShareReadReq            = dm.UserDeviceShareReadReq
	WithID                            = dm.WithID
	WithIDChildren                    = dm.WithIDChildren
	WithIDCode                        = dm.WithIDCode
	WithProfile                       = dm.WithProfile

	DeviceGroup interface {
		// 创建分组
		GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error)
		GroupInfoMultiCreate(ctx context.Context, in *GroupInfoMultiCreateReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取分组信息列表
		GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error)
		// 获取分组信息详情
		GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error)
		// 更新分组
		GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error)
		// 删除分组
		GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 创建分组设备
		GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		// 更新分组设备
		GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		// 删除分组设备
		GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultDeviceGroup struct {
		cli zrpc.Client
	}

	directDeviceGroup struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceGroupServer
	}
)

func NewDeviceGroup(cli zrpc.Client) DeviceGroup {
	return &defaultDeviceGroup{
		cli: cli,
	}
}

func NewDirectDeviceGroup(svcCtx *svc.ServiceContext, svr dm.DeviceGroupServer) DeviceGroup {
	return &directDeviceGroup{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 创建分组
func (m *defaultDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoCreate(ctx, in, opts...)
}

// 创建分组
func (d *directDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.GroupInfoCreate(ctx, in)
}

func (m *defaultDeviceGroup) GroupInfoMultiCreate(ctx context.Context, in *GroupInfoMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoMultiCreate(ctx, in, opts...)
}

func (d *directDeviceGroup) GroupInfoMultiCreate(ctx context.Context, in *GroupInfoMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupInfoMultiCreate(ctx, in)
}

// 获取分组信息列表
func (m *defaultDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoIndex(ctx, in, opts...)
}

// 获取分组信息列表
func (d *directDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	return d.svr.GroupInfoIndex(ctx, in)
}

// 获取分组信息详情
func (m *defaultDeviceGroup) GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoRead(ctx, in, opts...)
}

// 获取分组信息详情
func (d *directDeviceGroup) GroupInfoRead(ctx context.Context, in *GroupInfoReadReq, opts ...grpc.CallOption) (*GroupInfo, error) {
	return d.svr.GroupInfoRead(ctx, in)
}

// 更新分组
func (m *defaultDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoUpdate(ctx, in, opts...)
}

// 更新分组
func (d *directDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupInfoUpdate(ctx, in)
}

// 删除分组
func (m *defaultDeviceGroup) GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoDelete(ctx, in, opts...)
}

// 删除分组
func (d *directDeviceGroup) GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupInfoDelete(ctx, in)
}

// 创建分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiCreate(ctx, in, opts...)
}

// 创建分组设备
func (d *directDeviceGroup) GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiCreate(ctx, in)
}

// 更新分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiUpdate(ctx, in, opts...)
}

// 更新分组设备
func (d *directDeviceGroup) GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiUpdate(ctx, in)
}

// 删除分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiDelete(ctx, in, opts...)
}

// 删除分组设备
func (d *directDeviceGroup) GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiDelete(ctx, in)
}
