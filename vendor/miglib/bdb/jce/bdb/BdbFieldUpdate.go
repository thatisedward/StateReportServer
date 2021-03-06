// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbFieldUpdate struct {
	SValue []byte
	Oper   int
}

func (_obj *BdbFieldUpdate) resetDefault() {
}

func (_obj *BdbFieldUpdate) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.SValue), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.Oper), 2); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbFieldUpdate) ReadFrom(_is serializer.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.SValue), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.SValue = _i.([]byte)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.Oper), 2, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.Oper = _i.(int)
	}
	return nil
}

func (_obj *BdbFieldUpdate) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.SValue), "sValue")
	_ds.Display(reflect.ValueOf(&_obj.Oper), "oper")
}

func (_obj *BdbFieldUpdate) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.SValue), "sValue")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.Oper), "oper")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbFieldUpdate) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}
