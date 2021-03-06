//
// This file war generated by FastJce2go 1.0
// Generated from CacheShare.jce
// Tencent.

package DCache

import (
	"fmt"
	"tars/protocol/codec"
)

type Condition struct {
	FieldName string
	Op        Op
	Value     string
}

func (st *Condition) resetDefault() {
}

func (st *Condition) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.FieldName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32((*int32)(&st.Op), 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Value, 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (st *Condition) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Condition, but not exist. tag %d", tag)
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

func (st *Condition) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.FieldName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(st.Op), 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Value, 3)
	if err != nil {
		return err
	}

	return nil
}

func (st *Condition) WriteBlock(_os *codec.Buffer, tag byte) error {
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
