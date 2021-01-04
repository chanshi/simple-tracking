package tracking

type Config struct {
	EnableTracking   bool            // 是否开启跟踪 默认不开启
	Dsn              string          // 数据库连接地址
	MaxOpen          int             // 数据库最大连接数
	GeoCity          string          // GeoCity2 city 数据库地址
}
