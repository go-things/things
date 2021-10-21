// Code generated by goctl. DO NOT EDIT!
// Source: dc.proto

package dcclient

import (
	"context"

	"gitee.com/godLei6/things/src/dcsvr/dc"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	GroupMember          = dc.GroupMember
	ManageGroupInfoReq   = dc.ManageGroupInfoReq
	GetGroupInfoReq      = dc.GetGroupInfoReq
	GetGroupMemberResp   = dc.GetGroupMemberResp
	SendActionReq        = dc.SendActionReq
	SendPropertyReq      = dc.SendPropertyReq
	SendPropertyResp     = dc.SendPropertyResp
	PageInfo             = dc.PageInfo
	GroupInfo            = dc.GroupInfo
	ManageGroupMemberReq = dc.ManageGroupMemberReq
	GetGroupInfoResp     = dc.GetGroupInfoResp
	GetGroupMemberReq    = dc.GetGroupMemberReq
	SendActionResp       = dc.SendActionResp
	Response             = dc.Response

	Dc interface {
		// 管理组
		ManageGroupInfo(ctx context.Context, in *ManageGroupInfoReq) (*GroupInfo, error)
		// 管理组成员
		ManageGroupMember(ctx context.Context, in *ManageGroupMemberReq) (*GroupMember, error)
		// 获取组信息
		GetGroupInfo(ctx context.Context, in *GetGroupInfoReq) (*GetGroupInfoResp, error)
		// 获取组成员
		GetGroupMember(ctx context.Context, in *GetGroupMemberReq) (*GetGroupMemberResp, error)
		// 同步调用设备行为
		SendAction(ctx context.Context, in *SendActionReq) (*SendActionResp, error)
		// 同步调用设备属性
		SendProperty(ctx context.Context, in *SendPropertyReq) (*SendPropertyResp, error)
	}

	defaultDc struct {
		cli zrpc.Client
	}
)

func NewDc(cli zrpc.Client) Dc {
	return &defaultDc{
		cli: cli,
	}
}

// 管理组
func (m *defaultDc) ManageGroupInfo(ctx context.Context, in *ManageGroupInfoReq) (*GroupInfo, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.ManageGroupInfo(ctx, in)
}

// 管理组成员
func (m *defaultDc) ManageGroupMember(ctx context.Context, in *ManageGroupMemberReq) (*GroupMember, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.ManageGroupMember(ctx, in)
}

// 获取组信息
func (m *defaultDc) GetGroupInfo(ctx context.Context, in *GetGroupInfoReq) (*GetGroupInfoResp, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.GetGroupInfo(ctx, in)
}

// 获取组成员
func (m *defaultDc) GetGroupMember(ctx context.Context, in *GetGroupMemberReq) (*GetGroupMemberResp, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.GetGroupMember(ctx, in)
}

// 同步调用设备行为
func (m *defaultDc) SendAction(ctx context.Context, in *SendActionReq) (*SendActionResp, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.SendAction(ctx, in)
}

// 同步调用设备属性
func (m *defaultDc) SendProperty(ctx context.Context, in *SendPropertyReq) (*SendPropertyResp, error) {
	client := dc.NewDcClient(m.cli.Conn())
	return client.SendProperty(ctx, in)
}
