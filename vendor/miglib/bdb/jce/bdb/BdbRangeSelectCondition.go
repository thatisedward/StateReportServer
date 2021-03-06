// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbRangeSelectCondition struct {
	STableName      string
	MSelectStartKey map[string]BdbFieldValue
	MSelectEndKey   map[string]BdbFieldValue
	IBorderFlag     int32
	ILimit          int32
}

func (_obj *BdbRangeSelectCondition) resetDefault() {
	_obj.STableName = ""
	_obj.IBorderFlag = 0
	_obj.ILimit = -1
}

func (_obj *BdbRangeSelectCondition) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.STableName), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MSelectStartKey), 2); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MSelectEndKey), 3); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.IBorderFlag), 4); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.ILimit), 5); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbRangeSelectCondition) ReadFrom(_is serializer.JceInputStream) error {
	var _err error
	var _i interface{}
	_obj.resetDefault()
	_i, _err = _is.Read(reflect.TypeOf(_obj.STableName), 1, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.STableName = _i.(string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.MSelectStartKey), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MSelectStartKey = _i.(map[string]BdbFieldValue)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.MSelectEndKey), 3, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MSelectEndKey = _i.(map[string]BdbFieldValue)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.IBorderFlag), 4, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.IBorderFlag = _i.(int32)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.ILimit), 5, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.ILimit = _i.(int32)
	}
	return nil
}

func (_obj *BdbRangeSelectCondition) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.STableName), "sTableName")
	_ds.Display(reflect.ValueOf(&_obj.MSelectStartKey), "mSelectStartKey")
	_ds.Display(reflect.ValueOf(&_obj.MSelectEndKey), "mSelectEndKey")
	_ds.Display(reflect.ValueOf(&_obj.IBorderFlag), "iBorderFlag")
	_ds.Display(reflect.ValueOf(&_obj.ILimit), "iLimit")
}

func (_obj *BdbRangeSelectCondition) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.STableName), "sTableName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MSelectStartKey), "mSelectStartKey")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MSelectEndKey), "mSelectEndKey")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.IBorderFlag), "iBorderFlag")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.ILimit), "iLimit")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbRangeSelectCondition) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}
