//
// This file war generated by FastJce2go 1.0
// Generated from ProxyShare.jce
// Tencent.

package DCache

import (
	"fmt"
	"tars/protocol/codec"
)

type IKeyValueBS struct {
	KeyItem    int32
	Value      []int8
	Ret        int32
	Ver        int8
	ExpireTime int64
}

func (st *IKeyValueBS) resetDefault() {
}

func (st *IKeyValueBS) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int32(&st.KeyItem, 1, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Value = make([]int8, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_int8(&st.Value[i0], 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SIMPLE_LIST {

		err, _ = _is.SkipTo(codec.BYTE, 0, true)
		if err != nil {
			return err
		}
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		err = _is.Read_slice_int8(&st.Value, length, true)
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not.")
		if err != nil {
			return err
		}
	}

	err = _is.Read_int32(&st.Ret, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int8(&st.Ver, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.ExpireTime, 5, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (st *IKeyValueBS) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require IKeyValueBS, but not exist. tag %d", tag)
		} else {
			return nil
		}
	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

func (st *IKeyValueBS) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.KeyItem, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.SIMPLE_LIST, 2)
	if err != nil {
		return err
	}
	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.Value)), 0)
	if err != nil {
		return err
	}
	err = _os.Write_slice_int8(st.Value)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Ret, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int8(st.Ver, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.ExpireTime, 5)
	if err != nil {
		return err
	}

	return nil
}

func (st *IKeyValueBS) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, tag)
	if err != nil {
		return err
	}
	return nil
}
