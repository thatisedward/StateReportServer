package propertyplus

import (
	ls "miglib/propertyplus/jce/LogStat"
	"sync"
	ts "tars"
	"tars/util/logger"
	"time"
)

var singlePP *PropertyPlus
var lock *sync.Mutex = &sync.Mutex{}

var TLOG = logger.GetLogger()

const (
	ppObj = "WSDITIL.LogAllotServer4Comm.LogAllotObj"
)

type PropertyPlus struct {
	mlock      *sync.Mutex
	comm       *ts.Communicator
	ppf        *ls.PropertyPlushF
	ppmsg      map[string]*PropertyMsg
	serverName string
}

func CreatePropertyPlus(propertyName string, serverName string) *PropertyMsg {
	if serverName == "" || propertyName == "" {
		TLOG.Error("server name or propertyname is empty", serverName, propertyName)
		return nil
	}

	if singlePP == nil {
		lock.Lock()
		defer lock.Unlock()
		if singlePP == nil {
			singlePP = &PropertyPlus{}
			singlePP.init()
		}
	}

	return singlePP.CreatePP(propertyName, serverName)
}

func (pp *PropertyPlus) CreatePP(propertyName string, serverName string) *PropertyMsg {

	realName := "pp_" + serverName + "_" + propertyName

	pp.mlock.Lock()
	defer pp.mlock.Unlock()

	ppmsg, ok := pp.ppmsg[realName]
	if !ok {
		ppmsg = &PropertyMsg{}
		ppmsg.init()
		pp.ppmsg[realName] = ppmsg
	}

	return ppmsg
}

func (pp *PropertyPlus) init() {
	pp.comm = ts.NewCommunicator()
	pp.ppf = new(ls.PropertyPlushF)
	pp.comm.StringToProxy(ppObj, pp.ppf)
	pp.ppmsg = make(map[string]*PropertyMsg)
	pp.mlock = new(sync.Mutex)
	go pp.loop()
}

func (pp *PropertyPlus) doReport() {

	pp.mlock.Lock()
	defer pp.mlock.Unlock()

	var data []ls.StatLog
	for k, v := range pp.ppmsg {
		stat := v.doReport(k)
		if len(stat.Content) != 0 {
			data = append(data, stat)
		}
	}

	if len(data) != 0 {
		ret, err := pp.ppf.Mutillogstat(data)
		if ret != 0 {
			TLOG.Error("failed to report pp info, ", err)
		}
	}
}

func (pp *PropertyPlus) loop() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			pp.doReport()
		}
	}
}
