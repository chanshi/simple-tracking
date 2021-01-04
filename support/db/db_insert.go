package db

import (
	"strings"
)

// build sql
// table
// inert into <table>

type DataMap  map[string]interface{}


func InsertData(table string, data DataMap  ) (int64,error)   {
	if ! Get().IsConnection() {
		return 0,ERROR_NOT_CONNECTION
	}
	fieldK := []string{}
	valueK := []string{}

	for k,_:= range data {
		fieldK = append(fieldK,k)
		valueK = append(valueK,":"+k)
	}
	buildSql :="INSERT INTO "+table+"("+ strings.Join(fieldK,",") +" ) VALUES("+ strings.Join(valueK,",") +" ); "

	result,err :=Get().DB().NamedExec(buildSql,map[string]interface{}(data))
	if err!=nil {
		return 0, err
	}
	return result.LastInsertId()
}
