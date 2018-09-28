package dcache

import (
	"miglib/dcache/jce/DCache"
	. "tars"
)

type DCachev1 struct {
	comm  *Communicator
	proxy *DCache.Proxy
}

func (d *DCachev1) setObj(objName string) {
	d.comm = NewCommunicator()
	d.proxy = new(DCache.Proxy)
	d.comm.StringToProxy(objName, d.proxy)
}

func (d *DCachev1) GetProxy(objName string) *DCache.Proxy {
	d.setObj(objName)
	return d.proxy
}
