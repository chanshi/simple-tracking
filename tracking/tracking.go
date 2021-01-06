package tracking

import (
	"github.com/chanshi/simple-tracking/model"
	"github.com/chanshi/simple-tracking/support"
	"github.com/chanshi/simple-tracking/support/db"
	"sync"
)

/**
Tracking.Instance.Init()
Tracking.Instance.Start()
Tracking.Instance.Visit().Meta().Data()
Tracking.Instance.Stop()

 */


var server *Tracking
func init()  { server = &Tracking{} }
func Track() *Tracking  { return server }

/**
    对外的接口。
包含多个职责。
 */

type Tracking struct {
	conf *support.TrackingConfig
	sync.RWMutex
	item map[model.Vid] *Tracker
}

/**
 关注配置。
 */
func (this *Tracking) Init() {
	db.Get().Init()
	support.GetGeo().Init()
	this.conf = support.Config()
	this.item = make(map[model.Vid] *Tracker)
}


func (this *Tracking) EnableTracking( ) *Tracking  {
	this.conf.SetEnableTracking(true)
	return this
}
func (this *Tracking) DisableTracking() *Tracking  {
	this.conf.SetEnableTracking(false)
	return this
}

/**
 稍微封装下。
 */
func (this *Tracking) Config( config *Config ) *Tracking {
	if config !=nil{
		base :=support.Config()
		base.EnableTracking = config.EnableTracking
		base.DSN = config.Dsn
		base.MaxOpen = config.MaxOpen
		base.GeoCity = config.GeoCity
		support.SetConfig(base)
	}
	return this
}

func (this *Tracking) Visitor( visitorId int64 ) *Tracker {
	this.RLock()
	visitor,ok:=this.item[model.Vid(visitorId)]
	this.RUnlock()
	if ok && visitor.visitor!=nil{

		return visitor
	}else{
		this.Lock()
		this.item[model.Vid(visitorId)] = NewTracker(model.Vid(visitorId ) )
		this.Unlock()
		return this.item[model.Vid(visitorId)]
	}
}

func (this *Tracking) ExitVisitor( visitorId int64 )  {
	this.Lock()
	defer this.Unlock()
	id := model.Vid(visitorId)
	visitor ,ok:=this.item[id]
	if ok{
		//
		visitor.End()
		delete(this.item,model.Vid(visitorId))
	}
}

/**

 */
func (this *Tracking) Stop()  {
	//this.visits.stop()

	//
	db.Get().Close()
	support.GetGeo().Close()
}
