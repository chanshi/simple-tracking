package model

import (
	"github.com/chanshi/simple-tracking/support"
	"github.com/chanshi/simple-tracking/support/db"
	"time"
)


func newVisitPage( visit *Visit ) *VisitPage  {
	vp:=new(VisitPage)
	vp.VisitorId = visit.VisitorId
	vp.VisitId   = visit.VisitId
	vp.VisitTime = time.Now()
	vp.VisitSpent= 0
	vp.db        = NewVisitPageTable()
	//id,_ := vp.getVisitPageId()
	//vp.VisitPageId = id
	return vp
}


/**
页面跟踪
*/
type VisitPage struct {
	db                *db.Table
	VisitPageId       int64
	VisitId           int64
	VisitorId         Vid

	RefererPage       string                      //来源页面
	VisitPage         string                      //访问页面
	VisitTime         time.Time                   //访问时间
	VisitSpent        int64                       //访问时长
	Refresh           int                         //刷新次数。
	Data              DataSection
}

func (this *VisitPage) getVisitPageId() (int64,error) {
	return this.db.Insert(GetVisitPageData(this))

}

func (this *VisitPage) Page ( page string) *VisitPage {
	this.VisitPage     = page
	return this
}


func (this *VisitPage) Update()  {
	this.db.Update(db.DataMap{
		"refresh" :this.Refresh,
		"visitSpent":this.VisitSpent})
}

func (this *VisitPage) Action( group,name string , data ...interface{} ) *VisitAction {
	action :=newVisitAction(this)
	action.Action(group,name)
	action.Data.Data(data...)
	action.Save()
	return action
}


func (this *VisitPage) Save()  {
	_,err:=this.db.Insert(GetVisitPageData(this))
	if err!= nil{
		support.Log(err.Error())
	}
}