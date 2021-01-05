package model

import (
	"github.com/chanshi/simple-tracking/support"
	"github.com/chanshi/simple-tracking/support/db"
	"testing"
	"time"
)

func TestGetVisitor(t *testing.T) {
	db.Get().Init()
	defer db.Get().Close()
	support.GetGeo().Init()
	defer support.GetGeo().Close()

	//访问者
	visitor := GetVisitor(12)
	visitor.Data.Data("0","1",2,"3",4)
	visitor.Meta.Data(
		MetaApp{ AppName: "bate",  AppVersion: "v1.0.02" },
		MetaLocation{ Ip: "210.12.108.189"},
		MetaChannel{ChannelName: "local"},
		MetaDevice{ DeviceId: "44444!!!!!33",  DeviceOS: "IOS", DeviceOsVersion: "7", DeviceBrand: "HuaWei"},
	)
	_,err:=visitor.Save()
	if err!=nil{
		println(err.Error())
	}

	//访问
	visit:= NewVisit(visitor)

	//浏览页面
	visit.VisitPage("home")

	visit.Action("view","page",1,2)
	time.Sleep(time.Second * 5)

	//浏览页面
	visit.VisitPage("mine","10",12)

	//事件。
	visit.Event("充值","100",200)

	visit.Close()


}
