//
// This file war generated by FastJce2go 1.0
// Generated from CacheShare.jce
// Tencent.

package DCache

import (
	"fmt"
	"tars/protocol/codec"
)

type MainKeyValue struct {
	MainKey string
	Value   []map[string]string
	Ret     int32
}

func (st *MainKeyValue) resetDefault() {
}

func (st *MainKeyValue) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.MainKey, 1, true)
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
		st.Value = make([]map[string]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err, have = _is.SkipTo(codec.MAP, 0, false)
			if err != nil {
				return err
			}
			if have {
				err = _is.Read_int32(&length, 0, true)
				if err != nil {
					return err
				}
				st.Value[i0] = make(map[string]string)
				for i1, e1 := int32(0), length; i1 < e1; i1++ {
					var k1 string
					var v1 string

					err = _is.Read_string(&k1, 0, false)
					if err != nil {
						return err
					}

					err = _is.Read_string(&v1, 1, false)
					if err != nil {
						return err
					}

					st.Value[i0][k1] = v1
				}
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("type not support SIMPLE_LIST.")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not.")
		if err != nil {
			return err
		}
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

func (st *MainKeyValue) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MainKeyValue, but not exist. tag %d", tag)
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

func (st *MainKeyValue) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.MainKey, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.Value)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.Value {

		err = _os.WriteHead(codec.MAP, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(v)), 0)
		if err != nil {
			return err
		}
		for k2, v2 := range v {

			err = _os.Write_string(k2, 0)
			if err != nil {
				return err
			}

			err = _os.Write_string(v2, 1)
			if err != nil {
				return err
			}
		}
	}

	err = _os.Write_int32(st.Ret, 3)
	if err != nil {
		return err
	}

	return nil
}

func (st *MainKeyValue) WriteBlock(_os *codec.Buffer, tag byte) error {
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