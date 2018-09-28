// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbSelectCondition struct {
	STableName     string
	MSelectKey     [][]BdbFieldCondition
	ILimit         int32
	VFields        []string
	BReadBaseLocal bool
}

func (_obj *BdbSelectCondition) resetDefault() {
	_obj.STableName = ""
	_obj.ILimit = -1
	_obj.BReadBaseLocal = false
}

func (_obj *BdbSelectCondition) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.STableName), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MSelectKey), 2); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.ILimit), 3); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.VFields), 4); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.BReadBaseLocal), 5); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbSelectCondition) ReadFrom(_is serializer.JceInputStream) error {
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
	_i, _err = _is.Read(reflect.TypeOf(_obj.MSelectKey), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MSelectKey = _i.([][]BdbFieldCondition)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.ILimit), 3, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.ILimit = _i.(int32)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.VFields), 4, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.VFields = _i.([]string)
	}
	_i, _err = _is.Read(reflect.TypeOf(_obj.BReadBaseLocal), 5, false)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.BReadBaseLocal = _i.(bool)
	}
	return nil
}

func (_obj *BdbSelectCondition) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.STableName), "sTableName")
	_ds.Display(reflect.ValueOf(&_obj.MSelectKey), "mSelectKey")
	_ds.Display(reflect.ValueOf(&_obj.ILimit), "iLimit")
	_ds.Display(reflect.ValueOf(&_obj.VFields), "vFields")
	_ds.Display(reflect.ValueOf(&_obj.BReadBaseLocal), "bReadBaseLocal")
}

func (_obj *BdbSelectCondition) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.STableName), "sTableName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MSelectKey), "mSelectKey")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.ILimit), "iLimit")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.VFields), "vFields")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.BReadBaseLocal), "bReadBaseLocal")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbSelectCondition) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}