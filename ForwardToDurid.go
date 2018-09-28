package main

import (
	"net"
	"strconv"
	"time"
)

var Handler ForwardToDurid

type ForwardToDurid struct {
	newUdpListener *net.UDPConn
}

func (p *ForwardToDurid) Init() {

	checkListener, err := Listener.establish()

	if err != nil {
		dayLog.Debug("Failures! Establishing listener failed. ", err)
		return
	}
	p.newUdpListener = checkListener
	dayLog.Debug("Already bound Local IP:", p.newUdpListener.LocalAddr().String())

}

func (p *ForwardToDurid) TransferToDurid() {
	dayLog.Debug("Going into func TransferToDurid.")

	dayLog.Debug("Second check: bound Local IP:", p.newUdpListener.LocalAddr().String())

	data := make([]byte, Configurations.Control.dataBuffer)
	packetChan := make(chan string, Configurations.Control.packetChan)

	go func() {
		for {
			n, _ := Listener.read(p.newUdpListener, &data)
			report.channelConditions(len(packetChan))
			//dayLog.Info("1->",strconv.FormatInt(time.Now().Unix()*1000,10), "|and:", string(data[10:n]))
			packetChan <- strconv.FormatInt(time.Now().Unix()*1000, 10) + string(data[10:n])
			//packetChan <- string(data[:n])
		}
	}()

	dayLog.Debug("Establishing goRoutine base:", Configurations.Control.routineChan)
	for i := 0; i < Configurations.Control.routineChan; i++ {
		go handlingReport(&packetChan)
	}

	/**
	 * It's never gonna leave :)
	 */

	//select {}
}

func (p *ForwardToDurid) Start() {
	go p.TransferToDurid()
}
