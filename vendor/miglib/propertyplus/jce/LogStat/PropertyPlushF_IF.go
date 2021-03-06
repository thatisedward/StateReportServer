// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `PropertyPlusF.jce'
// **********************************************************************

package LogStat

import (
	"errors"
	"fmt"
	"reflect"
	m "tars/model"
	"tars/protocol/res/requestf"
	"tars/protocol/serializer"
)

type PropertyPlushF struct {
	s m.Servant
}

func (_obj *PropertyPlushF) Logstat(log StatLog, _opt ...map[string]string) (_ret int32, _err error) {
	_oe := serializer.NewOutputStream()
	_oe.Write(reflect.ValueOf(&log), 1)
	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err := _obj.s.Taf_invoke(0, "logstat", _oe.ToBytes(), _status, _context, _resp)
	if err != nil {
		return _ret, err
	}
	_is := serializer.NewInputStream(_resp.SBuffer)
	r0, err := _is.Read(reflect.TypeOf(_ret), 0, true)
	if err != nil {
		return _ret, err
	}
	return r0.(int32), nil
}
func (_obj *PropertyPlushF) Mutillogstat(logs []StatLog, _opt ...map[string]string) (_ret int32, _err error) {
	fmt.Println("servant info: ", _obj.s)
	_oe := serializer.NewOutputStream()
	_oe.Write(reflect.ValueOf(&logs), 1)
	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err := _obj.s.Taf_invoke(0, "mutillogstat", _oe.ToBytes(), _status, _context, _resp)
	if err != nil {
		return _ret, err
	}
	_is := serializer.NewInputStream(_resp.SBuffer)
	r0, err := _is.Read(reflect.TypeOf(_ret), 0, true)
	if err != nil {
		return _ret, err
	}
	return r0.(int32), nil
}
func (_obj *PropertyPlushF) SetServant(s m.Servant) {
	fmt.Println("Set servant: ", s)
	_obj.s = s
}

type _impPropertyPlushF interface {
	Logstat(log StatLog) (int32, error)
	Mutillogstat(logs []StatLog) (int32, error)
}

func (_obj *PropertyPlushF) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	parms := serializer.NewInputStream(req.SBuffer)
	oe := serializer.NewOutputStream()
	_imp := _val.(_impPropertyPlushF)
	switch req.SFuncName {
	case "logstat":
		var p_0 StatLog
		t_0, err := parms.Read(reflect.TypeOf(p_0), 1, true)
		if err != nil {
			return err
		}
		_ret, err := _imp.Logstat(t_0.(StatLog))
		if err != nil {
			return err
		}
		oe.Write(reflect.ValueOf(&_ret), 0)
	case "mutillogstat":
		var p_0 []StatLog
		t_0, err := parms.Read(reflect.TypeOf(p_0), 1, true)
		if err != nil {
			return err
		}
		_ret, err := _imp.Mutillogstat(t_0.([]StatLog))
		if err != nil {
			return err
		}
		oe.Write(reflect.ValueOf(&_ret), 0)
	default:
		return errors.New("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      oe.ToBytes(),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	return nil
}
