package support

var conf  *TrackingConfig
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
func init()  {
	conf = &TrackingConfig{
		EnableTracking:  true,
		EnableTrackingPage: true,
		SessionTimeOut: 1800,
		DSN: "",
		MaxOpen: 500,
		GeoCity: "",
	}
}

func Config() *TrackingConfig { return conf }
func SetConfig( data *TrackingConfig )  { conf = data }


/**
  需要用到再说
 */
type TrackingConfig struct {

	EnableTracking        bool       `json:"enableTracking"`        // 是否开启跟踪
	EnableTrackingPage    bool       `json:"enableTrackingPage"`    // 是否开启跟踪页面
	SessionTimeOut        int        `json:"sessionTimeOut"`        // 会话时长  30分钟  1800

	DSN                   string     `json:"dsn"`                   // mysql DSN 信息
	MaxOpen               int        `json:"maxOpen"`
	GeoCity               string     `json:"geoCity"`
}


func (this *TrackingConfig) IsTracking() bool {
	return this.EnableTracking == true
}

func (this *TrackingConfig) SetEnableTracking( enable bool )  {
	this.EnableTracking = enable
}

func (this *TrackingConfig)SetDsn( dsn string )  {
	this.DSN = dsn
}