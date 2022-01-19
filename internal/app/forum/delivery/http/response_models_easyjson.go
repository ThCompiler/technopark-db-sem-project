// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http_delivery

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

func easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp(in *jlexer.Lexer, out *UsersResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UsersResponse, 0, 1)
			} else {
				*out = UsersResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 UserResponse
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
func easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp(out *jwriter.Writer, in UsersResponse) {
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
func (v UsersResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp1(in *jlexer.Lexer, out *UserResponse) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp1(out *jwriter.Writer, in UserResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp1(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp2(in *jlexer.Lexer, out *ThreadsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ThreadsResponse, 0, 0)
			} else {
				*out = ThreadsResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 ThreadResponse
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp2(out *jwriter.Writer, in ThreadsResponse) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp2(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp3(in *jlexer.Lexer, out *ThreadResponse) {
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
func easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp3(out *jwriter.Writer, in ThreadResponse) {
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
	{
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
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp3(l, v)
}
func easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp4(in *jlexer.Lexer, out *ForumResponse) {
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
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "posts":
			out.Posts = int64(in.Int64())
		case "threads":
			out.Threads = int64(in.Int64())
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
func easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp4(out *jwriter.Writer, in ForumResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"posts\":"
		out.RawString(prefix)
		out.Int64(int64(in.Posts))
	}
	{
		const prefix string = ",\"threads\":"
		out.RawString(prefix)
		out.Int64(int64(in.Threads))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ForumResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ForumResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeTechDbForumInternalAppForumDeliveryHttp4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ForumResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ForumResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeTechDbForumInternalAppForumDeliveryHttp4(l, v)
}
