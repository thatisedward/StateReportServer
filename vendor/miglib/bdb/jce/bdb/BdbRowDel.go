// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbRowDel struct {
	MDeleteKey [][]BdbFieldCondition
}

func (_obj *BdbRowDel) resetDefault() {
}

func (_obj *BdbRowDel) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.MDeleteKey), 1); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbRowDel) ReadFrom(_is serializer.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.MDeleteKey), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MDeleteKey = _i.([][]BdbFieldCondition)
	}
	return nil
}

func (_obj *BdbRowDel) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.MDeleteKey), "mDeleteKey")
}

func (_obj *BdbRowDel) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MDeleteKey), "mDeleteKey")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbRowDel) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}