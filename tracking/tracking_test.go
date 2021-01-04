package tracking

import "testing"

func TestTracking_Start(t *testing.T) {

	Track().Config(&Config{EnableTracking: true,Dsn: ""}).Init()

	tracker := Track().Visitor(19).
		App("bate","v1.2.1").
		Channel("Local")

	tracker.Device("*&*^2","","","")
	tracker.Begin()

	tracker.Page("home")
	tracker.Page("home")
	tracker.Page("home")
	tracker.Page("home")
	tracker.Page("home")
	tracker.Page("mine")
	tracker.Page("game")
	tracker.Action("game","enter","800","437621")
	tracker.Action("game","exit","800","437621")
	tracker.Page("live")
	tracker.Page("liveRoom","","1283")
	tracker.Action("gift","give","","i2hs")
	tracker.End()

	Track().ExitVisitor(19)

	 Track().Stop()

}
