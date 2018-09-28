//
// This file war generated by FastJce2go 1.0
// Generated from CacheShare.jce
// Tencent.

package DCache

import (
	"fmt"
	"tars/protocol/codec"
)

type MainKeyCondition struct {
	MainKey    string
	Field      string
	Cond       [][]Condition
	Limit      Condition
	BGetMKCout bool
}

func (st *MainKeyCondition) resetDefault() {
	st.BGetMKCout = false
}

func (st *MainKeyCondition) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.MainKey, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Field, 2, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Cond = make([][]Condition, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err, have, ty = _is.SkipToNoCheck(0, false)
			if err != nil {
				return err
			}
			if have {
				if ty == codec.LIST {
					err = _is.Read_int32(&length, 0, true)
					if err != nil {
						return err
					}
					st.Cond[i0] = make([]Condition, length, length)
					for i1, e1 := int32(0), length; i1 < e1; i1++ {

						err = st.Cond[i0][i1].ReadBlock(_is, 0, false)
						if err != nil {
							return err
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

	err = st.Limit.ReadBlock(_is, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BGetMKCout, 5, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (st *MainKeyCondition) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MainKeyCondition, but not exist. tag %d", tag)
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

func (st *MainKeyCondition) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.MainKey, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Field, 2)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.Cond)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.Cond {

		err = _os.WriteHead(codec.LIST, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(v)), 0)
		if err != nil {
			return err
		}
		for _, v := range v {

			err = v.WriteBlock(_os, 0)
			if err != nil {
				return err
			}
		}
	}

	err = st.Limit.WriteBlock(_os, 4)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BGetMKCout, 5)
	if err != nil {
		return err
	}

	return nil
}

func (st *MainKeyCondition) WriteBlock(_os *codec.Buffer, tag byte) error {
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
