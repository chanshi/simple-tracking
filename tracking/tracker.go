package tracking

import (
	"github.com/chanshi/simple-tracking/model"
	"github.com/chanshi/simple-tracking/support"
	"time"
)

func NewTracker( VisitorId model.Vid ) *Tracker {
	tracker:= &Tracker{}
	tracker.setVisitor(VisitorId)
	return tracker
}


type Tracker struct {
	IsCanTracker bool
	visitor *model.Visitor
	proxy   *model.Visit
}


func (this *Tracker) setVisitor( visitorId model.Vid)  {
	this.visitor = model.GetVisitor(visitorId)
}

func (this *Tracker) SetFirstVisitTime( tm time.Time ) *Tracker {
	this.visitor.FirstVisitTime = tm
	return this
}

/**

 */

func (this *Tracker) Begin() *Tracker  {
	_,err:=this.visitor.Save()
	if err!=nil{
		support.Log("Visitor save error ==>",err.Error())
		return this
	}
	this.proxy = model.NewVisit(this.visitor)
	support.Log("Visitor Visit Begin ==>",this.visitor.VisitorId)
	return this
}

func (this *Tracker) CustomData( data ...interface{} ) *Tracker {
	this.visitor.Data.Data(data...)
	return this
}

func (this *Tracker) App( appName ,appVersion  string) *Tracker  {
	this.visitor.Meta.Data(
		model.MetaApp{
			AppName:    appName,
			AppVersion: appVersion,
	})
	return this
}

func (this  *Tracker) Ip(ip string ) *Tracker {
	this.visitor.Meta.Data(
		model.MetaLocation{
			Ip: ip,
		})
	return this
}

func (this *Tracker) Channel( channelName string  ) *Tracker {
	this.visitor.Meta.Data(
		model.MetaChannel{
			ChannelName: channelName,
		})
	return this
}

func (this *Tracker) Device( id ,os,osVersion,brand  string) *Tracker {
	this.visitor.Meta.Data(
		model.MetaDevice{
			DeviceId: id,
			DeviceOS: os,
			DeviceOsVersion: osVersion,
			DeviceBrand: brand,
		})
	return this
}

func (this *Tracker) Page( pageName string ,data ...interface{} ) *Tracker  {
	if  support.Config().EnableTracking && this.proxy !=nil{
		this.proxy.VisitPage(pageName,data... )
	}
	return this
}

func (this *Tracker) Action( group ,name string ,data ...interface{} ) *Tracker {
	if  support.Config().EnableTracking && this.proxy !=nil{
		this.proxy.Action(group,name,data... )
	}
	return this
}

func (this *Tracker) Event( name string,data ...interface{} ) *Tracker {
	if  support.Config().EnableTracking && this.proxy !=nil{
		this.proxy.Event(name,data... )
	}
	return this
}

func (this *Tracker) End()  {
	if this.proxy !=nil{
		this.proxy.Close()
		this.proxy   = nil
		this.visitor = nil
	}
}