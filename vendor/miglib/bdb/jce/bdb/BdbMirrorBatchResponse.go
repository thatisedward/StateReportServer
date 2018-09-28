// **********************************************************************
// This file was generated by a TAF parser!
// TAF version 2.0.0 by WSRD Tencent.
// Generated from `BdbComm.jce'
// **********************************************************************

package bdb

import "reflect"
import "tars/protocol/serializer"

type BdbMirrorBatchResponse struct {
	STableName        string
	MpFailIndexReason map[int32]int32
}

func (_obj *BdbMirrorBatchResponse) resetDefault() {
	_obj.STableName = ""
}

func (_obj *BdbMirrorBatchResponse) WriteTo(_os serializer.JceOutputStream) error {
	var _err error
	if _err = _os.Write(reflect.ValueOf(&_obj.STableName), 1); _err != nil {
		return _err
	}
	if _err = _os.Write(reflect.ValueOf(&_obj.MpFailIndexReason), 2); _err != nil {
		return _err
	}
	return nil
}

func (_obj *BdbMirrorBatchResponse) ReadFrom(_is serializer.JceInputStream) error {
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
	_i, _err = _is.Read(reflect.TypeOf(_obj.MpFailIndexReason), 2, true)
	if _err != nil {
		return _err
	}
	if _i != nil {
		_obj.MpFailIndexReason = _i.(map[int32]int32)
	}
	return nil
}

func (_obj *BdbMirrorBatchResponse) Display(_ds serializer.JceDisplayer) {
	_ds.Display(reflect.ValueOf(&_obj.STableName), "sTableName")
	_ds.Display(reflect.ValueOf(&_obj.MpFailIndexReason), "mpFailIndexReason")
}

func (_obj *BdbMirrorBatchResponse) WriteJson(_en serializer.JceJsonEncoder) ([]byte, error) {
	var _err error
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.STableName), "sTableName")
	if _err != nil {
		return nil, _err
	}
	_err = _en.EncodeJSON(reflect.ValueOf(&_obj.MpFailIndexReason), "mpFailIndexReason")
	if _err != nil {
		return nil, _err
	}
	return _en.ToBytes(), nil
}

func (_obj *BdbMirrorBatchResponse) ReadJson(_de serializer.JceJsonDecoder) error {
	return _de.DecodeJSON(reflect.ValueOf(_obj))
}