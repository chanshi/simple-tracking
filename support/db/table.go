package db

import (
	"github.com/chanshi/simple-tracking/support"
	"sync"
	"time"
)

const (
	FieldType_String   = 1
	FieldType_Int      = 2
	FieldType_DATETIME = 3
)


type Fields struct {
	FieldName      string         `json:"fieldName"`
	IsPrimary      bool           `json:"isPrimary"`
	DefaultData    interface{}    `json:"defaultData"`
	FieldType      int            `json:"dataType"`              // 字段类型
}

func NewTable( tableName string ) *Table  {
	table:=&Table{
		RWMutex:     sync.RWMutex{},
		TableName:   tableName,
		Fields:      map[string]Fields{},
		PrimaryData: map[string]interface{}{},
		ColumnData:  map[string]interface{}{},
		IsLoading:  false,
	}
	return table
}

type Table struct {
	sync.RWMutex
	TableName    string
	Fields       map[string]Fields
	PrimaryData  DataMap
	ColumnData   DataMap
	IsLoading      bool
}

func (this *Table) AddColumn( data []Fields )  {
	this.Lock()
	defer this.Unlock()

	for i:=0;i<len(data);i++{
		this.Fields[data[i].FieldName] = data[i]
	}
}

func (this *Table) PreColumn( data DataMap ) DataMap  {
	out:=DataMap{}
	for k,item:=range data{
		if col,ok:=this.Fields[k];ok{
			switch col.FieldType {
			case FieldType_DATETIME:
				out[col.FieldName] = item.(time.Time).Format("2006-01-02 15:04:05")
			//	support.Log(k,item)
			default:
				out[ col.FieldName ] = item
			//	support.Log(k,item)
			}
		}else{
			//support.Log("==> ",k,item)
		}
	}
	return out
}

/**
  一般条件  a= b  and  a=c
 */
func (this *Table) Query () ([]DataMap,error) {
	this.RLock()
	field := []string{}
	for k,_:=range this.Fields {
		field = append(field,k)
	}
	this.RUnlock()
	return SimpleQuery(this.TableName,field,this.PrimaryData)
}

func (this *Table) Record( key interface{} ) (DataMap,error) {
	this.RLock()
	field := []string{}
	//过滤
	primaryKey := DataMap{}
	for k,item:=range this.Fields {
		field = append(field,k)
		if item.IsPrimary{
			primaryKey[k] = key
		}
	}
	this.RUnlock()

	data,err := SimpleQuery(this.TableName,field,primaryKey)

	if err == nil && len(data) ==1 {
		this.ColumnData = data[0]
		this.Lock()
		for k,v:=range this.Fields{
			if v.IsPrimary{
				this.PrimaryData[k] = this.ColumnData[k]
				this.IsLoading = true
			}
		}
		this.Unlock()
		return this.ColumnData,nil
	}else{
		return nil,ERROR_NOT_FOUND_RECORD
	}
}

func (this *Table) Insert( data DataMap ) (int64,error) {
	raw:= this.PreColumn(data)
	id,e:= InsertData(this.TableName,raw)
	if e==nil{
		//保存数据
		this.ColumnData = data
		this.Lock()
		for k,v:=range this.Fields{
			if v.IsPrimary{
				this.PrimaryData[k] = id
				break
			}
		}
		this.Unlock()
	}else{
		support.Log("Insert Error ",e.Error())
	}
	return id ,nil
}

func (this *Table) Update( data DataMap ) (int64,error) {
	raw:= this.PreColumn(data)
	//跟新
	//根据
	if !this.IsLoading{
		for k,v:=range raw {
			if field,ok:= this.Fields[k] ;ok  {
				if field.IsPrimary{
					this.PrimaryData[k] =v
					delete(raw,k)
				}
			}
		}
	}

	changeId,err := UpdateData(this.TableName,raw,this.PrimaryData)
	if err == nil{
		this.Lock()
		for key,item:= range data{
			this.ColumnData[key] = item
		}
		this.Unlock()
	}
	return changeId,err

}

func (this *Table) Delete( key DataMap )  {

}

func (this *Table) Save( data DataMap ) (int64,error)  {
	return 0,nil
}