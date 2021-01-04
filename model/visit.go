package model

import (
	"simple-tracking/support"
	"simple-tracking/support/db"
	"sync"
	"time"
)

type Vid  int64

func NewVisit( visitor *Visitor ) *Visit {
	visit:= new(Visit)
	visit.VisitorId      = visitor.VisitorId
	visit.VisitTime      = time.Now()
	visit.LastActiveTime = time.Now()

	visit.db = NewVisitTable()
	visit.Meta = visitor.Meta
	visit.Data = visitor.Data
	visit.GetVisitId()

	//visit.currentPage = nil

	return visit
}

/**
   访问模型
 */
type Visit struct {
	db                       *db.Table
	VisitId                  int64                    // 会话ID
	VisitorId                Vid                      // 访客

	Meta                     MetaData
	Data                     DataSection

	RefererPage              string                   //
	EnterVisitPage           string                   // 访问的页面
	ExitVisitPage            string                   // 退出的页面
	VisitTime                time.Time
	LastActiveTime           time.Time                // 最后活动时间

	currentPage              *VisitPage

	sync.RWMutex
}

func (this *Visit) GetVisitId() int64 {
	id,err:=this.db.Insert(GetVisitData(this))
	if err!=nil{
		support.Log(err.Error())
		return 0
	}else{
		this.VisitId = id
	}
	return id
}

func (this *Visit) VisitPage( page string ,data ...interface{} ) *VisitPage {
	//判断刷新
	if this.currentPage != nil &&
		this.currentPage.VisitPage == page &&
		time.Now().Unix() - this.currentPage.VisitTime.Unix() < 2 {
		    this.currentPage.Refresh++
		    return this.currentPage
	}

	vp := newVisitPage(this).Page(page)
	vp.Data.Data(data...)
	if this.currentPage == nil{
		this.EnterVisitPage = page
	}else{
		vp.RefererPage = this.currentPage.VisitPage
		this.currentPage.VisitSpent = vp.VisitTime.Unix()- this.currentPage.VisitTime.Unix()
		this.currentPage.Update()
	}
	this.ExitVisitPage = page
	this.currentPage = vp
	vp.Save()
	this.updateLastTime()
	return vp
}

func (this *Visit) Action( group,name string ,data ...interface{} )  {
	if this.currentPage !=nil{
		this.currentPage.Action(group,name,data...)
	}
}

func (this *Visit) Event( name string , data ...interface{} )  {
	event:=newVisitEvent(this).Event(name)
	event.Data.Data(data...)
	event.Save()
	this.updateLastTime()
}

func (this *Visit) updateLastTime()  {
	this.LastActiveTime   = time.Now()
}

func (this *Visit) Close()  {
	this.db.Update(db.DataMap{
		"enterVisitPage":this.EnterVisitPage,
		"exitVisitPage" :this.ExitVisitPage,
		"lastActiveTime":this.LastActiveTime,
	})
}




