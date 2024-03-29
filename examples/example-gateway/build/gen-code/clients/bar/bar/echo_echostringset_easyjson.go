// Code generated by zanzibar
// @generated
// Checksum : MmjmARJxzHasqGV6K880lQ==
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

func easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(in *jlexer.Lexer, out *Echo_EchoStringSet_Result) {
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
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Success = make(map[string]struct{})
				} else {
					out.Success = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 struct{}
					easyjsonA149064aDecode(in, &v1)
					(out.Success)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
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
func easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(out *jwriter.Writer, in Echo_EchoStringSet_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Success) != 0 {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Success {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				easyjsonA149064aEncode(out, v2Value)
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoStringSet_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoStringSet_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoStringSet_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoStringSet_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet(l, v)
}
func easyjsonA149064aDecode(in *jlexer.Lexer, out *struct{}) {
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
func easyjsonA149064aEncode(out *jwriter.Writer, in struct{}) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
func easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(in *jlexer.Lexer, out *Echo_EchoStringSet_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
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
		case "arg":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Arg = make(map[string]struct{})
				} else {
					out.Arg = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 struct{}
					easyjsonA149064aDecode(in, &v3)
					(out.Arg)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(out *jwriter.Writer, in Echo_EchoStringSet_Args) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"arg\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Arg == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Arg {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				easyjsonA149064aEncode(out, v4Value)
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoStringSet_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoStringSet_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA149064aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoStringSet_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoStringSet_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA149064aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoStringSet1(l, v)
}
