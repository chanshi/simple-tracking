package db

import (
	"strings"
)

func UpdateData( table string , saveData ,whereDate DataMap ) (int64,error) {
	if ! Get().IsConnection() {
		return 0,ERROR_NOT_CONNECTION
	}

	fieldK:=[]string{}
	for k,_:=range saveData {
		fieldK= append(fieldK,k+"=:"+k)
	}
	whereK:=[]string{}
	mergeK:=[]string{}
	for k,_:=range whereDate {
		whereK = append(whereK,k+"=:"+k )
		mergeK =append(mergeK,k)
	}

	buildSql :="UPDATE "+ table +" SET "+strings.Join(fieldK,",")+" WHERE "+strings.Join(whereK,",")+ ";"
	//support.Log("UPDATE SQL==> ",buildSql)

	for i:=0;i<len(mergeK);i++ {
		key:=mergeK[i]
		saveData[key] = whereDate[key]
	}

	result,err :=Get().DB().NamedExec(buildSql,map[string]interface{}(saveData) )
	if err!=nil {
		return 0, err
	}

	return result.RowsAffected()
}