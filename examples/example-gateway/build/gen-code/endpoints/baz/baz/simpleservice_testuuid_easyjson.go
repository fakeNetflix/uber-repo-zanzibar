// Code generated by zanzibar
// @generated
// Checksum : uTHHdbYGhsZttgHEX5X/gA==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
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

func easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(in *jlexer.Lexer, out *SimpleService_TestUuid_Result) {
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
func easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(out *jwriter.Writer, in SimpleService_TestUuid_Result) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_TestUuid_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_TestUuid_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_TestUuid_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_TestUuid_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid(l, v)
}
func easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(in *jlexer.Lexer, out *SimpleService_TestUuid_Args) {
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
func easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(out *jwriter.Writer, in SimpleService_TestUuid_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_TestUuid_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_TestUuid_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9189346fEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_TestUuid_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_TestUuid_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9189346fDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTestUuid1(l, v)
}
