package tracking

type Config struct {
	EnableTracking   bool           `json:"enableTracking"` // 是否开启跟踪 默认不开启
	Dsn              string         `json:"dsn"` // 数据库连接地址
	MaxOpen          int            `json:"maxOpen"` // 数据库最大连接数
	GeoCity          string         `json:"geoCity"`// GeoCity2 city 数据库地址
}
