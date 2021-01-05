package support

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

// @SEE https://github.com/oschwald/geoip2-golang

const GeoLite2_Country_file =  "geoLite2/GeoLite2-Country.mmdb"
//const GeoLite2_City_file    =  "geoLite2/GeoLite2-City.mmdb"
const GeoLite2_City_file = "geoLite2/GeoLite2-City.mmdb"


var geoLite2 *geoLite

func init()  {
	geoLite2 = &geoLite{}
	geoLite2.Init()
}

func GetGeo() *geoLite{
	return geoLite2
}

type geoLite struct {
	dbCity      *geoip2.Reader
}

func (this *geoLite) Init()   {
	db,err := geoip2.Open(Config().GeoCity)
	if err!= nil{
		Log("GetGeoIp Open error",err.Error(),Config().GeoCity)
		return
	}
	this.dbCity = db
}

func (this *geoLite) GetCity( ip string) *geoip2.City  {
	defer func() {
		if e:=recover();e!=nil{
			fmt.Println("GEO SET ERROR ",e)
			return
		}
	}()
	if this.dbCity != nil{
		record,_:= this.dbCity.City( net.ParseIP(ip) )
		return record
	}
	return nil
}

func (this *geoLite) Close()  {
	if this.dbCity !=nil{
		if err:=this.dbCity.Close();err!=nil{
			Log("GeoIp Closed error ",err.Error())
		}else{
			Log("GeoIp Closed Success")
		}
	}
}


