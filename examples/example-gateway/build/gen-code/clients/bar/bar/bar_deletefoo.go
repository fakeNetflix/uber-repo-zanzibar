// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package bar

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Bar_DeleteFoo_Args represents the arguments for the Bar.deleteFoo function.
//
// The arguments for deleteFoo are sent and received over the wire as this struct.
type Bar_DeleteFoo_Args struct {
	Request *BarRequest `json:"request,required"`
}

// ToWire translates a Bar_DeleteFoo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_DeleteFoo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request == nil {
		return w, errors.New("field Request of Bar_DeleteFoo_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_DeleteFoo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_DeleteFoo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_DeleteFoo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_DeleteFoo_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _BarRequest_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_DeleteFoo_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_DeleteFoo_Args
// struct.
func (v *Bar_DeleteFoo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++

	return fmt.Sprintf("Bar_DeleteFoo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_DeleteFoo_Args match the
// provided Bar_DeleteFoo_Args.
//
// This function performs a deep comparison.
func (v *Bar_DeleteFoo_Args) Equals(rhs *Bar_DeleteFoo_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Request.Equals(rhs.Request) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_DeleteFoo_Args.
func (v *Bar_DeleteFoo_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("request", v.Request))
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_DeleteFoo_Args) GetRequest() (o *BarRequest) {
	if v != nil {
		o = v.Request
	}
	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Bar_DeleteFoo_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "deleteFoo" for this struct.
func (v *Bar_DeleteFoo_Args) MethodName() string {
	return "deleteFoo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_DeleteFoo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_DeleteFoo_Helper provides functions that aid in handling the
// parameters and return values of the Bar.deleteFoo
// function.
var Bar_DeleteFoo_Helper = struct {
	// Args accepts the parameters of deleteFoo in-order and returns
	// the arguments struct for the function.
	Args func(
		request *BarRequest,
	) *Bar_DeleteFoo_Args

	// IsException returns true if the given error can be thrown
	// by deleteFoo.
	//
	// An error can be thrown by deleteFoo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for deleteFoo
	// given the error returned by it. The provided error may
	// be nil if deleteFoo did not fail.
	//
	// This allows mapping errors returned by deleteFoo into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// deleteFoo
	//
	//   err := deleteFoo(args)
	//   result, err := Bar_DeleteFoo_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from deleteFoo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Bar_DeleteFoo_Result, error)

	// UnwrapResponse takes the result struct for deleteFoo
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if deleteFoo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Bar_DeleteFoo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_DeleteFoo_Result) error
}{}

func init() {
	Bar_DeleteFoo_Helper.Args = func(
		request *BarRequest,
	) *Bar_DeleteFoo_Args {
		return &Bar_DeleteFoo_Args{
			Request: request,
		}
	}

	Bar_DeleteFoo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Bar_DeleteFoo_Helper.WrapResponse = func(err error) (*Bar_DeleteFoo_Result, error) {
		if err == nil {
			return &Bar_DeleteFoo_Result{}, nil
		}

		return nil, err
	}
	Bar_DeleteFoo_Helper.UnwrapResponse = func(result *Bar_DeleteFoo_Result) (err error) {
		return
	}

}

// Bar_DeleteFoo_Result represents the result of a Bar.deleteFoo function call.
//
// The result of a deleteFoo execution is sent and received over the wire as this struct.
type Bar_DeleteFoo_Result struct {
}

// ToWire translates a Bar_DeleteFoo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar_DeleteFoo_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_DeleteFoo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_DeleteFoo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_DeleteFoo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_DeleteFoo_Result) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Bar_DeleteFoo_Result
// struct.
func (v *Bar_DeleteFoo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Bar_DeleteFoo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_DeleteFoo_Result match the
// provided Bar_DeleteFoo_Result.
//
// This function performs a deep comparison.
func (v *Bar_DeleteFoo_Result) Equals(rhs *Bar_DeleteFoo_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_DeleteFoo_Result.
func (v *Bar_DeleteFoo_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "deleteFoo" for this struct.
func (v *Bar_DeleteFoo_Result) MethodName() string {
	return "deleteFoo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_DeleteFoo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
