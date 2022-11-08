package sdkLogRepo

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/store"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"github.com/zeromicro/go-zero/core/logx"
)

func (d SDKLogRepo) GetDeviceSDKLog(ctx context.Context,
	filter deviceMsg.SdkLogFilter, page def.PageInfo2) ([]*deviceMsg.SDKLog, error) {
	sqSql := sq.Select("*").From(d.GetSDKLogStableName())
	sqSql = d.fillFilter(sqSql, filter)
	sqSql = page.FmtSql(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := d.t.QueryContext(ctx, sqlStr, value...)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return []*deviceMsg.SDKLog{}, nil
		}
	}
	var datas []map[string]any
	store.Scan(rows, &datas)
	retLogs := make([]*deviceMsg.SDKLog, 0, len(datas))
	for _, v := range datas {
		retLogs = append(retLogs, ToDeviceSDKLog(filter.ProductID, v))
	}
	return retLogs, nil
}

func (d SDKLogRepo) Insert(ctx context.Context, data *deviceMsg.SDKLog) error {
	sql := fmt.Sprintf(
		"insert into %s using %s tags('%s','%s')(`ts`, `content`,`log_level`,`client_token`) values (?,?,?,?);",
		d.GetSDKLogTableName(data.ProductID, data.DeviceName), d.GetSDKLogStableName(), data.ProductID, data.DeviceName)

	if _, err := d.t.ExecContext(ctx, sql, data.Timestamp, data.Content, data.LogLevel, data.ClientToken); err != nil {
		logx.WithContext(ctx).Errorf(
			sql+"%s.EventTable productID:%v deviceName:%v err:%v",
			utils.FuncName(), data.ProductID, data.DeviceName, err)
		return err
	}
	return nil
}
func (d SDKLogRepo) fillFilter(sql sq.SelectBuilder, filter deviceMsg.SdkLogFilter) sq.SelectBuilder {
	if len(filter.ProductID) != 0 {
		sql = sql.Where("`product_id`=?", filter.ProductID)
	}
	if len(filter.DeviceName) != 0 {
		sql = sql.Where("`device_name`=?", filter.DeviceName)
	}
	if filter.LogLevel != 0 {
		sql = sql.Where("`log_level`=?", filter.LogLevel)
	}
	return sql
}
func (d SDKLogRepo) GetCountLog(ctx context.Context, filter deviceMsg.SdkLogFilter, page def.PageInfo2) (int64, error) {
	sqSql := sq.Select("Count(1)").From(d.GetSDKLogStableName())
	sqSql = d.fillFilter(sqSql, filter)
	sqSql = page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := d.t.QueryRowContext(ctx, sqlStr, value...)
	if err != nil {
		return 0, err
	}
	var (
		total int64
	)

	err = row.Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}
