// Code generated by goctl. DO NOT EDIT!

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userCoreFieldNames          = builder.RawFieldNames(&UserCore{})
	userCoreRows                = strings.Join(userCoreFieldNames, ",")
	userCoreRowsExpectAutoSet   = strings.Join(stringx.Remove(userCoreFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userCoreRowsWithPlaceHolder = strings.Join(stringx.Remove(userCoreFieldNames, "`uid`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	userCoreModel interface {
		Insert(ctx context.Context, data *UserCore) (sql.Result, error)
		FindOne(ctx context.Context, uid int64) (*UserCore, error)
		FindOneByPhone(ctx context.Context, phone string) (*UserCore, error)
		FindOneByWechat(ctx context.Context, phone string) (*UserCore, error)
		FindOneByUserName(ctx context.Context, userName string) (*UserCore, error)
		Update(ctx context.Context, newData *UserCore) error
		Delete(ctx context.Context, uid int64) error
	}

	defaultUserCoreModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserCore struct {
		Uid         int64        `db:"uid"`      // 用户id
		UserName    string       `db:"userName"` // 登录用户名
		Password    string       `db:"password"` // 登录密码
		Email       string       `db:"email"`    // 邮箱
		Phone       string       `db:"phone"`    // 手机号
		Wechat      string       `db:"wechat"`   // 微信union id
		LastIP      string       `db:"lastIP"`   // 最后登录ip
		RegIP       string       `db:"regIP"`    // 注册ip
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime time.Time    `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
		Status      int64        `db:"status"`      // 用户状态:0为未注册状态
		AuthorityId int64        `db:"authorityId"` // 用户角色
	}
)

func newUserCoreModel(conn sqlx.SqlConn) *defaultUserCoreModel {
	return &defaultUserCoreModel{
		conn:  conn,
		table: "`user_core`",
	}
}

func (m *defaultUserCoreModel) Delete(ctx context.Context, uid int64) error {
	query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, uid)
	return err
}

func (m *defaultUserCoreModel) FindOne(ctx context.Context, uid int64) (*UserCore, error) {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userCoreRows, m.table)
	var resp UserCore
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserCoreModel) FindOneByPhone(ctx context.Context, phone string) (*UserCore, error) {
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userCoreRows, m.table)
	var resp UserCore
	err := m.conn.QueryRowCtx(ctx, &resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserCoreModel) FindOneByWechat(ctx context.Context, weChat string) (*UserCore, error) {
	query := fmt.Sprintf("select %s from %s where `wechat` = ? limit 1", userCoreRows, m.table)
	var resp UserCore
	err := m.conn.QueryRowCtx(ctx, &resp, query, weChat)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserCoreModel) Insert(ctx context.Context, data *UserCore) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userCoreRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uid, data.UserName, data.Password, data.Email, data.Phone, data.Wechat, data.LastIP, data.RegIP, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.Status, data.AuthorityId)
	return ret, err
}

func (m *defaultUserCoreModel) Update(ctx context.Context, data *UserCore) error {
	query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, "`userName`=?,`password`=?,`email`=?,`phone`=?,`wechat`=?,`lastIP`=?,`regIP`=?,`status`=?,`authorityId`=?")
	_, err := m.conn.ExecCtx(ctx, query, data.UserName, data.Password, data.Email, data.Phone, data.Wechat, data.LastIP, data.RegIP, data.Status, data.AuthorityId, data.Uid)
	return err
}
func (m *defaultUserCoreModel) FindOneByUserName(ctx context.Context, userName string) (*UserCore, error) {
	query := fmt.Sprintf("select %s from %s where `userName` = ? limit 1", userCoreRows, m.table)
	var resp UserCore
	err := m.conn.QueryRowCtx(ctx, &resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserCoreModel) tableName() string {
	return m.table
}
