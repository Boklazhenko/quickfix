package tradingsessionstatusrequest

import (
	"github.com/Boklazhenko/quickfix"
	"github.com/Boklazhenko/quickfix/gen/enum"
	"github.com/Boklazhenko/quickfix/gen/field"
	"github.com/Boklazhenko/quickfix/gen/fixt11"
	"github.com/Boklazhenko/quickfix/gen/tag"
)

// TradingSessionStatusRequest is the fix50 TradingSessionStatusRequest type, MsgType = g.
type TradingSessionStatusRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TradingSessionStatusRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) TradingSessionStatusRequest {
	return TradingSessionStatusRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m TradingSessionStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a TradingSessionStatusRequest initialized with the required fields for TradingSessionStatusRequest.
func New(tradsesreqid field.TradSesReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m TradingSessionStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("g"))
	m.Set(tradsesreqid)
	m.Set(subscriptionrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "g", r
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m TradingSessionStatusRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m TradingSessionStatusRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetTradSesReqID sets TradSesReqID, Tag 335.
func (m TradingSessionStatusRequest) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m TradingSessionStatusRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradSesMethod sets TradSesMethod, Tag 338.
func (m TradingSessionStatusRequest) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

// SetTradSesMode sets TradSesMode, Tag 339.
func (m TradingSessionStatusRequest) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m TradingSessionStatusRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m TradingSessionStatusRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m TradingSessionStatusRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesReqID gets TradSesReqID, Tag 335.
func (m TradingSessionStatusRequest) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m TradingSessionStatusRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMethod gets TradSesMethod, Tag 338.
func (m TradingSessionStatusRequest) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMode gets TradSesMode, Tag 339.
func (m TradingSessionStatusRequest) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m TradingSessionStatusRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m TradingSessionStatusRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m TradingSessionStatusRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasTradSesReqID returns true if TradSesReqID is present, Tag 335.
func (m TradingSessionStatusRequest) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m TradingSessionStatusRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradSesMethod returns true if TradSesMethod is present, Tag 338.
func (m TradingSessionStatusRequest) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

// HasTradSesMode returns true if TradSesMode is present, Tag 339.
func (m TradingSessionStatusRequest) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m TradingSessionStatusRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
