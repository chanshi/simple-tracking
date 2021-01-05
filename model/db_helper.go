package model

import "github.com/chanshi/simple-tracking/support/db"

/**
CREATE TABLE `visitor` (
  `visitorId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `firstVisitTime` datetime NOT NULL,
  `appName` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `appVersion` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channelName` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceId` char(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceOs` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceOsVersion` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceBrand` char(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` char(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ipCountry` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ipLatitude` double DEFAULT NULL,
  `ipLongitude` double DEFAULT NULL,
  `ipCity` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_0` text COLLATE utf8mb4_unicode_ci,
  `data_1` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_2` int(11) DEFAULT NULL,
  `data_3` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_4` int(11) DEFAULT NULL,
  `lastVisitTime` datetime DEFAULT NULL,
  PRIMARY KEY (`visitorId`),
  KEY `timel` (`visitorId`,`firstVisitTime`,`lastVisitTime`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

*/

func NewVisitorTable() *db.Table {
	table:=db.NewTable("visitor")
	fields:=[]db.Fields{
		{FieldName: "visitorId", IsPrimary:   true},
		{FieldName: "firstVisitTime",FieldType: db.FieldType_DATETIME},
		{FieldName: "lastVisitTime",FieldType: db.FieldType_DATETIME},
		{FieldName: "appName"},
		{FieldName: "appVersion"},
		{FieldName: "channelName"},
		{FieldName: "deviceId"},
		{FieldName: "deviceOs"},
		{FieldName: "deviceOsVersion"},
		{FieldName: "deviceBrand"},
		{FieldName: "ip"},
		{FieldName: "ipCountry"},
		{FieldName: "ipCity"},
		{FieldName: "ipLatitude"},
		{FieldName: "ipLongitude"},
	}
	fields =append(fields,DataSectionField()...)
	table.AddColumn(fields)
	return table
}

func GetVisitorData( visitor *Visitor ) db.DataMap {
	result:= db.DataMap{}
	result["visitorId"]         = visitor.VisitorId
	result["firstVisitTime"]    = visitor.FirstVisitTime
	result["lastVisitTime"]     = visitor.LastVisitTime

	for k,item:= range visitor.Meta.value(){
		result[k] = item
	}
	for k,item:=range visitor.Data.value(){
		result[k] =item
	}

	return result
}

/**
CREATE TABLE `visit` (
  `visitId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `visitorId` int(11) NOT NULL,
  `appName` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `appVersion` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channelName` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `refererPage` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `enterVisitPage` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `visitTime` datetime NOT NULL,
  `lastActiveTime` datetime NOT NULL,
  `exitVisitPage` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceId` char(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceOs` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceOsVersion` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deviceBrand` char(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` char(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ipCountry` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ipLatitude` float DEFAULT NULL,
  `ipLongitude` float DEFAULT NULL,
  `ipCity` char(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_0` text COLLATE utf8mb4_unicode_ci,
  `data_1` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_2` int(11) DEFAULT NULL,
  `data_3` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_4` int(11) DEFAULT NULL,
  PRIMARY KEY (`visitId`),
  KEY `visitTime` (`visitorId`),
  KEY `visitorId` (`visitorId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
 */

func NewVisitTable() *db.Table  {
	table:=db.NewTable("visit")
	fields := []db.Fields{
		{FieldName: "visitId",IsPrimary: true},
		{FieldName: "visitorId"},
		{FieldName: "refererPage"},
		{FieldName: "enterVisitPage"},
		{FieldName: "exitVisitPage"},    //退出页面
		{FieldName: "visitTime",FieldType: db.FieldType_DATETIME},
		{FieldName: "lastActiveTime",FieldType: db.FieldType_DATETIME},
		{FieldName: "appName"},
		{FieldName: "appVersion"},
		{FieldName: "channelName"},
		{FieldName: "deviceId"},
		{FieldName: "deviceOs"},
		{FieldName: "deviceOsVersion"},
		{FieldName: "deviceBrand"},
		{FieldName: "ip"},
		{FieldName: "ipCountry"},
		{FieldName: "ipCity"},
		{FieldName: "ipLatitude"},
		{FieldName: "ipLongitude"},
	}
	fields =append(fields,DataSectionField()...)
	table.AddColumn(fields)
	return table
}

func GetVisitData( visit *Visit ) db.DataMap  {
	result:= db.DataMap{}
	result["visitorId"]      = visit.VisitorId
	result["refererPage"]    = visit.RefererPage
	result["enterVisitPage"] = visit.EnterVisitPage
	result["exitVisitPage"]  = visit.ExitVisitPage
	result["visitTime"]      = visit.VisitTime
	result["lastActiveTime"] = visit.LastActiveTime
	for k,item:= range visit.Meta.value(){
		result[k] = item
	}
	for k,item:=range visit.Data.value(){
		result[k] =item
	}

	return result
}

/**
CREATE TABLE `visit_page` (
  `visitPageId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `visitId` int(11) DEFAULT NULL,
  `visitorId` int(11) DEFAULT NULL,
  `refererPage` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `visitPage` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `visitTime` datetime DEFAULT NULL,
  `visitSpent` int(11) DEFAULT NULL,
  `data_0` text COLLATE utf8mb4_unicode_ci,
  `data_1` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_2` int(11) DEFAULT NULL,
  `data_3` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_4` int(11) DEFAULT NULL,
  PRIMARY KEY (`visitPageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

 */

func NewVisitPageTable() *db.Table  {
	table:= db.NewTable("visit_page")
	fields := []db.Fields{
		{FieldName: "visitPageId",IsPrimary: true},
		{FieldName: "visitId"},
		{FieldName: "visitorId"},
		{FieldName: "refererPage"},
		{FieldName: "visitPage"},
		{FieldName: "visitTime",FieldType: db.FieldType_DATETIME},
		{FieldName: "visitSpent"},
		{FieldName: "refresh"},
	}
	fields = append(fields,DataSectionField()...)
	table.AddColumn(fields)
	return table
}
func GetVisitPageData( page *VisitPage ) db.DataMap  {
	result:= db.DataMap{}
	result["visitPageId"]    = page.VisitPageId
	result["visitId"]        = page.VisitId
	result["visitorId"]      = page.VisitorId

	result["refererPage"]    = page.RefererPage
	result["visitPage"]      = page.VisitPage
	result["visitTime"]      = page.VisitTime
	result["visitSpent"]     = page.VisitSpent
	result["refresh"]        = page.Refresh

	for k,item:=range page.Data.value(){
		result[k] =item
	}
	return result
}

/**
CREATE TABLE `visit_action` (
  `actionId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `visitId` int(11) DEFAULT NULL,
  `visitorId` int(11) DEFAULT NULL,
  `visitPageId` int(11) DEFAULT NULL,
  `actionGroup` char(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `actionName` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `actionTime` datetime DEFAULT NULL,
  `data_0` text COLLATE utf8mb4_unicode_ci,
  `data_1` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_2` int(11) DEFAULT NULL,
  `data_3` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_4` int(11) DEFAULT NULL,
  PRIMARY KEY (`actionId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
 */

func NewVisitActionTable() *db.Table  {
	table:= db.NewTable("visit_action")
	fields:= []db.Fields{
		{FieldName: "actionId",IsPrimary: true},
		{FieldName: "visitId"},
		{FieldName: "visitorId"},
		{FieldName: "visitPageId"},

		{FieldName: "actionGroup"},
		{FieldName: "actionName"},
		{FieldName: "actionTime",FieldType: db.FieldType_DATETIME},
	}
	fields = append(fields,DataSectionField()...)
	table.AddColumn(fields)
	return table
}
func GetVisitActionData( action *VisitAction ) db.DataMap {
	result:= db.DataMap{}
	result["actionId"]       = action.ActionId
	result["visitPageId"]    = action.VisitPageId
	result["visitId"]        = action.VisitId
	result["visitorId"]      = action.VisitorId

	result["actionGroup"]    = action.ActionGroup
	result["actionName"]      = action.ActionName
	result["actionTime"]      = action.ActionTime

	for k,item:=range action.Data.value(){
		result[k] =item
	}
	return result
}
/**
CREATE TABLE `visit_event` (
  `eventId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `visitId` int(11) DEFAULT NULL,
  `visitorId` int(11) DEFAULT NULL,
  `eventName` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `eventTime` datetime DEFAULT NULL,
  `data_0` text COLLATE utf8mb4_unicode_ci,
  `data_1` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_2` int(11) DEFAULT NULL,
  `data_3` varchar(300) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `data_4` int(11) DEFAULT NULL,
  PRIMARY KEY (`eventId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
 */

func NewVisitEventTable() *db.Table  {
	table:= db.NewTable("visit_event")
	fields := []db.Fields{
		{FieldName: "eventId",IsPrimary: true},
		{FieldName: "visitId"},
		{FieldName: "visitorId"},

		{FieldName: "eventName"},
		{FieldName: "eventTime",FieldType: db.FieldType_DATETIME},
	}
	fields=append(fields,DataSectionField()...)
	table.AddColumn(fields)
	return table
}

func GetVisitEventData( event *VisitEvent ) db.DataMap  {
	result:= db.DataMap{}
	result["eventId"]         = event.EventId
	result["visitId"]         = event.VisitId
	result["visitorId"]       = event.VisitorId

	result["eventName"]       = event.EventName
	result["eventTime"]       = event.EventTime

	for k,item:= range event.Data.value(){
		result[k] =item
	}
	return result
}

