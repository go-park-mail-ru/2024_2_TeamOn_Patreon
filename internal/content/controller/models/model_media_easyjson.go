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

func easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(in *jlexer.Lexer, out *MediaResponse) {
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
		case "postId":
			out.PostID = string(in.String())
		case "mediaContent":
			if in.IsNull() {
				in.Skip()
				out.MediaContent = nil
			} else {
				in.Delim('[')
				if out.MediaContent == nil {
					if !in.IsDelim(']') {
						out.MediaContent = make([]*Media, 0, 8)
					} else {
						out.MediaContent = []*Media{}
					}
				} else {
					out.MediaContent = (out.MediaContent)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Media
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Media)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.MediaContent = append(out.MediaContent, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(out *jwriter.Writer, in MediaResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"postId\":"
		out.RawString(prefix[1:])
		out.String(string(in.PostID))
	}
	{
		const prefix string = ",\"mediaContent\":"
		out.RawString(prefix)
		if in.MediaContent == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MediaContent {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MediaResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MediaResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MediaResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MediaResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(l, v)
}
func easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(in *jlexer.Lexer, out *Media) {
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
		case "mediaID":
			out.MediaID = string(in.String())
		case "mediaType":
			out.MediaType = string(in.String())
		case "mediaURL":
			out.MediaURL = string(in.String())
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
func easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(out *jwriter.Writer, in Media) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mediaID\":"
		out.RawString(prefix[1:])
		out.String(string(in.MediaID))
	}
	{
		const prefix string = ",\"mediaType\":"
		out.RawString(prefix)
		out.String(string(in.MediaType))
	}
	{
		const prefix string = ",\"mediaURL\":"
		out.RawString(prefix)
		out.String(string(in.MediaURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Media) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Media) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9574b300EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Media) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Media) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9574b300DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels1(l, v)
}
