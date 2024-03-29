// Code generated by zanzibar
// @generated
// Checksum : X2uDcuKcJRdUoc1Nfk7TVQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package corge

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

func easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(in *jlexer.Lexer, out *Corge_EchoString_Result) {
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
					out.Success = new(string)
				}
				*out.Success = string(in.String())
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
func easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(out *jwriter.Writer, in Corge_EchoString_Result) {
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
func (v Corge_EchoString_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoString_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoString_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoString_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString(l, v)
}
func easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(in *jlexer.Lexer, out *Corge_EchoString_Args) {
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
			out.Arg = string(in.String())
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
func easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(out *jwriter.Writer, in Corge_EchoString_Args) {
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
func (v Corge_EchoString_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Corge_EchoString_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC839be93EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Corge_EchoString_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Corge_EchoString_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC839be93DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsCorgeCorgeCorgeEchoString1(l, v)
}
