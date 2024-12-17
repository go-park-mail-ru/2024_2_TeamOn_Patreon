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

func easyjson366e1779DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(in *jlexer.Lexer, out *SubscriptionRequest) {
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
		case "subscriptionRequestID":
			out.SubscriptionRequestID = string(in.String())
		case "authorID":
			out.AuthorID = string(in.String())
		case "monthCount":
			out.MonthCount = int(in.Int())
		case "layer":
			out.Layer = int(in.Int())
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
func easyjson366e1779EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(out *jwriter.Writer, in SubscriptionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"subscriptionRequestID\":"
		out.RawString(prefix[1:])
		out.String(string(in.SubscriptionRequestID))
	}
	{
		const prefix string = ",\"authorID\":"
		out.RawString(prefix)
		out.String(string(in.AuthorID))
	}
	{
		const prefix string = ",\"monthCount\":"
		out.RawString(prefix)
		out.Int(int(in.MonthCount))
	}
	{
		const prefix string = ",\"layer\":"
		out.RawString(prefix)
		out.Int(int(in.Layer))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscriptionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson366e1779EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscriptionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson366e1779EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscriptionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson366e1779DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscriptionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson366e1779DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(l, v)
}