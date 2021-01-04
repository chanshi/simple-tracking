package model

import "simple-tracking/support/db"

type MetaDevice struct {
	//移动设备信息
	DeviceId        string                  // 设配ID
	DeviceOS        string                  // 设备操作系统
	DeviceOsVersion string                  // 设备系统版本
	DeviceBrand     string                  // 设备品牌
}

func (this MetaDevice) value() db.DataMap  {
	value:= make(db.DataMap)
	value["deviceId"]        = this.DeviceId
	value["deviceOs"]        = this.DeviceOS
	value["deviceOsVersion"] = this.DeviceOsVersion
	value["deviceBrand"]     = this.DeviceBrand
	return value
}
