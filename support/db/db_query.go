package db

import (
	"errors"
	"github.com/chanshi/simple-tracking/support"
	"strings"
)

/**
SELECT
FROM
WHERE
GROUP BY
ORDER BY
LIMIT


 */


var ERROR_NOT_FOUND_RECORD = errors.New("NOT FOUND RECORD")

func SimpleQuery( table string , field []string ,whereMap DataMap ) ( []DataMap, error ){
	if ! Get().IsConnection() {
		return nil,ERROR_NOT_CONNECTION
	}

	fieldRaw := "*"
	if len(field) > 0 {
		fieldRaw = strings.Join(field,",")
	}
	whereK:=[]string{}
	for k,_:=range whereMap {
		whereK = append(whereK,k+"=:"+k )
	}

	buildSql := "SELECT "+fieldRaw+" FROM "+table+" WHERE "+strings.Join(whereK," AND ")+" LIMIT 100;"
	result := []DataMap{}

	row,err:=Get().DB().NamedQuery( buildSql,map[string]interface{}(whereMap) )
	if err !=nil{
		return result, err
	}

	for row.Next() {
		item:= map[string]interface{}{}
		err:=row.MapScan(item)
		if err!=nil{
			support.Log("SQL:[ "+ buildSql +"] \n ",err.Error())
		}else{
			result = append(result,item)
		}
	}
	return result,nil

}