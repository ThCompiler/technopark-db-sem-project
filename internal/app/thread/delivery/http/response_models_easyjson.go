// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http_delivery

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	thread "tech-db-forum/internal/app/thread"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp(in *jlexer.Lexer, out *ThreadResponse) {
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
		case "id":
			out.Id = int64(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "forum":
			out.Forum = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "votes":
			out.Votes = int64(in.Int64())
		case "slug":
			out.Slug = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp(out *jwriter.Writer, in ThreadResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		out.Int64(int64(in.Votes))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp1(in *jlexer.Lexer, out *PostsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostsResponse, 0, 0)
			} else {
				*out = PostsResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 thread.Post
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp1(out *jwriter.Writer, in PostsResponse) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp1(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp2(in *jlexer.Lexer, out *PostResponse) {
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
		case "id":
			out.Id = int64(in.Int64())
		case "parent":
			out.Parent = int64(in.Int64())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "isEdited":
			out.Is_Edited = bool(in.Bool())
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = int64(in.Int64())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp2(out *jwriter.Writer, in PostResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Id))
	}
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int64(int64(in.Parent))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.Is_Edited))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int64(int64(in.Thread))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppThreadDeliveryHttp2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppThreadDeliveryHttp2(l, v)
}
