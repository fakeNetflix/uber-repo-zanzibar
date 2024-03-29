// Code generated by zanzibar
// @generated
// Checksum : pCBivHdeILFrGKDQdJyqHw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package withexceptions

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

func easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(in *jlexer.Lexer, out *WithExceptions_Func1_Result) {
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
					out.Success = new(Response)
				}
				(*out.Success).UnmarshalEasyJSON(in)
			}
		case "e1":
			if in.IsNull() {
				in.Skip()
				out.E1 = nil
			} else {
				if out.E1 == nil {
					out.E1 = new(ExceptionType1)
				}
				(*out.E1).UnmarshalEasyJSON(in)
			}
		case "e2":
			if in.IsNull() {
				in.Skip()
				out.E2 = nil
			} else {
				if out.E2 == nil {
					out.E2 = new(ExceptionType2)
				}
				(*out.E2).UnmarshalEasyJSON(in)
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
func easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(out *jwriter.Writer, in WithExceptions_Func1_Result) {
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
		(*in.Success).MarshalEasyJSON(out)
	}
	if in.E1 != nil {
		const prefix string = ",\"e1\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.E1).MarshalEasyJSON(out)
	}
	if in.E2 != nil {
		const prefix string = ",\"e2\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.E2).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithExceptions_Func1_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WithExceptions_Func1_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithExceptions_Func1_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WithExceptions_Func1_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc1(l, v)
}
func easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(in *jlexer.Lexer, out *WithExceptions_Func1_Args) {
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
func easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(out *jwriter.Writer, in WithExceptions_Func1_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WithExceptions_Func1_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WithExceptions_Func1_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFd689198EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WithExceptions_Func1_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WithExceptions_Func1_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFd689198DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsWithexceptionsWithexceptionsWithExceptionsFunc11(l, v)
}
