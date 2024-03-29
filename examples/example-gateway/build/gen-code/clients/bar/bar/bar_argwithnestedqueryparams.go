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

// Bar_ArgWithNestedQueryParams_Args represents the arguments for the Bar.argWithNestedQueryParams function.
//
// The arguments for argWithNestedQueryParams are sent and received over the wire as this struct.
type Bar_ArgWithNestedQueryParams_Args struct {
	Request *QueryParamsStruct     `json:"request,required"`
	Opt     *QueryParamsOptsStruct `json:"opt,omitempty"`
}

// ToWire translates a Bar_ArgWithNestedQueryParams_Args struct into a Thrift-level intermediate
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
func (v *Bar_ArgWithNestedQueryParams_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request == nil {
		return w, errors.New("field Request of Bar_ArgWithNestedQueryParams_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Opt != nil {
		w, err = v.Opt.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _QueryParamsStruct_Read(w wire.Value) (*QueryParamsStruct, error) {
	var v QueryParamsStruct
	err := v.FromWire(w)
	return &v, err
}

func _QueryParamsOptsStruct_Read(w wire.Value) (*QueryParamsOptsStruct, error) {
	var v QueryParamsOptsStruct
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_ArgWithNestedQueryParams_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithNestedQueryParams_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithNestedQueryParams_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithNestedQueryParams_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _QueryParamsStruct_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Opt, err = _QueryParamsOptsStruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_ArgWithNestedQueryParams_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithNestedQueryParams_Args
// struct.
func (v *Bar_ArgWithNestedQueryParams_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++
	if v.Opt != nil {
		fields[i] = fmt.Sprintf("Opt: %v", v.Opt)
		i++
	}

	return fmt.Sprintf("Bar_ArgWithNestedQueryParams_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgWithNestedQueryParams_Args match the
// provided Bar_ArgWithNestedQueryParams_Args.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithNestedQueryParams_Args) Equals(rhs *Bar_ArgWithNestedQueryParams_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Request.Equals(rhs.Request) {
		return false
	}
	if !((v.Opt == nil && rhs.Opt == nil) || (v.Opt != nil && rhs.Opt != nil && v.Opt.Equals(rhs.Opt))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithNestedQueryParams_Args.
func (v *Bar_ArgWithNestedQueryParams_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("request", v.Request))
	if v.Opt != nil {
		err = multierr.Append(err, enc.AddObject("opt", v.Opt))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithNestedQueryParams_Args) GetRequest() (o *QueryParamsStruct) {
	if v != nil {
		o = v.Request
	}
	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Bar_ArgWithNestedQueryParams_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// GetOpt returns the value of Opt if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithNestedQueryParams_Args) GetOpt() (o *QueryParamsOptsStruct) {
	if v != nil && v.Opt != nil {
		return v.Opt
	}

	return
}

// IsSetOpt returns true if Opt is not nil.
func (v *Bar_ArgWithNestedQueryParams_Args) IsSetOpt() bool {
	return v != nil && v.Opt != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "argWithNestedQueryParams" for this struct.
func (v *Bar_ArgWithNestedQueryParams_Args) MethodName() string {
	return "argWithNestedQueryParams"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_ArgWithNestedQueryParams_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_ArgWithNestedQueryParams_Helper provides functions that aid in handling the
// parameters and return values of the Bar.argWithNestedQueryParams
// function.
var Bar_ArgWithNestedQueryParams_Helper = struct {
	// Args accepts the parameters of argWithNestedQueryParams in-order and returns
	// the arguments struct for the function.
	Args func(
		request *QueryParamsStruct,
		opt *QueryParamsOptsStruct,
	) *Bar_ArgWithNestedQueryParams_Args

	// IsException returns true if the given error can be thrown
	// by argWithNestedQueryParams.
	//
	// An error can be thrown by argWithNestedQueryParams only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for argWithNestedQueryParams
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// argWithNestedQueryParams into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by argWithNestedQueryParams
	//
	//   value, err := argWithNestedQueryParams(args)
	//   result, err := Bar_ArgWithNestedQueryParams_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from argWithNestedQueryParams: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponse, error) (*Bar_ArgWithNestedQueryParams_Result, error)

	// UnwrapResponse takes the result struct for argWithNestedQueryParams
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if argWithNestedQueryParams threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_ArgWithNestedQueryParams_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_ArgWithNestedQueryParams_Result) (*BarResponse, error)
}{}

func init() {
	Bar_ArgWithNestedQueryParams_Helper.Args = func(
		request *QueryParamsStruct,
		opt *QueryParamsOptsStruct,
	) *Bar_ArgWithNestedQueryParams_Args {
		return &Bar_ArgWithNestedQueryParams_Args{
			Request: request,
			Opt:     opt,
		}
	}

	Bar_ArgWithNestedQueryParams_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Bar_ArgWithNestedQueryParams_Helper.WrapResponse = func(success *BarResponse, err error) (*Bar_ArgWithNestedQueryParams_Result, error) {
		if err == nil {
			return &Bar_ArgWithNestedQueryParams_Result{Success: success}, nil
		}

		return nil, err
	}
	Bar_ArgWithNestedQueryParams_Helper.UnwrapResponse = func(result *Bar_ArgWithNestedQueryParams_Result) (success *BarResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_ArgWithNestedQueryParams_Result represents the result of a Bar.argWithNestedQueryParams function call.
//
// The result of a argWithNestedQueryParams execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_ArgWithNestedQueryParams_Result struct {
	// Value returned by argWithNestedQueryParams after a successful execution.
	Success *BarResponse `json:"success,omitempty"`
}

// ToWire translates a Bar_ArgWithNestedQueryParams_Result struct into a Thrift-level intermediate
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
func (v *Bar_ArgWithNestedQueryParams_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_ArgWithNestedQueryParams_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_ArgWithNestedQueryParams_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_ArgWithNestedQueryParams_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_ArgWithNestedQueryParams_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_ArgWithNestedQueryParams_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponse_Read(field.Value)
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
		return fmt.Errorf("Bar_ArgWithNestedQueryParams_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_ArgWithNestedQueryParams_Result
// struct.
func (v *Bar_ArgWithNestedQueryParams_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Bar_ArgWithNestedQueryParams_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_ArgWithNestedQueryParams_Result match the
// provided Bar_ArgWithNestedQueryParams_Result.
//
// This function performs a deep comparison.
func (v *Bar_ArgWithNestedQueryParams_Result) Equals(rhs *Bar_ArgWithNestedQueryParams_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_ArgWithNestedQueryParams_Result.
func (v *Bar_ArgWithNestedQueryParams_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_ArgWithNestedQueryParams_Result) GetSuccess() (o *BarResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_ArgWithNestedQueryParams_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "argWithNestedQueryParams" for this struct.
func (v *Bar_ArgWithNestedQueryParams_Result) MethodName() string {
	return "argWithNestedQueryParams"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_ArgWithNestedQueryParams_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
