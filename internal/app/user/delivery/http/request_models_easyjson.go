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

func easyjson7df0efccDecodeTechDbForumInternalAppUserDeliveryHttp(in *jlexer.Lexer, out *UserUpdateRequest) {
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
func easyjson7df0efccEncodeTechDbForumInternalAppUserDeliveryHttp(out *jwriter.Writer, in UserUpdateRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Fullname))
	}
	if in.About != "" {
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
func (v UserUpdateRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodeTechDbForumInternalAppUserDeliveryHttp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUpdateRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodeTechDbForumInternalAppUserDeliveryHttp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUpdateRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodeTechDbForumInternalAppUserDeliveryHttp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUpdateRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodeTechDbForumInternalAppUserDeliveryHttp(l, v)
}
