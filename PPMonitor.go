package main

import (
	"miglib/propertyplus"
	"strconv"
)

const (
	POST_HTTP_ERROR = iota
	POST_PARCEL_EMPTY
	DURID_PARSE_ERROR
	POST_RETRY
	POST_FAILED
)

const SERVER = "NewsAD.StateReportServer"

var report PPMonitor

var reportError *propertyplus.PropertyMsg
var reportPacket *propertyplus.PropertyMsg
var reportChanCond *propertyplus.PropertyMsg

type PPMonitor struct {
}

func (p *PPMonitor) Init() {
	reportError = propertyplus.CreatePropertyPlus("Errors", SERVER)
	reportPacket = propertyplus.CreatePropertyPlus("Event", SERVER)
	reportChanCond = propertyplus.CreatePropertyPlus("PacketChannel", SERVER)
}

func (p *PPMonitor) errors(errorCode int) {
	quantity := []float32{}

	switch errorCode {

	case 0:
		quantity = []float32{float32(1), float32(1), float32(0), float32(0), float32(0), float32(0)}
	case 1:
		quantity = []float32{float32(1), float32(0), float32(1), float32(0), float32(0), float32(0)}
	case 2:
		quantity = []float32{float32(1), float32(0), float32(0), float32(1), float32(0), float32(0)}
	case 3:
		quantity = []float32{float32(1), float32(0), float32(0), float32(0), float32(1), float32(0)}
	case 4:
		quantity = []float32{float32(1), float32(0), float32(0), float32(0), float32(0), float32(1)}
	default:
		dayLog.Error("PLEASE CHECK! Invalid Error Code ->", errorCode)
		return
	}

	errorType := []string{strconv.FormatInt(int64(errorCode), 10)}

	reportError.Report(errorType, quantity)
}

func (p *PPMonitor) packet(pRRS int) {

	if pRRS <= 0 {
		dayLog.Error("PLEASE CHECK! Invalid postResp.Result.Sent value", pRRS)
		return
	}

	filterTypes := []string{}
	oneParcel := []float32{float32(1)}

	reportPacket.Report(filterTypes, oneParcel)
}

func (p *PPMonitor) channelConditions(len int) {
	settedLen := float32(Configurations.Control.packetChan)
	addOne := []float32{}
	switch {
	case len < int(0.25*settedLen):
		addOne = []float32{float32(1), float32(0), float32(0), float32(0)}
	case len < int(0.5*settedLen):
		addOne = []float32{float32(0), float32(1), float32(0), float32(0)}
	case len < int(0.75*settedLen):
		addOne = []float32{float32(0), float32(0), float32(1), float32(0)}
	default:
		addOne = []float32{float32(0), float32(0), float32(0), float32(1)}
	}
	onceFull := []string{"PacketChannel"}

	reportChanCond.Report(onceFull, addOne)
}
