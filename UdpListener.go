package main

import (
	"net"
	"tars"
)

var Listener UdpListener

type UdpListener struct {
}

func (p *UdpListener) establish() (*net.UDPConn, error) {

	address := &net.UDPAddr{IP: net.ParseIP(Configurations.Server.ip), Port: Configurations.Server.port}
	//listener, error := net.ListenUDP("udp", address)
	return net.ListenUDP("udp", address)
}

func (p *UdpListener) read(listener *net.UDPConn, data *[]byte) (int, error) {

	n, remoteAddr, error := listener.ReadFromUDP(*data)

	if error != nil {
		dayLog.Error("Failures! Reading data failed. errors: ", error, "n: ", n, "|remoteAddress: ", remoteAddr)
		return 0, error
	}
	//fmt.Printf("receive: %s, %s\n", data[:n], remoteAddr)
	tars.TLOG.Info("receiving:", string((*data)[:n]), "remote Address: ", remoteAddr)
	return n, error
}
