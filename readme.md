#### Simple Tracking

使用 go 语言，简单开发的用户访问跟踪系统。包括用户的行为和事件的跟踪。

### 安装

```shell script
go get -t github.com/chanshi/simple-tracking 
```


### 使用方式

配置
```go
type Config struct {
	EnableTracking   bool           `json:"enableTracking"`  // 是否开启跟踪 默认不开启
	Dsn              string         `json:"dsn"`             // 数据库连接地址
	MaxOpen          int            `json:"maxOpen"`         // 数据库最大连接数
	GeoCity          string         `json:"geoCity"`         // GeoCity2 city 数据库文件地址
}
```

```go

// 1. 配置

// 2. 初始化服务
Track().Config(&Config{EnableTracking: true,Dsn: ""}).Init()

// 3. 创建跟踪者
tracker := Track().
          Visitor(19).
	      App("bate","v1.2.1").
		  Channel("Local").
		  Ip("223.104.103.23")

tracker.Device("*&*^2","","","")

// 4. 跟踪开始
tracker.Begin()

// 5. 跟踪页面  5个用户自定义参数
tracker.Page("home")

// 6. 跟踪页面上的行为 5个用户自定义参数
tracker.Action("game","enter","800","437621")

// 7. 跟踪事件  5个用户自定义参数
tracker.Event("eventName")

// 8. 跟踪结束 
tracker.End()

// 9. 推出跟踪
Track().ExitVisitor(19)

// 10. 结束服务
Track().Stop()
```

#### 模型

### Visit 访问模型
  访问，访客，设备信息，来源信息。  
  
  可以统计  
  1. 基于用户
   * 新赠用户
   * 用户次日留存，7日 15日 30日
   * 3个月后再次登录，唤醒。
   * 用户地理位置分析
   * 用户设备信息分析
   
  2. app 信息
   * app 版本号
   * app 使用设备号
   * app 设备长宽
   * app 设备是否启用麦克风/摄像机 等
  
  3. 来源信息
   * 用户是从哪个渠道唤醒的
   * 用户是从哪里下载的
   * 用户因为什么打开了APP 
   * 分析来源带来的用户量，以及转化率 ，留存率的分析。得出优质的来源或者渠道 
  
  4. 访问之间的关系
   * 可以计算用户的粘性
   
  5. 基本页面，行为，事件统计
   * 可以页面的访问深度。
   * 平均页面深度
   * 退出页面
   * 以及退出页面的退出率



