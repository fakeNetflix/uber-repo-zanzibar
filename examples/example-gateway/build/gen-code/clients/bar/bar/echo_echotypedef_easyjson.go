// Code generated by zanzibar
// @generated
// Checksum : 1+Y9C5H4AU7cD6ee/bhofA==
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

func easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(in *jlexer.Lexer, out *Echo_EchoTypedef_Result) {
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
					out.Success = new(UUID)
				}
				*out.Success = UUID(in.String())
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
func easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(out *jwriter.Writer, in Echo_EchoTypedef_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoTypedef_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoTypedef_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoTypedef_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoTypedef_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef(l, v)
}
func easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(in *jlexer.Lexer, out *Echo_EchoTypedef_Args) {
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
			out.Arg = UUID(in.String())
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
func easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(out *jwriter.Writer, in Echo_EchoTypedef_Args) {
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
		out.String(string(in.Arg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Echo_EchoTypedef_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Echo_EchoTypedef_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28d9171aEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Echo_EchoTypedef_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Echo_EchoTypedef_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28d9171aDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBarBarEchoEchoTypedef1(l, v)
}
