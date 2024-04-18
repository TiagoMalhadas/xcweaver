// Code generated by "xcweaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package model

import (
	"fmt"
	"github.com/TiagoMalhadas/xcweaver"
	"github.com/TiagoMalhadas/xcweaver/runtime/codegen"
	"socialnetwork/pkg/trace"
)

// xcweaver.InstanceOf checks.

// xcweaver.Router checks.

// Local stub implementations.

// Client stub implementations.

// Note that "xcweaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'xcweaver generate' v0.5.22 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/TiagoMalhadas/xcweaver module that you're using. The xcweaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/TiagoMalhadas/xcweaver

We recommend updating the xcweaver module and the 'xcweaver generate' command by
running the following.

    go get github.com/TiagoMalhadas/xcweaver@latest
    go install github.com/TiagoMalhadas/xcweaver/cmd/weaver@latest

Then, re-run 'xcweaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/TiagoMalhadas/xcweaver/issues.

`)

// Server stub implementations.

// Reflect stub implementations.

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Creator)(nil)

type __is_Creator[T ~struct {
	xcweaver.AutoMarshal
	UserID   int64  "bson:\"user_id\""
	Username string "bson:\"username\""
}] struct{}

var _ __is_Creator[Creator]

func (x *Creator) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Creator.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.UserID)
	enc.String(x.Username)
}

func (x *Creator) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Creator.WeaverUnmarshal: nil receiver"))
	}
	x.UserID = dec.Int64()
	x.Username = dec.String()
}

var _ codegen.AutoMarshal = (*Media)(nil)

type __is_Media[T ~struct {
	xcweaver.AutoMarshal
	MediaID   int64  "bson:\"media_id\""
	MediaType string "bson:\"media_type\""
}] struct{}

var _ __is_Media[Media]

func (x *Media) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Media.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.MediaID)
	enc.String(x.MediaType)
}

func (x *Media) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Media.WeaverUnmarshal: nil receiver"))
	}
	x.MediaID = dec.Int64()
	x.MediaType = dec.String()
}

var _ codegen.AutoMarshal = (*Message)(nil)

type __is_Message[T ~struct {
	xcweaver.AutoMarshal
	ReqID              int64             "json:\"req_id\""
	UserID             int64             "json:\"user_id\""
	PostID             int64             "json:\"post_id\""
	Timestamp          int64             "json:\"timestamp\""
	UserMentionIDs     []int64           "json:\"user_mention_ids\""
	SpanContext        trace.SpanContext "json:\"span_context\""
	NotificationSendTs int64             "json:\"notification_write\""
}] struct{}

var _ __is_Message[Message]

func (x *Message) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Message.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.ReqID)
	enc.Int64(x.UserID)
	enc.Int64(x.PostID)
	enc.Int64(x.Timestamp)
	serviceweaver_enc_slice_int64_a8f7f092(enc, x.UserMentionIDs)
	(x.SpanContext).WeaverMarshal(enc)
	enc.Int64(x.NotificationSendTs)
}

func (x *Message) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Message.WeaverUnmarshal: nil receiver"))
	}
	x.ReqID = dec.Int64()
	x.UserID = dec.Int64()
	x.PostID = dec.Int64()
	x.Timestamp = dec.Int64()
	x.UserMentionIDs = serviceweaver_dec_slice_int64_a8f7f092(dec)
	(&x.SpanContext).WeaverUnmarshal(dec)
	x.NotificationSendTs = dec.Int64()
}

func serviceweaver_enc_slice_int64_a8f7f092(enc *codegen.Encoder, arg []int64) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Int64(arg[i])
	}
}

func serviceweaver_dec_slice_int64_a8f7f092(dec *codegen.Decoder) []int64 {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Int64()
	}
	return res
}

var _ codegen.AutoMarshal = (*Post)(nil)

type __is_Post[T ~struct {
	xcweaver.AutoMarshal
	PostID       int64         "bson:\"post_id\""
	ReqID        int64         "bson:\"req_id\""
	Creator      Creator       "bson:\"creator\""
	Text         string        "bson:\"text\""
	UserMentions []UserMention "bson:\"user_mentions\""
	Media        []Media       "bson:\"media\""
	URLs         []URL         "bson:\"urls\""
	Timestamp    int64         "bson:\"timestamp\""
	PostType     PostType      "bson:\"posttype\""
}] struct{}

var _ __is_Post[Post]

func (x *Post) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Post.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.PostID)
	enc.Int64(x.ReqID)
	(x.Creator).WeaverMarshal(enc)
	enc.String(x.Text)
	serviceweaver_enc_slice_UserMention_6ddd52fd(enc, x.UserMentions)
	serviceweaver_enc_slice_Media_d3cc456b(enc, x.Media)
	serviceweaver_enc_slice_URL_0b7470c4(enc, x.URLs)
	enc.Int64(x.Timestamp)
	enc.Int((int)(x.PostType))
}

func (x *Post) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Post.WeaverUnmarshal: nil receiver"))
	}
	x.PostID = dec.Int64()
	x.ReqID = dec.Int64()
	(&x.Creator).WeaverUnmarshal(dec)
	x.Text = dec.String()
	x.UserMentions = serviceweaver_dec_slice_UserMention_6ddd52fd(dec)
	x.Media = serviceweaver_dec_slice_Media_d3cc456b(dec)
	x.URLs = serviceweaver_dec_slice_URL_0b7470c4(dec)
	x.Timestamp = dec.Int64()
	*(*int)(&x.PostType) = dec.Int()
}

func serviceweaver_enc_slice_UserMention_6ddd52fd(enc *codegen.Encoder, arg []UserMention) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_UserMention_6ddd52fd(dec *codegen.Decoder) []UserMention {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]UserMention, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_Media_d3cc456b(enc *codegen.Encoder, arg []Media) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Media_d3cc456b(dec *codegen.Decoder) []Media {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Media, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

func serviceweaver_enc_slice_URL_0b7470c4(enc *codegen.Encoder, arg []URL) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_URL_0b7470c4(dec *codegen.Decoder) []URL {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]URL, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}

var _ codegen.AutoMarshal = (*URL)(nil)

type __is_URL[T ~struct {
	xcweaver.AutoMarshal
	ExpandedUrl  string "bson:\"expanded_url\""
	ShortenedUrl string "bson:\"shortened_url\""
}] struct{}

var _ __is_URL[URL]

func (x *URL) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("URL.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ExpandedUrl)
	enc.String(x.ShortenedUrl)
}

func (x *URL) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("URL.WeaverUnmarshal: nil receiver"))
	}
	x.ExpandedUrl = dec.String()
	x.ShortenedUrl = dec.String()
}

var _ codegen.AutoMarshal = (*User)(nil)

type __is_User[T ~struct {
	xcweaver.AutoMarshal
	UserID    int64  "bson:\"user_id\""
	FirstName string "bson:\"first_name\""
	LastName  string "bson:\"last_name\""
	Username  string "bson:\"username\""
	PwdHashed string "bson:\"pwd_hashed\""
	Salt      string "bson:\"salt\""
}] struct{}

var _ __is_User[User]

func (x *User) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.UserID)
	enc.String(x.FirstName)
	enc.String(x.LastName)
	enc.String(x.Username)
	enc.String(x.PwdHashed)
	enc.String(x.Salt)
}

func (x *User) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("User.WeaverUnmarshal: nil receiver"))
	}
	x.UserID = dec.Int64()
	x.FirstName = dec.String()
	x.LastName = dec.String()
	x.Username = dec.String()
	x.PwdHashed = dec.String()
	x.Salt = dec.String()
}

var _ codegen.AutoMarshal = (*UserMention)(nil)

type __is_UserMention[T ~struct {
	xcweaver.AutoMarshal
	UserID   int64  "bson:\"user_id\""
	Username string "bson:\"username\""
}] struct{}

var _ __is_UserMention[UserMention]

func (x *UserMention) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("UserMention.WeaverMarshal: nil receiver"))
	}
	enc.Int64(x.UserID)
	enc.String(x.Username)
}

func (x *UserMention) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("UserMention.WeaverUnmarshal: nil receiver"))
	}
	x.UserID = dec.Int64()
	x.Username = dec.String()
}