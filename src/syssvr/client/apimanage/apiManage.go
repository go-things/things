// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package apimanage

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq              = sys.ApiDeleteReq
	ApiInfo                   = sys.ApiInfo
	ApiInfoIndexReq           = sys.ApiInfoIndexReq
	ApiInfoIndexResp          = sys.ApiInfoIndexResp
	AppInfo                   = sys.AppInfo
	AppInfoIndexReq           = sys.AppInfoIndexReq
	AppInfoIndexResp          = sys.AppInfoIndexResp
	AppModuleIndexReq         = sys.AppModuleIndexReq
	AppModuleIndexResp        = sys.AppModuleIndexResp
	AppModuleMultiUpdateReq   = sys.AppModuleMultiUpdateReq
	AreaInfo                  = sys.AreaInfo
	AreaInfoIndexReq          = sys.AreaInfoIndexReq
	AreaInfoIndexResp         = sys.AreaInfoIndexResp
	AreaInfoReadReq           = sys.AreaInfoReadReq
	AreaWithID                = sys.AreaWithID
	AuthApiInfo               = sys.AuthApiInfo
	ConfigResp                = sys.ConfigResp
	DateRange                 = sys.DateRange
	JwtToken                  = sys.JwtToken
	LoginLogCreateReq         = sys.LoginLogCreateReq
	LoginLogIndexReq          = sys.LoginLogIndexReq
	LoginLogIndexResp         = sys.LoginLogIndexResp
	LoginLogInfo              = sys.LoginLogInfo
	Map                       = sys.Map
	MenuInfo                  = sys.MenuInfo
	MenuInfoIndexReq          = sys.MenuInfoIndexReq
	MenuInfoIndexResp         = sys.MenuInfoIndexResp
	ModuleInfo                = sys.ModuleInfo
	ModuleInfoIndexReq        = sys.ModuleInfoIndexReq
	ModuleInfoIndexResp       = sys.ModuleInfoIndexResp
	OperLogCreateReq          = sys.OperLogCreateReq
	OperLogIndexReq           = sys.OperLogIndexReq
	OperLogIndexResp          = sys.OperLogIndexResp
	OperLogInfo               = sys.OperLogInfo
	PageInfo                  = sys.PageInfo
	PageInfo_OrderBy          = sys.PageInfo_OrderBy
	Point                     = sys.Point
	ProjectInfo               = sys.ProjectInfo
	ProjectInfoIndexReq       = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp      = sys.ProjectInfoIndexResp
	ProjectWithID             = sys.ProjectWithID
	Response                  = sys.Response
	RoleApiAuthReq            = sys.RoleApiAuthReq
	RoleApiIndexReq           = sys.RoleApiIndexReq
	RoleApiIndexResp          = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq     = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq           = sys.RoleAppIndexReq
	RoleAppIndexResp          = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq     = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq          = sys.RoleAppUpdateReq
	RoleInfo                  = sys.RoleInfo
	RoleInfoIndexReq          = sys.RoleInfoIndexReq
	RoleInfoIndexResp         = sys.RoleInfoIndexResp
	RoleMenuIndexReq          = sys.RoleMenuIndexReq
	RoleMenuIndexResp         = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq    = sys.RoleMenuMultiUpdateReq
	RoleModuleIndexReq        = sys.RoleModuleIndexReq
	RoleModuleIndexResp       = sys.RoleModuleIndexResp
	RoleModuleMultiUpdateReq  = sys.RoleModuleMultiUpdateReq
	TenantApiInfo             = sys.TenantApiInfo
	TenantAppApiIndexReq      = sys.TenantAppApiIndexReq
	TenantAppApiIndexResp     = sys.TenantAppApiIndexResp
	TenantAppCreateReq        = sys.TenantAppCreateReq
	TenantAppIndexReq         = sys.TenantAppIndexReq
	TenantAppIndexResp        = sys.TenantAppIndexResp
	TenantAppMenu             = sys.TenantAppMenu
	TenantAppMenuIndexReq     = sys.TenantAppMenuIndexReq
	TenantAppMenuIndexResp    = sys.TenantAppMenuIndexResp
	TenantAppModule           = sys.TenantAppModule
	TenantAppMultiUpdateReq   = sys.TenantAppMultiUpdateReq
	TenantAppWithIDOrCode     = sys.TenantAppWithIDOrCode
	TenantInfo                = sys.TenantInfo
	TenantInfoCreateReq       = sys.TenantInfoCreateReq
	TenantInfoIndexReq        = sys.TenantInfoIndexReq
	TenantInfoIndexResp       = sys.TenantInfoIndexResp
	TenantModuleCreateReq     = sys.TenantModuleCreateReq
	TenantModuleIndexReq      = sys.TenantModuleIndexReq
	TenantModuleIndexResp     = sys.TenantModuleIndexResp
	TenantModuleWithIDOrCode  = sys.TenantModuleWithIDOrCode
	UserArea                  = sys.UserArea
	UserAreaApplyCreateReq    = sys.UserAreaApplyCreateReq
	UserAreaApplyDealReq      = sys.UserAreaApplyDealReq
	UserAreaApplyIndexReq     = sys.UserAreaApplyIndexReq
	UserAreaApplyIndexResp    = sys.UserAreaApplyIndexResp
	UserAreaApplyInfo         = sys.UserAreaApplyInfo
	UserAreaIndexReq          = sys.UserAreaIndexReq
	UserAreaIndexResp         = sys.UserAreaIndexResp
	UserAreaMultiDeleteReq    = sys.UserAreaMultiDeleteReq
	UserAreaMultiUpdateReq    = sys.UserAreaMultiUpdateReq
	UserCaptchaReq            = sys.UserCaptchaReq
	UserCaptchaResp           = sys.UserCaptchaResp
	UserChangePwdReq          = sys.UserChangePwdReq
	UserCheckTokenReq         = sys.UserCheckTokenReq
	UserCheckTokenResp        = sys.UserCheckTokenResp
	UserCreateResp            = sys.UserCreateResp
	UserForgetPwdReq          = sys.UserForgetPwdReq
	UserInfo                  = sys.UserInfo
	UserInfoCreateReq         = sys.UserInfoCreateReq
	UserInfoDeleteReq         = sys.UserInfoDeleteReq
	UserInfoIndexReq          = sys.UserInfoIndexReq
	UserInfoIndexResp         = sys.UserInfoIndexResp
	UserInfoReadReq           = sys.UserInfoReadReq
	UserLoginReq              = sys.UserLoginReq
	UserLoginResp             = sys.UserLoginResp
	UserProject               = sys.UserProject
	UserProjectIndexReq       = sys.UserProjectIndexReq
	UserProjectIndexResp      = sys.UserProjectIndexResp
	UserProjectMultiUpdateReq = sys.UserProjectMultiUpdateReq
	UserRegisterReq           = sys.UserRegisterReq
	UserRegisterResp          = sys.UserRegisterResp
	UserRoleIndexReq          = sys.UserRoleIndexReq
	UserRoleIndexResp         = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq    = sys.UserRoleMultiUpdateReq
	WithAppCodeID             = sys.WithAppCodeID
	WithID                    = sys.WithID
	WithIDCode                = sys.WithIDCode

	ApiManage interface {
		ApiInfoCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error)
		ApiInfoIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error)
		ApiInfoUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error)
		ApiInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
	}

	defaultApiManage struct {
		cli zrpc.Client
	}

	directApiManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.ApiManageServer
	}
)

func NewApiManage(cli zrpc.Client) ApiManage {
	return &defaultApiManage{
		cli: cli,
	}
}

func NewDirectApiManage(svcCtx *svc.ServiceContext, svr sys.ApiManageServer) ApiManage {
	return &directApiManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultApiManage) ApiInfoCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewApiManageClient(m.cli.Conn())
	return client.ApiInfoCreate(ctx, in, opts...)
}

func (d *directApiManage) ApiInfoCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ApiInfoCreate(ctx, in)
}

func (m *defaultApiManage) ApiInfoIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error) {
	client := sys.NewApiManageClient(m.cli.Conn())
	return client.ApiInfoIndex(ctx, in, opts...)
}

func (d *directApiManage) ApiInfoIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error) {
	return d.svr.ApiInfoIndex(ctx, in)
}

func (m *defaultApiManage) ApiInfoUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewApiManageClient(m.cli.Conn())
	return client.ApiInfoUpdate(ctx, in, opts...)
}

func (d *directApiManage) ApiInfoUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ApiInfoUpdate(ctx, in)
}

func (m *defaultApiManage) ApiInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewApiManageClient(m.cli.Conn())
	return client.ApiInfoDelete(ctx, in, opts...)
}

func (d *directApiManage) ApiInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ApiInfoDelete(ctx, in)
}
