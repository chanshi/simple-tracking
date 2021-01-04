package model

import (
	"simple-tracking/support/db"
)

type IMetaData interface {
	value()  db.DataMap
}

type MetaData struct {
	Location    MetaLocation
	Device      MetaDevice
	App         MetaApp
	Channel     MetaChannel
}

/**
  可以配置Meta
 */
func (this *MetaData) Data( data ...interface{} )  {
	for _,item:=range data{
		switch item.(type) {
		case MetaLocation :
			this.Location = item.(MetaLocation)
		case MetaDevice:
			this.Device = item.(MetaDevice)
		case MetaApp:
			this.App = item.(MetaApp)
		case MetaChannel:
			this.Channel = item.(MetaChannel)
		default:
		}
	}
}

func (this *MetaData) value() db.DataMap  {
	//合并
	result:=db.DataMap{}
	for k,item:=range this.App.value(){
		result[k] = item
	}
	for k,item:= range this.Channel.value(){
		result[k] =item
	}
	for k,item:=range this.Device.value(){
		result[k] = item
	}
	for k,item:=range this.Location.value(){
		result[k] = item
	}
	return result
}

