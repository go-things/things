package devicemanagelogic

import (
	"encoding/json"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/domain/device"
	mysql "github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

func ToDeviceInfo(di *mysql.DeviceInfo) *dm.DeviceInfo {
	var (
		tags map[string]string
	)

	_ = json.Unmarshal([]byte(di.Tags), &tags)

	if di.IsOnline == def.Unknown {
		di.IsOnline = def.False
	}
	if di.LogLevel == def.Unknown {
		di.LogLevel = def.LogClose
	}
	return &dm.DeviceInfo{
		Version:     &wrappers.StringValue{Value: di.Version},
		LogLevel:    di.LogLevel,
		Cert:        di.Cert,
		ProductID:   di.ProductID,
		DeviceName:  di.DeviceName,
		CreatedTime: di.CreatedTime.Unix(),
		FirstLogin:  utils.GetNullTime(di.FirstLogin),
		LastLogin:   utils.GetNullTime(di.LastLogin),
		Secret:      di.Secret,
		IsOnline:    di.IsOnline,
		Tags:        tags,
	}
}
func ToDeviceCoreDos(in []*dm.DeviceCore) (ret []*device.Core) {
	for _, v := range in {
		ret = append(ret, &device.Core{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
		})
	}
	return
}
