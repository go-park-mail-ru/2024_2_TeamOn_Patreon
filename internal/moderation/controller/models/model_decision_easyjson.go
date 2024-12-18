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

func easyjsonCc5c9a00DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(in *jlexer.Lexer, out *Decision) {
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
		case "postID":
			out.PostID = string(in.String())
		case "status":
			out.Status = string(in.String())
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
func easyjsonCc5c9a00EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(out *jwriter.Writer, in Decision) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"postID\":"
		out.RawString(prefix[1:])
		out.String(string(in.PostID))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Decision) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCc5c9a00EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Decision) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCc5c9a00EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Decision) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCc5c9a00DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Decision) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCc5c9a00DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalModerationControllerModels(l, v)
}
