// Code generated by zanzibar
// @generated
// Checksum : TJPAs6uD0PdxOpA1u8hqxg==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

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

func easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(in *jlexer.Lexer, out *SimpleService_Ping_Result) {
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
					out.Success = new(BazResponse)
				}
				easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in, out.Success)
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
func easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(out *jwriter.Writer, in SimpleService_Ping_Result) {
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
		easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out, *in.Success)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Ping_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Ping_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Ping_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Ping_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing(l, v)
}
func easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in *jlexer.Lexer, out *BazResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
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
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out *jwriter.Writer, in BazResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}
func easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(in *jlexer.Lexer, out *SimpleService_Ping_Args) {
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
func easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(out *jwriter.Writer, in SimpleService_Ping_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Ping_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Ping_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Ping_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Ping_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServicePing1(l, v)
}
