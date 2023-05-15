package svc

import (
	"github.com/casbin/casbin/v2"
	cas "github.com/i-Things/things/shared/casbin"
	"github.com/i-Things/things/shared/third/weixin"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/config"
	"github.com/i-Things/things/src/syssvr/internal/repo/mysql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel mysql.SysUserInfoModel
	RoleInfoModle mysql.SysRoleInfoModel
	MenuInfoModle mysql.SysMenuInfoModel
	RoleMenuModle mysql.SysRoleMenuModel
	UserModel     mysql.UserModel
	RoleModel     mysql.RoleModel
	MenuModel     mysql.MenuModel
	WxMiniProgram *weixin.MiniProgram
	UserID        *utils.SnowFlake
	LogLoginModel mysql.SysLoginLogModel
	LogOperModel  mysql.SysOperLogModel
	LogModel      mysql.LogModel
	ApiModel      mysql.SysApiInfoModel
	Casbin        *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ui := mysql.NewSysUserInfoModel(conn)
	ro := mysql.NewSysRoleInfoModel(conn)
	me := mysql.NewSysMenuInfoModel(conn)
	ll := mysql.NewSysLoginLogModel(conn)
	lo := mysql.NewSysOperLogModel(conn)
	l := mysql.NewLogModel(conn)
	api := mysql.NewSysApiInfoModel(conn)
	rom := mysql.NewRoleModel(conn, c.CacheRedis)
	mem := mysql.NewMenuModel(conn, c.CacheRedis)
	um := mysql.NewUserModel(conn, c.CacheRedis)

	WxMiniProgram := weixin.NewWexinMiniProgram(c.WexinMiniprogram, c.CacheRedis)
	nodeId := utils.GetNodeID(c.CacheRedis, c.Name)
	UserID := utils.NewSnowFlake(nodeId)
	db, _ := conn.RawDB()
	ca := cas.NewCasbinWithRedisWatcher(db, "mysql", c.CacheRedis[0].RedisConf)

	return &ServiceContext{
		Config:        c,
		UserInfoModel: ui,
		UserModel:     um,
		RoleModel:     rom,
		MenuModel:     mem,
		RoleInfoModle: ro,
		MenuInfoModle: me,
		WxMiniProgram: WxMiniProgram,
		UserID:        UserID,
		LogLoginModel: ll,
		LogOperModel:  lo,
		LogModel:      l,
		ApiModel:      api,
		Casbin:        ca,
	}
}
