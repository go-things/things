package authlogic

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/spf13/cast"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorityApiMultiUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthorityApiMultiUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorityApiMultiUpdateLogic {
	return &AuthorityApiMultiUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthorityApiMultiUpdateLogic) AuthorityApiMultiUpdate(in *sys.AuthorityApiMultiUpdateReq) (*sys.Response, error) {
	// clear old policies
	var oldPolicies [][]string
	oldPolicies = l.svcCtx.Casbin.GetFilteredPolicy(0, cast.ToString(in.RoleID))
	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, cast.ToString(in.RoleID))
		if err != nil {
			l.Errorf("%s.Casbin.RemoveFilteredPolicy req=%v err=%+v", utils.FuncName(), in, err)
			return nil, errors.Permissions.AddDetail(err)
		}
		if !removeResult {
			l.Errorf("%s.Casbin.RemoveFilteredPolicy req=%v", utils.FuncName(), in)
			return nil, errors.System.AddDetail("RemoveFilteredPolicy Failed")
		}
	}

	// add new policies
	var policies [][]string
	for _, v := range in.List {
		policies = append(policies, []string{cast.ToString(in.RoleID), v.Route, cast.ToString(v.Method)})
	}
	addResult, err := l.svcCtx.Casbin.AddPolicies(policies) //如果是删除策略是否还有用?
	if err != nil {
		err := errors.Fmt(err)
		l.Errorf("%s.Casbin.AddPolicies req=%v err=%+v", utils.FuncName(), in, err)
		return nil, errors.Permissions.AddDetail(err)
	}
	if !addResult {
		l.Errorf("%s Casbin.AddPolicies return nil req=%+v", utils.FuncName(), in)
		return nil, errors.System.AddDetail(err)
	}

	return &sys.Response{}, nil
}
