package model

import (
	"simple-tracking/support/db"
	"strconv"
)

/**
 设置最大的段位制 5
 */
const MAX_DataSection = 5

type DataSection struct {
 	 DataSection          []interface{}
}

func (this *DataSection) Data( data ...interface{} ) {
	this.DataSection = data
}

func (this *DataSection) value() db.DataMap  {
	value := make(db.DataMap)
	for i:=0;i<len(this.DataSection);i++{
		if i >= MAX_DataSection{
			break
		}
		dataIndex := strconv.Itoa(i)
		value["data_"+dataIndex+""]=this.DataSection[i]
	}
	return value
}

func DataSectionField() []db.Fields  {
	fields:=[]db.Fields{}
	for i:=0;i< MAX_DataSection;i++{
		dataIndex := strconv.Itoa(i)
		fields = append(fields,db.Fields{FieldName: "data_"+dataIndex})
	}
	return fields
}
