package model

import (
	"fmt"
	"github.com/chanshi/simple-tracking/support"
)

/**
  可以根据  GEO  来解析
 */
type MetaLocation struct {
	Ip              string                  // IP 地址
	IpCountry       string                  // 国家        //时区？
	IpCity          string
	IpLatitude      float64                 // 经纬度
	IpLongitude     float64                 // 经纬度
}

func (this MetaLocation) value() MetaArray  {
	defer func() {
		if e:=recover();e!=nil{
			fmt.Println("GEO SET ERROR ",e)
			return
		}
	}()
	if this.Ip != "" && this.IpCountry == "" && support.GetGeo() !=nil {
		 city:= support.GetGeo().GetCity(this.Ip)
		 this.IpCountry   = city.Country.Names["en"]
		 this.IpCity      = city.City.Names["en"]
		 this.IpLongitude = city.Location.Longitude
		 this.IpLatitude  = city.Location.Latitude
	}
	value := make(MetaArray)
	value["ip"]          =  this.Ip
	value["ipCountry"]   =  this.IpCountry
	value["ipCity"]      =  this.IpCity
	value["ipLatitude"]  =  this.IpLatitude
	value["ipLongitude"] = this.IpLongitude
	return value
}