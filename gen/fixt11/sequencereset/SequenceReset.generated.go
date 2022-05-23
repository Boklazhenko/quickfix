package sequencereset

import (
	"github.com/Boklazhenko/quickfix"
	"github.com/Boklazhenko/quickfix/gen/field"
	"github.com/Boklazhenko/quickfix/gen/fixt11"
	"github.com/Boklazhenko/quickfix/gen/tag"
)

// SequenceReset is the fixt11 SequenceReset type, MsgType = 4.
type SequenceReset struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a SequenceReset from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) SequenceReset {
	return SequenceReset{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m SequenceReset) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a SequenceReset initialized with the required fields for SequenceReset.
func New(newseqno field.NewSeqNoField) (m SequenceReset) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("4"))
	m.Set(newseqno)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg SequenceReset, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIXT.1.1", "4", r
}

// SetNewSeqNo sets NewSeqNo, Tag 36.
func (m SequenceReset) SetNewSeqNo(v int) {
	m.Set(field.NewNewSeqNo(v))
}

// SetGapFillFlag sets GapFillFlag, Tag 123.
func (m SequenceReset) SetGapFillFlag(v bool) {
	m.Set(field.NewGapFillFlag(v))
}

// GetNewSeqNo gets NewSeqNo, Tag 36.
func (m SequenceReset) GetNewSeqNo() (v int, err quickfix.MessageRejectError) {
	var f field.NewSeqNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetGapFillFlag gets GapFillFlag, Tag 123.
func (m SequenceReset) GetGapFillFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.GapFillFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNewSeqNo returns true if NewSeqNo is present, Tag 36.
func (m SequenceReset) HasNewSeqNo() bool {
	return m.Has(tag.NewSeqNo)
}

// HasGapFillFlag returns true if GapFillFlag is present, Tag 123.
func (m SequenceReset) HasGapFillFlag() bool {
	return m.Has(tag.GapFillFlag)
}
