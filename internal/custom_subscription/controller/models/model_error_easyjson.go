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

func easyjson2b7acbfeDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(in *jlexer.Lexer, out *ModelError) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjson2b7acbfeEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(out *jwriter.Writer, in ModelError) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ModelError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2b7acbfeEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ModelError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2b7acbfeEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ModelError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2b7acbfeDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ModelError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2b7acbfeDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalCustomSubscriptionControllerModels(l, v)
}
