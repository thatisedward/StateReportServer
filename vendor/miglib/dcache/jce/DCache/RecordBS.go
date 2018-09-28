//
// This file war generated by FastJce2go 1.0
// Generated from CacheShare.jce
// Tencent.

package DCache

import (
	"fmt"
	"tars/protocol/codec"
)

type RecordBS struct {
	MainKey  []int8
	MpRecord map[string][]int8
	Ret      int32
}

func (st *RecordBS) resetDefault() {
}

func (st *RecordBS) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err, have, ty = _is.SkipToNoCheck(1, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.MainKey = make([]int8, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_int8(&st.MainKey[i0], 0, false)
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
		err = _is.Read_slice_int8(&st.MainKey, length, true)
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not.")
		if err != nil {
			return err
		}
	}

	err, have = _is.SkipTo(codec.MAP, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}
	st.MpRecord = make(map[string][]int8)
	for i1, e1 := int32(0), length; i1 < e1; i1++ {
		var k1 string
		var v1 []int8

		err = _is.Read_string(&k1, 0, false)
		if err != nil {
			return err
		}

		err, have, ty = _is.SkipToNoCheck(1, false)
		if err != nil {
			return err
		}
		if have {
			if ty == codec.LIST {
				err = _is.Read_int32(&length, 0, true)
				if err != nil {
					return err
				}
				v1 = make([]int8, length, length)
				for i2, e2 := int32(0), length; i2 < e2; i2++ {

					err = _is.Read_int8(&v1[i2], 0, false)
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
				err = _is.Read_slice_int8(&v1, length, true)
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not.")
				if err != nil {
					return err
				}
			}
		}

		st.MpRecord[k1] = v1
	}

	err = _is.Read_int32(&st.Ret, 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (st *RecordBS) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require RecordBS, but not exist. tag %d", tag)
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

func (st *RecordBS) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.SIMPLE_LIST, 1)
	if err != nil {
		return err
	}
	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MainKey)), 0)
	if err != nil {
		return err
	}
	err = _os.Write_slice_int8(st.MainKey)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MpRecord)), 0)
	if err != nil {
		return err
	}
	for k3, v3 := range st.MpRecord {

		err = _os.Write_string(k3, 0)
		if err != nil {
			return err
		}

		err = _os.WriteHead(codec.SIMPLE_LIST, 1)
		if err != nil {
			return err
		}
		err = _os.WriteHead(codec.BYTE, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(v3)), 0)
		if err != nil {
			return err
		}
		err = _os.Write_slice_int8(v3)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int32(st.Ret, 3)
	if err != nil {
		return err
	}

	return nil
}

func (st *RecordBS) WriteBlock(_os *codec.Buffer, tag byte) error {
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
