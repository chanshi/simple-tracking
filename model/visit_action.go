package model

import (
	"simple-tracking/support/db"
	"time"
)

/**
  用户行为跟踪。
 */
type VisitAction struct {
	db                *db.Table
	ActionId          int64
	VisitId           int64
	VisitorId         Vid
	VisitPageId       int64                // 时间发生的页面

	ActionGroup       string               // 行为分组
	ActionName        string               // 行为名称
	Data              DataSection
	ActionTime        time.Time
}

func newVisitAction( page *VisitPage ) *VisitAction {
	return &VisitAction{
		ActionId:       0,
		VisitId:        page.VisitId,
		VisitPageId:    page.VisitPageId,
		VisitorId:      page.VisitorId,
		ActionGroup:     "",
		ActionName:      "",
		ActionTime:     time.Now(),
		db : NewVisitActionTable(),
	}
}

func (this *VisitAction) Action( group,name string ) *VisitAction {
	this.ActionGroup = group
	this.ActionName  = name
	return this
}

func (this *VisitAction) Save()  {
	data := GetVisitActionData(this)
	this.db.Insert(data)
}