package model

import (
	"github.com/chanshi/simple-tracking/support/db"
	"time"
)

/**
  普通事件
*/
type VisitEvent struct {
	db                 *db.Table
	EventId           int64                // 事件ID
	VisitId           int64
	VisitorId         Vid                  // 访问

	EventName         string               // 事件
	Data              DataSection
	EventTime         time.Time            // 事件时间
}

func newVisitEvent( v *Visit ) *VisitEvent {
	ve:= &VisitEvent{}
	ve.VisitId    = v.VisitId
	ve.VisitorId  = v.VisitorId
	ve.EventTime = time.Now()

	ve.db = NewVisitEventTable()
	return ve
}


func (this *VisitEvent) Event( name string ) *VisitEvent {
	this.EventName = name
	return this
}


func (this *VisitEvent) Save()  {
	data := GetVisitEventData(this)
	this.db.Insert(data)
}