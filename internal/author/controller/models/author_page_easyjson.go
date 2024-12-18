// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	models "github.com/go-park-mail-ru/2024_2_TeamOn_Patreon/internal/author/service/models"
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

func easyjsonD2ce8f2fDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(in *jlexer.Lexer, out *AuthorPage) {
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
		case "authorUsername":
			out.Username = string(in.String())
		case "info":
			out.Info = string(in.String())
		case "followers":
			out.Followers = int(in.Int())
		case "subscriptions":
			if in.IsNull() {
				in.Skip()
				out.Subscriptions = nil
			} else {
				in.Delim('[')
				if out.Subscriptions == nil {
					if !in.IsDelim(']') {
						out.Subscriptions = make([]models.Subscription, 0, 2)
					} else {
						out.Subscriptions = []models.Subscription{}
					}
				} else {
					out.Subscriptions = (out.Subscriptions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.Subscription
					easyjsonD2ce8f2fDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorServiceModels(in, &v1)
					out.Subscriptions = append(out.Subscriptions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "isSubscribe":
			out.UserIsSubscribe = bool(in.Bool())
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
func easyjsonD2ce8f2fEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(out *jwriter.Writer, in AuthorPage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"authorUsername\":"
		out.RawString(prefix[1:])
		out.String(string(in.Username))
	}
	if in.Info != "" {
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"followers\":"
		out.RawString(prefix)
		out.Int(int(in.Followers))
	}
	{
		const prefix string = ",\"subscriptions\":"
		out.RawString(prefix)
		if in.Subscriptions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Subscriptions {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjsonD2ce8f2fEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorServiceModels(out, v3)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"isSubscribe\":"
		out.RawString(prefix)
		out.Bool(bool(in.UserIsSubscribe))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthorPage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2ce8f2fEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthorPage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2ce8f2fEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthorPage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2ce8f2fDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthorPage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2ce8f2fDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorControllerModels(l, v)
}
func easyjsonD2ce8f2fDecodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorServiceModels(in *jlexer.Lexer, out *models.Subscription) {
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
		case "AuthorID":
			out.AuthorID = string(in.String())
		case "AuthorName":
			out.AuthorName = string(in.String())
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
func easyjsonD2ce8f2fEncodeGithubComGoParkMailRu20242TeamOnPatreonInternalAuthorServiceModels(out *jwriter.Writer, in models.Subscription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"AuthorID\":"
		out.RawString(prefix[1:])
		out.String(string(in.AuthorID))
	}
	{
		const prefix string = ",\"AuthorName\":"
		out.RawString(prefix)
		out.String(string(in.AuthorName))
	}
	out.RawByte('}')
}
