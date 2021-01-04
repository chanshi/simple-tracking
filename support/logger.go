package support

import (
	"fmt"
	"time"
)

func Log(data ...interface{})  {
	result:=[]interface{}{}
	tm    := time.Now().Format("2006-01-02 15:04:05")
	space := ""
	//tag   := "simple-tracking =>"
	tag:=""

	result = append(result,tm,space,tag)
	for i:=0;i<len(data);i++{
		result = append(result,data[i])
	}
	fmt.Println(result...)
}
