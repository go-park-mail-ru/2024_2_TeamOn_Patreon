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

func easyjson8ba07311DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(in *jlexer.Lexer, out *MediaDeleteRequest) {
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
			if in.IsNull() {
				in.Skip()
				out.MediaIDs = nil
			} else {
				in.Delim('[')
				if out.MediaIDs == nil {
					if !in.IsDelim(']') {
						out.MediaIDs = make([]string, 0, 4)
					} else {
						out.MediaIDs = []string{}
					}
				} else {
					out.MediaIDs = (out.MediaIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.MediaIDs = append(out.MediaIDs, v1)
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
func easyjson8ba07311EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(out *jwriter.Writer, in MediaDeleteRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mediaID\":"
		out.RawString(prefix[1:])
		if in.MediaIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.MediaIDs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MediaDeleteRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8ba07311EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MediaDeleteRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8ba07311EncodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MediaDeleteRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8ba07311DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MediaDeleteRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8ba07311DecodeGithubComGoParkMailRu20242TeamOnPatreonInternalContentControllerModels(l, v)
}
