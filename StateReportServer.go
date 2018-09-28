package main

import (
	"tars"

	"NewsAD"
)

var dayLog = tars.GetDayLogger("log.DLOG", 1)

func init() {
	dayLog.Debug("Initializing StateReportServer")
	imp := new(StateReportImp)     //New Imp
	app := new(NewsAD.StateReport) //New init the A JCE
	cfg := tars.GetServerConfig()  //Get Config File Object

	dayLog.Debug(imp, app)
	findPort := cfg.Adapters["NewsAD.StateReportServer.StateReportObjAdapter"].Endpoint.Port
	//app.AddServant(imp, cfg.App+"."+cfg.Server+".StateReportObj") //Register Servant
	report.Init()
	Configurations.Init(cfg.LocalIP, int(findPort))
	Handler.Init()
}
func main() { //Init servant

	dayLog.Debug("StateReportServer starts running!")
	Handler.Start()
	tars.Run()

}
