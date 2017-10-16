// Code generated by zanzibar
// @generated
// Checksum : U0lB6sBMsW+7sD4mjl8uaw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bar

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(in *jlexer.Lexer, out *Bar_ArgWithManyQueryParams_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(BarResponse)
				}
				easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in, &*out.Success)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(out *jwriter.Writer, in Bar_ArgWithManyQueryParams_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out, *in.Success)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams(l, v)
}
func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(in *jlexer.Lexer, out *BarResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var StringFieldSet bool
	var IntWithRangeSet bool
	var IntWithoutRangeSet bool
	var MapIntWithRangeSet bool
	var MapIntWithoutRangeSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "stringField":
			out.StringField = string(in.String())
			StringFieldSet = true
		case "intWithRange":
			out.IntWithRange = int32(in.Int32())
			IntWithRangeSet = true
		case "intWithoutRange":
			out.IntWithoutRange = int32(in.Int32())
			IntWithoutRangeSet = true
		case "mapIntWithRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithRange = make(map[UUID]int32)
				} else {
					out.MapIntWithRange = nil
				}
				for !in.IsDelim('}') {
					key := UUID(in.String())
					in.WantColon()
					var v1 int32
					v1 = int32(in.Int32())
					(out.MapIntWithRange)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithRangeSet = true
		case "mapIntWithoutRange":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MapIntWithoutRange = make(map[string]int32)
				} else {
					out.MapIntWithoutRange = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 int32
					v2 = int32(in.Int32())
					(out.MapIntWithoutRange)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
			MapIntWithoutRangeSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !StringFieldSet {
		in.AddError(fmt.Errorf("key 'stringField' is required"))
	}
	if !IntWithRangeSet {
		in.AddError(fmt.Errorf("key 'intWithRange' is required"))
	}
	if !IntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'intWithoutRange' is required"))
	}
	if !MapIntWithRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithRange' is required"))
	}
	if !MapIntWithoutRangeSet {
		in.AddError(fmt.Errorf("key 'mapIntWithoutRange' is required"))
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBar(out *jwriter.Writer, in BarResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stringField\":")
	out.String(string(in.StringField))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithRange\":")
	out.Int32(int32(in.IntWithRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"intWithoutRange\":")
	out.Int32(int32(in.IntWithoutRange))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithRange\":")
	if in.MapIntWithRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in.MapIntWithRange {
			if !v3First {
				out.RawByte(',')
			}
			v3First = false
			out.String(string(v3Name))
			out.RawByte(':')
			out.Int32(int32(v3Value))
		}
		out.RawByte('}')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"mapIntWithoutRange\":")
	if in.MapIntWithoutRange == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.MapIntWithoutRange {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			out.Int32(int32(v4Value))
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}
func easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(in *jlexer.Lexer, out *Bar_ArgWithManyQueryParams_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var AStrSet bool
	var ABoolSet bool
	var AInt8Set bool
	var AInt16Set bool
	var AInt32Set bool
	var AInt64Set bool
	var AFloat64Set bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "aStr":
			out.AStr = string(in.String())
			AStrSet = true
		case "anOptStr":
			if in.IsNull() {
				in.Skip()
				out.AnOptStr = nil
			} else {
				if out.AnOptStr == nil {
					out.AnOptStr = new(string)
				}
				*out.AnOptStr = string(in.String())
			}
		case "aBool":
			out.ABool = bool(in.Bool())
			ABoolSet = true
		case "anOptBool":
			if in.IsNull() {
				in.Skip()
				out.AnOptBool = nil
			} else {
				if out.AnOptBool == nil {
					out.AnOptBool = new(bool)
				}
				*out.AnOptBool = bool(in.Bool())
			}
		case "aInt8":
			out.AInt8 = int8(in.Int8())
			AInt8Set = true
		case "anOptInt8":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt8 = nil
			} else {
				if out.AnOptInt8 == nil {
					out.AnOptInt8 = new(int8)
				}
				*out.AnOptInt8 = int8(in.Int8())
			}
		case "aInt16":
			out.AInt16 = int16(in.Int16())
			AInt16Set = true
		case "anOptInt16":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt16 = nil
			} else {
				if out.AnOptInt16 == nil {
					out.AnOptInt16 = new(int16)
				}
				*out.AnOptInt16 = int16(in.Int16())
			}
		case "aInt32":
			out.AInt32 = int32(in.Int32())
			AInt32Set = true
		case "anOptInt32":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt32 = nil
			} else {
				if out.AnOptInt32 == nil {
					out.AnOptInt32 = new(int32)
				}
				*out.AnOptInt32 = int32(in.Int32())
			}
		case "aInt64":
			out.AInt64 = int64(in.Int64())
			AInt64Set = true
		case "anOptInt64":
			if in.IsNull() {
				in.Skip()
				out.AnOptInt64 = nil
			} else {
				if out.AnOptInt64 == nil {
					out.AnOptInt64 = new(int64)
				}
				*out.AnOptInt64 = int64(in.Int64())
			}
		case "aFloat64":
			out.AFloat64 = float64(in.Float64())
			AFloat64Set = true
		case "anOptFloat64":
			if in.IsNull() {
				in.Skip()
				out.AnOptFloat64 = nil
			} else {
				if out.AnOptFloat64 == nil {
					out.AnOptFloat64 = new(float64)
				}
				*out.AnOptFloat64 = float64(in.Float64())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !AStrSet {
		in.AddError(fmt.Errorf("key 'aStr' is required"))
	}
	if !ABoolSet {
		in.AddError(fmt.Errorf("key 'aBool' is required"))
	}
	if !AInt8Set {
		in.AddError(fmt.Errorf("key 'aInt8' is required"))
	}
	if !AInt16Set {
		in.AddError(fmt.Errorf("key 'aInt16' is required"))
	}
	if !AInt32Set {
		in.AddError(fmt.Errorf("key 'aInt32' is required"))
	}
	if !AInt64Set {
		in.AddError(fmt.Errorf("key 'aInt64' is required"))
	}
	if !AFloat64Set {
		in.AddError(fmt.Errorf("key 'aFloat64' is required"))
	}
}
func easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(out *jwriter.Writer, in Bar_ArgWithManyQueryParams_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aStr\":")
	out.String(string(in.AStr))
	if in.AnOptStr != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptStr\":")
		if in.AnOptStr == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.AnOptStr))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aBool\":")
	out.Bool(bool(in.ABool))
	if in.AnOptBool != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptBool\":")
		if in.AnOptBool == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.AnOptBool))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aInt8\":")
	out.Int8(int8(in.AInt8))
	if in.AnOptInt8 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptInt8\":")
		if in.AnOptInt8 == nil {
			out.RawString("null")
		} else {
			out.Int8(int8(*in.AnOptInt8))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aInt16\":")
	out.Int16(int16(in.AInt16))
	if in.AnOptInt16 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptInt16\":")
		if in.AnOptInt16 == nil {
			out.RawString("null")
		} else {
			out.Int16(int16(*in.AnOptInt16))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aInt32\":")
	out.Int32(int32(in.AInt32))
	if in.AnOptInt32 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptInt32\":")
		if in.AnOptInt32 == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.AnOptInt32))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aInt64\":")
	out.Int64(int64(in.AInt64))
	if in.AnOptInt64 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptInt64\":")
		if in.AnOptInt64 == nil {
			out.RawString("null")
		} else {
			out.Int64(int64(*in.AnOptInt64))
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"aFloat64\":")
	out.Float64(float64(in.AFloat64))
	if in.AnOptFloat64 != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"anOptFloat64\":")
		if in.AnOptFloat64 == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.AnOptFloat64))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar_ArgWithManyQueryParams_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a92023EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar_ArgWithManyQueryParams_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a92023DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBarBarBarArgWithManyQueryParams1(l, v)
}
