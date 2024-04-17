// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package remoteconfig

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                       = dm.ActionRespReq
	ActionSendReq                       = dm.ActionSendReq
	ActionSendResp                      = dm.ActionSendResp
	CommonSchemaCreateReq               = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq                = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp               = dm.CommonSchemaIndexResp
	CommonSchemaInfo                    = dm.CommonSchemaInfo
	CommonSchemaUpdateReq               = dm.CommonSchemaUpdateReq
	CustomTopic                         = dm.CustomTopic
	DeviceCore                          = dm.DeviceCore
	DeviceCountInfo                     = dm.DeviceCountInfo
	DeviceCountReq                      = dm.DeviceCountReq
	DeviceCountResp                     = dm.DeviceCountResp
	DeviceGatewayBindDevice             = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq               = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp              = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq         = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq         = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                   = dm.DeviceGatewaySign
	DeviceInfo                          = dm.DeviceInfo
	DeviceInfoBindReq                   = dm.DeviceInfoBindReq
	DeviceInfoCount                     = dm.DeviceInfoCount
	DeviceInfoCountReq                  = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                 = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                  = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                 = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq            = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                   = dm.DeviceInfoReadReq
	DeviceProfile                       = dm.DeviceProfile
	DeviceProfileIndexReq               = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp              = dm.DeviceProfileIndexResp
	DeviceProfileReadReq                = dm.DeviceProfileReadReq
	DeviceTypeCountReq                  = dm.DeviceTypeCountReq
	DeviceTypeCountResp                 = dm.DeviceTypeCountResp
	Empty                               = dm.Empty
	EventLogIndexReq                    = dm.EventLogIndexReq
	EventLogIndexResp                   = dm.EventLogIndexResp
	EventLogInfo                        = dm.EventLogInfo
	Firmware                            = dm.Firmware
	FirmwareFile                        = dm.FirmwareFile
	FirmwareInfo                        = dm.FirmwareInfo
	FirmwareInfoDeleteReq               = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp              = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq                = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp               = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                 = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp                = dm.FirmwareInfoReadResp
	FirmwareResp                        = dm.FirmwareResp
	GroupDeviceIndexReq                 = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp                = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq           = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq             = dm.GroupDeviceMultiSaveReq
	GroupInfo                           = dm.GroupInfo
	GroupInfoCreateReq                  = dm.GroupInfoCreateReq
	GroupInfoIndexReq                   = dm.GroupInfoIndexReq
	GroupInfoIndexResp                  = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                  = dm.GroupInfoUpdateReq
	HubLogIndexReq                      = dm.HubLogIndexReq
	HubLogIndexResp                     = dm.HubLogIndexResp
	HubLogInfo                          = dm.HubLogInfo
	OtaFirmwareDeviceCancelReq          = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceIndexReq           = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp          = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo               = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq           = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                     = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq             = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp            = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                 = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                  = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                 = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                     = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq            = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq             = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp            = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq            = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq              = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp             = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                  = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq              = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                   = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                    = dm.OtaJobStaticInfo
	OtaModuleInfo                       = dm.OtaModuleInfo
	OtaModuleInfoIndexReq               = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp              = dm.OtaModuleInfoIndexResp
	OtaPromptIndexReq                   = dm.OtaPromptIndexReq
	OtaPromptIndexResp                  = dm.OtaPromptIndexResp
	PageInfo                            = dm.PageInfo
	PageInfo_OrderBy                    = dm.PageInfo_OrderBy
	Point                               = dm.Point
	ProductCategory                     = dm.ProductCategory
	ProductCategoryIndexReq             = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp            = dm.ProductCategoryIndexResp
	ProductCategoryReadReq              = dm.ProductCategoryReadReq
	ProductCategorySchemaIndexReq       = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp      = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiUpdateReq = dm.ProductCategorySchemaMultiUpdateReq
	ProductCustom                       = dm.ProductCustom
	ProductCustomReadReq                = dm.ProductCustomReadReq
	ProductInfo                         = dm.ProductInfo
	ProductInfoDeleteReq                = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                 = dm.ProductInfoIndexReq
	ProductInfoIndexResp                = dm.ProductInfoIndexResp
	ProductInfoReadReq                  = dm.ProductInfoReadReq
	ProductInitReq                      = dm.ProductInitReq
	ProductRemoteConfig                 = dm.ProductRemoteConfig
	ProductSchemaCreateReq              = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq              = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq               = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp              = dm.ProductSchemaIndexResp
	ProductSchemaInfo                   = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq         = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq           = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq             = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp            = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq              = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq         = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp        = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg              = dm.PropertyControlSendMsg
	PropertyControlSendReq              = dm.PropertyControlSendReq
	PropertyControlSendResp             = dm.PropertyControlSendResp
	PropertyGetReportSendReq            = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp           = dm.PropertyGetReportSendResp
	PropertyLogIndexReq                 = dm.PropertyLogIndexReq
	PropertyLogIndexResp                = dm.PropertyLogIndexResp
	PropertyLogInfo                     = dm.PropertyLogInfo
	PropertyLogLatestIndexReq           = dm.PropertyLogLatestIndexReq
	ProtocolConfigField                 = dm.ProtocolConfigField
	ProtocolConfigInfo                  = dm.ProtocolConfigInfo
	ProtocolInfo                        = dm.ProtocolInfo
	ProtocolInfoIndexReq                = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp               = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq               = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq                = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp               = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq             = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp            = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq              = dm.RemoteConfigPushAllReq
	RespReadReq                         = dm.RespReadReq
	RootCheckReq                        = dm.RootCheckReq
	SdkLogIndexReq                      = dm.SdkLogIndexReq
	SdkLogIndexResp                     = dm.SdkLogIndexResp
	SdkLogInfo                          = dm.SdkLogInfo
	SendLogIndexReq                     = dm.SendLogIndexReq
	SendLogIndexResp                    = dm.SendLogIndexResp
	SendLogInfo                         = dm.SendLogInfo
	SendMsgReq                          = dm.SendMsgReq
	SendMsgResp                         = dm.SendMsgResp
	SendOption                          = dm.SendOption
	ShadowIndex                         = dm.ShadowIndex
	ShadowIndexResp                     = dm.ShadowIndexResp
	StatusLogIndexReq                   = dm.StatusLogIndexReq
	StatusLogIndexResp                  = dm.StatusLogIndexResp
	StatusLogInfo                       = dm.StatusLogInfo
	TimeRange                           = dm.TimeRange
	UserDeviceCollectSave               = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq             = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp            = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo                 = dm.UserDeviceShareInfo
	UserDeviceShareMultiDeleteReq       = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareReadReq              = dm.UserDeviceShareReadReq
	WithID                              = dm.WithID
	WithIDCode                          = dm.WithIDCode

	RemoteConfig interface {
		RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Empty, error)
		RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error)
		RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Empty, error)
		RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error)
	}

	defaultRemoteConfig struct {
		cli zrpc.Client
	}

	directRemoteConfig struct {
		svcCtx *svc.ServiceContext
		svr    dm.RemoteConfigServer
	}
)

func NewRemoteConfig(cli zrpc.Client) RemoteConfig {
	return &defaultRemoteConfig{
		cli: cli,
	}
}

func NewDirectRemoteConfig(svcCtx *svc.ServiceContext, svr dm.RemoteConfigServer) RemoteConfig {
	return &directRemoteConfig{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRemoteConfig) RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigCreate(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigCreate(ctx context.Context, in *RemoteConfigCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.RemoteConfigCreate(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigIndex(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigIndex(ctx context.Context, in *RemoteConfigIndexReq, opts ...grpc.CallOption) (*RemoteConfigIndexResp, error) {
	return d.svr.RemoteConfigIndex(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigPushAll(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigPushAll(ctx context.Context, in *RemoteConfigPushAllReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.RemoteConfigPushAll(ctx, in)
}

func (m *defaultRemoteConfig) RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error) {
	client := dm.NewRemoteConfigClient(m.cli.Conn())
	return client.RemoteConfigLastRead(ctx, in, opts...)
}

func (d *directRemoteConfig) RemoteConfigLastRead(ctx context.Context, in *RemoteConfigLastReadReq, opts ...grpc.CallOption) (*RemoteConfigLastReadResp, error) {
	return d.svr.RemoteConfigLastRead(ctx, in)
}
