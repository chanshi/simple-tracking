### Simple Tracking

简单的数据跟踪统计服务

#### 安装

```shell script
go get -t github.com/chanshi/simple-tracking 
```


### 使用方式
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
* Visit
  访问，访客，设备信息，来源信息。
  一次访问，默认过期时间为 30分钟。
  可以统计
  1. 基于用户
   新赠用户
   用户次日留存，7日 15日 30日
   3个月后再次登录，唤醒。
   用户地理位置分析
   用户设备信息分析
   
  2. app 信息
   app 版本号
   app 使用设备号
   app 设备长宽
   app 设备是否启用麦克风/摄像机 等
  
  3. 来源信息
   用户是从哪个渠道唤醒的
   用户是从哪里下载的
   用户因为什么打开了APP 
   分析来源带来的用户量，以及转化率 ，留存率的分析。得出优质的来源或者渠道 
  
  4. 访问之间的关系
   可以计算用户的粘性
   
  5. 基本页面，行为，事件统计
   可以页面的访问深度。
   平均页面深度
   退出页面
   以及退出页面的退出率

*  Page
  应用内的页面  可以分为 3 个层次
  应该的逻辑为 第一级 > 第二级 > 第三级
  页面 可以统计 
  PV 
  PUV
  

* Action

* Event 

* Visitor


