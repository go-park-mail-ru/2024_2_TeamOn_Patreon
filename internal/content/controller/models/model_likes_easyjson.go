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

func easyjsonDc9f0cf4DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(in *jlexer.Lexer, out *Likes) {
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
		case "count":
			out.Count = int(in.Int())
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
func easyjsonDc9f0cf4EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(out *jwriter.Writer, in Likes) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Count))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Likes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDc9f0cf4EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Likes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDc9f0cf4EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Likes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDc9f0cf4DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Likes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDc9f0cf4DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(l, v)
}