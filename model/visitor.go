package model

import (
	"simple-tracking/support/db"
	"time"
)

func GetVisitor( visitorId  Vid) *Visitor {
	visitor := new(Visitor)
	visitor.VisitorId = visitorId
	visitor.db = NewVisitorTable()
	visitor.FirstVisitTime = time.Now()
	visitor.LastVisitTime  = time.Now()
	return visitor
}

//访客
type Visitor struct {
	VisitorId               Vid                // 用户ID
	FirstVisitTime          time.Time          // 首次访问时间 ，确认是否为新用户。
	Meta                    MetaData
	Data                    DataSection
	LastVisitTime           time.Time          // 最后活动时间
	db                      *db.Table
}


func (this *Visitor) Save() (int64,error) {
	_,err:= this.db.Record(this.VisitorId)
	data:= GetVisitorData(this)
	if err!=nil{
		return this.db.Insert(data)
	}else{
		// 智能保存部分数据 自定义数据 LastVisitTime
		result:= db.DataMap{}
		result["lastVisitTime"]     = this.LastVisitTime

		for k,item:=range this.Data.value(){
			result[k] =item
		}
		return this.db.Update(result)
	}
}

