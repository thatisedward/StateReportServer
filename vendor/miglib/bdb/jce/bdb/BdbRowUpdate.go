// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbRowUpdate struct {
	MField     map[string]BdbFieldUpdate
	MFieldCond [][]BdbFieldCondition
	IVersion   int64
}

func (_obj *BdbRowUpdate) resetDefault() {
	_obj.IVersion = 0
}

func (_obj *BdbRowUpdate) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.MField), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MFieldCond), 2); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.IVersion), 3); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbRowUpdate) ReadFrom(_is serializer.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.MField), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MField = _i.(map[string]BdbFieldUpdate)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.MFieldCond), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MFieldCond = _i.([][]BdbFieldCondition)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.IVersion), 3, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.IVersion = _i.(int64)
	}
	return nil
}

func (_obj *BdbRowUpdate) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.MField), "mField")
	_ds.Display(reflect.ValueOf(&_obj.MFieldCond), "mFieldCond")
	_ds.Display(reflect.ValueOf(&_obj.IVersion), "iVersion")
}

func (_obj *BdbRowUpdate) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MField), "mField")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MFieldCond), "mFieldCond")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.IVersion), "iVersion")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbRowUpdate) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}
