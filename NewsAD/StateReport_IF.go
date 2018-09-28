//
// This file war generated by FastJce2go 1.0
// Generated from StateReport.jce
// Tencent.

package NewsAD

import (
	"fmt"
	"tars"
	m "tars/model"
	"tars/protocol/codec"
	"tars/protocol/res/requestf"
)

type StateReport struct {
	s m.Servant
}

func (_obj *StateReport) Test(A int32, B int32, C *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(A, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(B, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Taf_invoke(0, "Test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*C), 3, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}
func (_obj *StateReport) TransferToDurid(Num int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(Num, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Taf_invoke(0, "TransferToDurid", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(_resp.SBuffer)
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_ = length
	_ = have
	_ = ty
	return ret, nil
}

func (_obj *StateReport) SetServant(s m.Servant) {
	_obj.s = s
}

func (_obj *StateReport) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

func (_obj *StateReport) AddServant(imp _impStateReport, obj string) {
	tars.AddServant(_obj, imp, obj)
}

type _impStateReport interface {
	Test(A int32, B int32, C *int32) (ret int32, err error)
	TransferToDurid(Num int32) (ret int32, err error)
}

func (_obj *StateReport) Dispatch(_val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(req.SBuffer)
	_os := codec.NewBuffer()
	_imp := _val.(_impStateReport)
	switch req.SFuncName {
	case "Test":
		var A int32
		err = _is.Read_int32(&A, 1, true)
		if err != nil {
			return err
		}
		var B int32
		err = _is.Read_int32(&B, 2, true)
		if err != nil {
			return err
		}
		var C int32
		ret, err := _imp.Test(A, B, &C)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}

		err = _os.Write_int32(C, 3)
		if err != nil {
			return err
		}
	case "TransferToDurid":
		var Num int32
		err = _is.Read_int32(&Num, 1, true)
		if err != nil {
			return err
		}
		ret, err := _imp.TransferToDurid(Num)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      _os.ToBytes(),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	_ = length
	_ = have
	_ = ty
	return nil
}
