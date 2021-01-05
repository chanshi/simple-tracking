package model

import "github.com/chanshi/simple-tracking/support/db"

type MetaApp struct {
	AppName        string
	AppVersion     string
}

func (this MetaApp)  value() db.DataMap {
	v := make(db.DataMap)
	v["appName"]    = this.AppName
	v["appVersion"] = this.AppVersion
	return v
}

