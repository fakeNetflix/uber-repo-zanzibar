// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_EchoStructList_Args struct {
	Arg []*base.BazResponse `json:"arg,required"`
}

type _List_BazResponse_ValueList []*base.BazResponse

func (v _List_BazResponse_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_BazResponse_ValueList) Size() int {
	return len(v)
}

func (_List_BazResponse_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_BazResponse_ValueList) Close() {
}

func (v *SecondService_EchoStructList_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		return w, errors.New("field Arg of SecondService_EchoStructList_Args is required")
	}
	w, err = wire.NewValueList(_List_BazResponse_ValueList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_BazResponse_Read(l wire.ValueList) ([]*base.BazResponse, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*base.BazResponse, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _BazResponse_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *SecondService_EchoStructList_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Arg, err = _List_BazResponse_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoStructList_Args is required")
	}
	return nil
}

func (v *SecondService_EchoStructList_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("SecondService_EchoStructList_Args{%v}", strings.Join(fields[:i], ", "))
}

func _List_BazResponse_Equals(lhs, rhs []*base.BazResponse) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, lv := range lhs {
		rv := rhs[i]
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

func (v *SecondService_EchoStructList_Args) Equals(rhs *SecondService_EchoStructList_Args) bool {
	if !_List_BazResponse_Equals(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_EchoStructList_Args) MethodName() string {
	return "echoStructList"
}

func (v *SecondService_EchoStructList_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_EchoStructList_Helper = struct {
	Args           func(arg []*base.BazResponse) *SecondService_EchoStructList_Args
	IsException    func(error) bool
	WrapResponse   func([]*base.BazResponse, error) (*SecondService_EchoStructList_Result, error)
	UnwrapResponse func(*SecondService_EchoStructList_Result) ([]*base.BazResponse, error)
}{}

func init() {
	SecondService_EchoStructList_Helper.Args = func(arg []*base.BazResponse) *SecondService_EchoStructList_Args {
		return &SecondService_EchoStructList_Args{Arg: arg}
	}
	SecondService_EchoStructList_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_EchoStructList_Helper.WrapResponse = func(success []*base.BazResponse, err error) (*SecondService_EchoStructList_Result, error) {
		if err == nil {
			return &SecondService_EchoStructList_Result{Success: success}, nil
		}
		return nil, err
	}
	SecondService_EchoStructList_Helper.UnwrapResponse = func(result *SecondService_EchoStructList_Result) (success []*base.BazResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SecondService_EchoStructList_Result struct {
	Success []*base.BazResponse `json:"success"`
}

func (v *SecondService_EchoStructList_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(_List_BazResponse_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoStructList_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_EchoStructList_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_BazResponse_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SecondService_EchoStructList_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SecondService_EchoStructList_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("SecondService_EchoStructList_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_EchoStructList_Result) Equals(rhs *SecondService_EchoStructList_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _List_BazResponse_Equals(v.Success, rhs.Success))) {
		return false
	}
	return true
}

func (v *SecondService_EchoStructList_Result) MethodName() string {
	return "echoStructList"
}

func (v *SecondService_EchoStructList_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}