package resendrequest

import (
	"github.com/Boklazhenko/quickfix"
	"github.com/Boklazhenko/quickfix/gen/field"
	"github.com/Boklazhenko/quickfix/gen/fixt11"
	"github.com/Boklazhenko/quickfix/gen/tag"
)

// ResendRequest is the fixt11 ResendRequest type, MsgType = 2.
type ResendRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a ResendRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) ResendRequest {
	return ResendRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m ResendRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a ResendRequest initialized with the required fields for ResendRequest.
func New(beginseqno field.BeginSeqNoField, endseqno field.EndSeqNoField) (m ResendRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("2"))
	m.Set(beginseqno)
	m.Set(endseqno)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg ResendRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIXT.1.1", "2", r
}

// SetBeginSeqNo sets BeginSeqNo, Tag 7.
func (m ResendRequest) SetBeginSeqNo(v int) {
	m.Set(field.NewBeginSeqNo(v))
}

// SetEndSeqNo sets EndSeqNo, Tag 16.
func (m ResendRequest) SetEndSeqNo(v int) {
	m.Set(field.NewEndSeqNo(v))
}

// GetBeginSeqNo gets BeginSeqNo, Tag 7.
func (m ResendRequest) GetBeginSeqNo() (v int, err quickfix.MessageRejectError) {
	var f field.BeginSeqNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndSeqNo gets EndSeqNo, Tag 16.
func (m ResendRequest) GetEndSeqNo() (v int, err quickfix.MessageRejectError) {
	var f field.EndSeqNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasBeginSeqNo returns true if BeginSeqNo is present, Tag 7.
func (m ResendRequest) HasBeginSeqNo() bool {
	return m.Has(tag.BeginSeqNo)
}

// HasEndSeqNo returns true if EndSeqNo is present, Tag 16.
func (m ResendRequest) HasEndSeqNo() bool {
	return m.Has(tag.EndSeqNo)
}
