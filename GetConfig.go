package main

import (
	"IpInfo"
	"io/ioutil"
	"tars"
	"tars/util/conf"
	"tars/util/rconf"
)

var Configurations GetConfig

type GetConfig struct {
	Durid struct {
		url string
	}
	Server struct {
		ip   string
		port int
	}
	Control struct {
		packetChan     int
		routineChan    int
		dataBuffer     int
		postRetry      int
		parcelCapacity int
	}
}

func (c *GetConfig) Init(localIp string, port int) {

	c.parseConfig("server.conf", localIp, port)
}

func (c *GetConfig) getConf(file string) string {

	if false {
		data, _ := ioutil.ReadFile(file)
		return string(data)
	}
	cfg := tars.GetServerConfig()
	tc := rconf.NewRConf(cfg.App, cfg.Server, cfg.BasePath)
	data, err := tc.GetConfig(file)
	if err != nil {
		dayLog.Error("config error:", err)
		panic(err)
	}

	return data
}

func (c *GetConfig) parseConfig(config, localIp string, port int) {

	var data string

	data = c.getConf(config)

	dayLog.Debug("config:\n" + data)

	cfg := conf.New()
	cfg.InitFromString(data)

	c.Durid.url = cfg.GetString("/Main/Durid/<url>")

	//c.Server.port = cfg.GetInt("/Main/Server/<port>")
	c.Server.port = port
	//c.Server.ip = c.getUdpIp()
	c.Server.ip = localIp
	//dayLog.Debug("c.Server.ip", c.Server.ip)
	c.Control.packetChan = cfg.GetInt("/Main/Control/<packetChan>")
	c.Control.routineChan = cfg.GetInt("/Main/Control/<routineChan>")
	c.Control.dataBuffer = cfg.GetInt("/Main/Control/<dataBuf>")
	c.Control.postRetry = cfg.GetInt("/Main/Control/<postRetry>")
	c.Control.parcelCapacity = cfg.GetInt("/Main/Control/<parcelCap>")

}

func (c *GetConfig) getUdpIp() string {
	udpIp := IpInfo.GetIntranetIp()

	if udpIp == "" {
		dayLog.Error("Failed to find UDP IP of current server.")
		return udpIp
	}

	dayLog.Error("The UDP IP is:", udpIp)
	return udpIp
}
