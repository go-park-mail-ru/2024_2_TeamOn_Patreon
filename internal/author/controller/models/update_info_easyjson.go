// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonE0a3d8c0DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(in *jlexer.Lexer, out *UpdateInfo) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "info":
			out.Info = string(in.String())
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
func easyjsonE0a3d8c0EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(out *jwriter.Writer, in UpdateInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix[1:])
		out.String(string(in.Info))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE0a3d8c0EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE0a3d8c0EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE0a3d8c0DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE0a3d8c0DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(l, v)
}
