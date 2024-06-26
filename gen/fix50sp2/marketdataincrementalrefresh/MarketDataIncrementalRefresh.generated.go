package marketdataincrementalrefresh

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/Boklazhenko/quickfix"
	"github.com/Boklazhenko/quickfix/gen/enum"
	"github.com/Boklazhenko/quickfix/gen/field"
	"github.com/Boklazhenko/quickfix/gen/fixt11"
	"github.com/Boklazhenko/quickfix/gen/tag"
)

// MarketDataIncrementalRefresh is the fix50sp2 MarketDataIncrementalRefresh type, MsgType = X.
type MarketDataIncrementalRefresh struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataIncrementalRefresh from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataIncrementalRefresh {
	return MarketDataIncrementalRefresh{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataIncrementalRefresh) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataIncrementalRefresh initialized with the required fields for MarketDataIncrementalRefresh.
func New() (m MarketDataIncrementalRefresh) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("X"))

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "X", r
}

// SetTradeDate sets TradeDate, Tag 75.
func (m MarketDataIncrementalRefresh) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetNoRoutingIDs sets NoRoutingIDs, Tag 215.
func (m MarketDataIncrementalRefresh) SetNoRoutingIDs(f NoRoutingIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetMDReqID sets MDReqID, Tag 262.
func (m MarketDataIncrementalRefresh) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

// SetNoMDEntries sets NoMDEntries, Tag 268.
func (m MarketDataIncrementalRefresh) SetNoMDEntries(f NoMDEntriesRepeatingGroup) {
	m.SetGroup(f)
}

// SetApplQueueDepth sets ApplQueueDepth, Tag 813.
func (m MarketDataIncrementalRefresh) SetApplQueueDepth(v int) {
	m.Set(field.NewApplQueueDepth(v))
}

// SetApplQueueResolution sets ApplQueueResolution, Tag 814.
func (m MarketDataIncrementalRefresh) SetApplQueueResolution(v enum.ApplQueueResolution) {
	m.Set(field.NewApplQueueResolution(v))
}

// SetMDBookType sets MDBookType, Tag 1021.
func (m MarketDataIncrementalRefresh) SetMDBookType(v enum.MDBookType) {
	m.Set(field.NewMDBookType(v))
}

// SetMDFeedType sets MDFeedType, Tag 1022.
func (m MarketDataIncrementalRefresh) SetMDFeedType(v string) {
	m.Set(field.NewMDFeedType(v))
}

// SetApplID sets ApplID, Tag 1180.
func (m MarketDataIncrementalRefresh) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

// SetApplSeqNum sets ApplSeqNum, Tag 1181.
func (m MarketDataIncrementalRefresh) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

// SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350.
func (m MarketDataIncrementalRefresh) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

// SetApplResendFlag sets ApplResendFlag, Tag 1352.
func (m MarketDataIncrementalRefresh) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

// GetTradeDate gets TradeDate, Tag 75.
func (m MarketDataIncrementalRefresh) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRoutingIDs gets NoRoutingIDs, Tag 215.
func (m MarketDataIncrementalRefresh) GetNoRoutingIDs() (NoRoutingIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRoutingIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataIncrementalRefresh) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMDEntries gets NoMDEntries, Tag 268.
func (m MarketDataIncrementalRefresh) GetNoMDEntries() (NoMDEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetApplQueueDepth gets ApplQueueDepth, Tag 813.
func (m MarketDataIncrementalRefresh) GetApplQueueDepth() (v int, err quickfix.MessageRejectError) {
	var f field.ApplQueueDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplQueueResolution gets ApplQueueResolution, Tag 814.
func (m MarketDataIncrementalRefresh) GetApplQueueResolution() (v enum.ApplQueueResolution, err quickfix.MessageRejectError) {
	var f field.ApplQueueResolutionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDBookType gets MDBookType, Tag 1021.
func (m MarketDataIncrementalRefresh) GetMDBookType() (v enum.MDBookType, err quickfix.MessageRejectError) {
	var f field.MDBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDFeedType gets MDFeedType, Tag 1022.
func (m MarketDataIncrementalRefresh) GetMDFeedType() (v string, err quickfix.MessageRejectError) {
	var f field.MDFeedTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplID gets ApplID, Tag 1180.
func (m MarketDataIncrementalRefresh) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplSeqNum gets ApplSeqNum, Tag 1181.
func (m MarketDataIncrementalRefresh) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350.
func (m MarketDataIncrementalRefresh) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetApplResendFlag gets ApplResendFlag, Tag 1352.
func (m MarketDataIncrementalRefresh) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTradeDate returns true if TradeDate is present, Tag 75.
func (m MarketDataIncrementalRefresh) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasNoRoutingIDs returns true if NoRoutingIDs is present, Tag 215.
func (m MarketDataIncrementalRefresh) HasNoRoutingIDs() bool {
	return m.Has(tag.NoRoutingIDs)
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataIncrementalRefresh) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasNoMDEntries returns true if NoMDEntries is present, Tag 268.
func (m MarketDataIncrementalRefresh) HasNoMDEntries() bool {
	return m.Has(tag.NoMDEntries)
}

// HasApplQueueDepth returns true if ApplQueueDepth is present, Tag 813.
func (m MarketDataIncrementalRefresh) HasApplQueueDepth() bool {
	return m.Has(tag.ApplQueueDepth)
}

// HasApplQueueResolution returns true if ApplQueueResolution is present, Tag 814.
func (m MarketDataIncrementalRefresh) HasApplQueueResolution() bool {
	return m.Has(tag.ApplQueueResolution)
}

// HasMDBookType returns true if MDBookType is present, Tag 1021.
func (m MarketDataIncrementalRefresh) HasMDBookType() bool {
	return m.Has(tag.MDBookType)
}

// HasMDFeedType returns true if MDFeedType is present, Tag 1022.
func (m MarketDataIncrementalRefresh) HasMDFeedType() bool {
	return m.Has(tag.MDFeedType)
}

// HasApplID returns true if ApplID is present, Tag 1180.
func (m MarketDataIncrementalRefresh) HasApplID() bool {
	return m.Has(tag.ApplID)
}

// HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181.
func (m MarketDataIncrementalRefresh) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

// HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350.
func (m MarketDataIncrementalRefresh) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

// HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352.
func (m MarketDataIncrementalRefresh) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

// NoRoutingIDs is a repeating group element, Tag 215.
type NoRoutingIDs struct {
	*quickfix.Group
}

// SetRoutingType sets RoutingType, Tag 216.
func (m NoRoutingIDs) SetRoutingType(v enum.RoutingType) {
	m.Set(field.NewRoutingType(v))
}

// SetRoutingID sets RoutingID, Tag 217.
func (m NoRoutingIDs) SetRoutingID(v string) {
	m.Set(field.NewRoutingID(v))
}

// GetRoutingType gets RoutingType, Tag 216.
func (m NoRoutingIDs) GetRoutingType() (v enum.RoutingType, err quickfix.MessageRejectError) {
	var f field.RoutingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRoutingID gets RoutingID, Tag 217.
func (m NoRoutingIDs) GetRoutingID() (v string, err quickfix.MessageRejectError) {
	var f field.RoutingIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRoutingType returns true if RoutingType is present, Tag 216.
func (m NoRoutingIDs) HasRoutingType() bool {
	return m.Has(tag.RoutingType)
}

// HasRoutingID returns true if RoutingID is present, Tag 217.
func (m NoRoutingIDs) HasRoutingID() bool {
	return m.Has(tag.RoutingID)
}

// NoRoutingIDsRepeatingGroup is a repeating group, Tag 215.
type NoRoutingIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRoutingIDsRepeatingGroup returns an initialized, NoRoutingIDsRepeatingGroup.
func NewNoRoutingIDsRepeatingGroup() NoRoutingIDsRepeatingGroup {
	return NoRoutingIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRoutingIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RoutingType), quickfix.GroupElement(tag.RoutingID)})}
}

// Add create and append a new NoRoutingIDs to this group.
func (m NoRoutingIDsRepeatingGroup) Add() NoRoutingIDs {
	g := m.RepeatingGroup.Add()
	return NoRoutingIDs{g}
}

// Get returns the ith NoRoutingIDs in the NoRoutingIDsRepeatinGroup.
func (m NoRoutingIDsRepeatingGroup) Get(i int) NoRoutingIDs {
	return NoRoutingIDs{m.RepeatingGroup.Get(i)}
}

// NoMDEntries is a repeating group element, Tag 268.
type NoMDEntries struct {
	*quickfix.Group
}

// SetMDUpdateAction sets MDUpdateAction, Tag 279.
func (m NoMDEntries) SetMDUpdateAction(v enum.MDUpdateAction) {
	m.Set(field.NewMDUpdateAction(v))
}

// SetDeleteReason sets DeleteReason, Tag 285.
func (m NoMDEntries) SetDeleteReason(v enum.DeleteReason) {
	m.Set(field.NewDeleteReason(v))
}

// SetMDEntryType sets MDEntryType, Tag 269.
func (m NoMDEntries) SetMDEntryType(v enum.MDEntryType) {
	m.Set(field.NewMDEntryType(v))
}

// SetMDEntryID sets MDEntryID, Tag 278.
func (m NoMDEntries) SetMDEntryID(v string) {
	m.Set(field.NewMDEntryID(v))
}

// SetMDEntryRefID sets MDEntryRefID, Tag 280.
func (m NoMDEntries) SetMDEntryRefID(v string) {
	m.Set(field.NewMDEntryRefID(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m NoMDEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m NoMDEntries) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoMDEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoMDEntries) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m NoMDEntries) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m NoMDEntries) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m NoMDEntries) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m NoMDEntries) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m NoMDEntries) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m NoMDEntries) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m NoMDEntries) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m NoMDEntries) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m NoMDEntries) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m NoMDEntries) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m NoMDEntries) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m NoMDEntries) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m NoMDEntries) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m NoMDEntries) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m NoMDEntries) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m NoMDEntries) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m NoMDEntries) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m NoMDEntries) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m NoMDEntries) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m NoMDEntries) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m NoMDEntries) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m NoMDEntries) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m NoMDEntries) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m NoMDEntries) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m NoMDEntries) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m NoMDEntries) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m NoMDEntries) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m NoMDEntries) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m NoMDEntries) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m NoMDEntries) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m NoMDEntries) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetPool sets Pool, Tag 691.
func (m NoMDEntries) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m NoMDEntries) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m NoMDEntries) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m NoMDEntries) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetNoEvents sets NoEvents, Tag 864.
func (m NoMDEntries) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m NoMDEntries) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m NoMDEntries) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetSecurityStatus sets SecurityStatus, Tag 965.
func (m NoMDEntries) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

// SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966.
func (m NoMDEntries) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

// SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049.
func (m NoMDEntries) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

// SetStrikeMultiplier sets StrikeMultiplier, Tag 967.
func (m NoMDEntries) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

// SetStrikeValue sets StrikeValue, Tag 968.
func (m NoMDEntries) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

// SetMinPriceIncrement sets MinPriceIncrement, Tag 969.
func (m NoMDEntries) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

// SetPositionLimit sets PositionLimit, Tag 970.
func (m NoMDEntries) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

// SetNTPositionLimit sets NTPositionLimit, Tag 971.
func (m NoMDEntries) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

// SetNoInstrumentParties sets NoInstrumentParties, Tag 1018.
func (m NoMDEntries) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnitOfMeasure sets UnitOfMeasure, Tag 996.
func (m NoMDEntries) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

// SetTimeUnit sets TimeUnit, Tag 997.
func (m NoMDEntries) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

// SetMaturityTime sets MaturityTime, Tag 1079.
func (m NoMDEntries) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

// SetSecurityGroup sets SecurityGroup, Tag 1151.
func (m NoMDEntries) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

// SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146.
func (m NoMDEntries) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

// SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147.
func (m NoMDEntries) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

// SetSecurityXMLLen sets SecurityXMLLen, Tag 1184.
func (m NoMDEntries) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

// SetSecurityXML sets SecurityXML, Tag 1185.
func (m NoMDEntries) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

// SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186.
func (m NoMDEntries) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

// SetProductComplex sets ProductComplex, Tag 1227.
func (m NoMDEntries) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

// SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191.
func (m NoMDEntries) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

// SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192.
func (m NoMDEntries) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

// SetSettlMethod sets SettlMethod, Tag 1193.
func (m NoMDEntries) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

// SetExerciseStyle sets ExerciseStyle, Tag 1194.
func (m NoMDEntries) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

// SetOptPayoutAmount sets OptPayoutAmount, Tag 1195.
func (m NoMDEntries) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

// SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196.
func (m NoMDEntries) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

// SetListMethod sets ListMethod, Tag 1198.
func (m NoMDEntries) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

// SetCapPrice sets CapPrice, Tag 1199.
func (m NoMDEntries) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

// SetFloorPrice sets FloorPrice, Tag 1200.
func (m NoMDEntries) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m NoMDEntries) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetFlexibleIndicator sets FlexibleIndicator, Tag 1244.
func (m NoMDEntries) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

// SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242.
func (m NoMDEntries) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

// SetValuationMethod sets ValuationMethod, Tag 1197.
func (m NoMDEntries) SetValuationMethod(v enum.ValuationMethod) {
	m.Set(field.NewValuationMethod(v))
}

// SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435.
func (m NoMDEntries) SetContractMultiplierUnit(v enum.ContractMultiplierUnit) {
	m.Set(field.NewContractMultiplierUnit(v))
}

// SetFlowScheduleType sets FlowScheduleType, Tag 1439.
func (m NoMDEntries) SetFlowScheduleType(v enum.FlowScheduleType) {
	m.Set(field.NewFlowScheduleType(v))
}

// SetRestructuringType sets RestructuringType, Tag 1449.
func (m NoMDEntries) SetRestructuringType(v enum.RestructuringType) {
	m.Set(field.NewRestructuringType(v))
}

// SetSeniority sets Seniority, Tag 1450.
func (m NoMDEntries) SetSeniority(v enum.Seniority) {
	m.Set(field.NewSeniority(v))
}

// SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451.
func (m NoMDEntries) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

// SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m NoMDEntries) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

// SetAttachmentPoint sets AttachmentPoint, Tag 1457.
func (m NoMDEntries) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

// SetDetachmentPoint sets DetachmentPoint, Tag 1458.
func (m NoMDEntries) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

// SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478.
func (m NoMDEntries) SetStrikePriceDeterminationMethod(v enum.StrikePriceDeterminationMethod) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

// SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479.
func (m NoMDEntries) SetStrikePriceBoundaryMethod(v enum.StrikePriceBoundaryMethod) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

// SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480.
func (m NoMDEntries) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

// SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m NoMDEntries) SetUnderlyingPriceDeterminationMethod(v enum.UnderlyingPriceDeterminationMethod) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

// SetOptPayoutType sets OptPayoutType, Tag 1482.
func (m NoMDEntries) SetOptPayoutType(v enum.OptPayoutType) {
	m.Set(field.NewOptPayoutType(v))
}

// SetNoComplexEvents sets NoComplexEvents, Tag 1483.
func (m NoMDEntries) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoUnderlyings sets NoUnderlyings, Tag 711.
func (m NoMDEntries) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoLegs sets NoLegs, Tag 555.
func (m NoMDEntries) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

// SetFinancialStatus sets FinancialStatus, Tag 291.
func (m NoMDEntries) SetFinancialStatus(v enum.FinancialStatus) {
	m.Set(field.NewFinancialStatus(v))
}

// SetCorporateAction sets CorporateAction, Tag 292.
func (m NoMDEntries) SetCorporateAction(v enum.CorporateAction) {
	m.Set(field.NewCorporateAction(v))
}

// SetMDEntryPx sets MDEntryPx, Tag 270.
func (m NoMDEntries) SetMDEntryPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntryPx(value, scale))
}

// SetCurrency sets Currency, Tag 15.
func (m NoMDEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetMDEntrySize sets MDEntrySize, Tag 271.
func (m NoMDEntries) SetMDEntrySize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntrySize(value, scale))
}

// SetMDEntryDate sets MDEntryDate, Tag 272.
func (m NoMDEntries) SetMDEntryDate(v string) {
	m.Set(field.NewMDEntryDate(v))
}

// SetMDEntryTime sets MDEntryTime, Tag 273.
func (m NoMDEntries) SetMDEntryTime(v string) {
	m.Set(field.NewMDEntryTime(v))
}

// SetTickDirection sets TickDirection, Tag 274.
func (m NoMDEntries) SetTickDirection(v enum.TickDirection) {
	m.Set(field.NewTickDirection(v))
}

// SetMDMkt sets MDMkt, Tag 275.
func (m NoMDEntries) SetMDMkt(v string) {
	m.Set(field.NewMDMkt(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m NoMDEntries) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m NoMDEntries) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetQuoteCondition sets QuoteCondition, Tag 276.
func (m NoMDEntries) SetQuoteCondition(v enum.QuoteCondition) {
	m.Set(field.NewQuoteCondition(v))
}

// SetTradeCondition sets TradeCondition, Tag 277.
func (m NoMDEntries) SetTradeCondition(v enum.TradeCondition) {
	m.Set(field.NewTradeCondition(v))
}

// SetMDEntryOriginator sets MDEntryOriginator, Tag 282.
func (m NoMDEntries) SetMDEntryOriginator(v string) {
	m.Set(field.NewMDEntryOriginator(v))
}

// SetLocationID sets LocationID, Tag 283.
func (m NoMDEntries) SetLocationID(v string) {
	m.Set(field.NewLocationID(v))
}

// SetDeskID sets DeskID, Tag 284.
func (m NoMDEntries) SetDeskID(v string) {
	m.Set(field.NewDeskID(v))
}

// SetOpenCloseSettlFlag sets OpenCloseSettlFlag, Tag 286.
func (m NoMDEntries) SetOpenCloseSettlFlag(v enum.OpenCloseSettlFlag) {
	m.Set(field.NewOpenCloseSettlFlag(v))
}

// SetTimeInForce sets TimeInForce, Tag 59.
func (m NoMDEntries) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

// SetExpireDate sets ExpireDate, Tag 432.
func (m NoMDEntries) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

// SetExpireTime sets ExpireTime, Tag 126.
func (m NoMDEntries) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

// SetMinQty sets MinQty, Tag 110.
func (m NoMDEntries) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

// SetExecInst sets ExecInst, Tag 18.
func (m NoMDEntries) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

// SetSellerDays sets SellerDays, Tag 287.
func (m NoMDEntries) SetSellerDays(v int) {
	m.Set(field.NewSellerDays(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m NoMDEntries) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetQuoteEntryID sets QuoteEntryID, Tag 299.
func (m NoMDEntries) SetQuoteEntryID(v string) {
	m.Set(field.NewQuoteEntryID(v))
}

// SetMDEntryBuyer sets MDEntryBuyer, Tag 288.
func (m NoMDEntries) SetMDEntryBuyer(v string) {
	m.Set(field.NewMDEntryBuyer(v))
}

// SetMDEntrySeller sets MDEntrySeller, Tag 289.
func (m NoMDEntries) SetMDEntrySeller(v string) {
	m.Set(field.NewMDEntrySeller(v))
}

// SetNumberOfOrders sets NumberOfOrders, Tag 346.
func (m NoMDEntries) SetNumberOfOrders(v int) {
	m.Set(field.NewNumberOfOrders(v))
}

// SetMDEntryPositionNo sets MDEntryPositionNo, Tag 290.
func (m NoMDEntries) SetMDEntryPositionNo(v int) {
	m.Set(field.NewMDEntryPositionNo(v))
}

// SetScope sets Scope, Tag 546.
func (m NoMDEntries) SetScope(v enum.Scope) {
	m.Set(field.NewScope(v))
}

// SetPriceDelta sets PriceDelta, Tag 811.
func (m NoMDEntries) SetPriceDelta(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceDelta(value, scale))
}

// SetNetChgPrevDay sets NetChgPrevDay, Tag 451.
func (m NoMDEntries) SetNetChgPrevDay(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetChgPrevDay(value, scale))
}

// SetText sets Text, Tag 58.
func (m NoMDEntries) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m NoMDEntries) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m NoMDEntries) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetOrderCapacity sets OrderCapacity, Tag 528.
func (m NoMDEntries) SetOrderCapacity(v enum.OrderCapacity) {
	m.Set(field.NewOrderCapacity(v))
}

// SetMDOriginType sets MDOriginType, Tag 1024.
func (m NoMDEntries) SetMDOriginType(v enum.MDOriginType) {
	m.Set(field.NewMDOriginType(v))
}

// SetHighPx sets HighPx, Tag 332.
func (m NoMDEntries) SetHighPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewHighPx(value, scale))
}

// SetLowPx sets LowPx, Tag 333.
func (m NoMDEntries) SetLowPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLowPx(value, scale))
}

// SetTradeVolume sets TradeVolume, Tag 1020.
func (m NoMDEntries) SetTradeVolume(value decimal.Decimal, scale int32) {
	m.Set(field.NewTradeVolume(value, scale))
}

// SetSettlType sets SettlType, Tag 63.
func (m NoMDEntries) SetSettlType(v enum.SettlType) {
	m.Set(field.NewSettlType(v))
}

// SetSettlDate sets SettlDate, Tag 64.
func (m NoMDEntries) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

// SetMDQuoteType sets MDQuoteType, Tag 1070.
func (m NoMDEntries) SetMDQuoteType(v enum.MDQuoteType) {
	m.Set(field.NewMDQuoteType(v))
}

// SetRptSeq sets RptSeq, Tag 83.
func (m NoMDEntries) SetRptSeq(v int) {
	m.Set(field.NewRptSeq(v))
}

// SetDealingCapacity sets DealingCapacity, Tag 1048.
func (m NoMDEntries) SetDealingCapacity(v enum.DealingCapacity) {
	m.Set(field.NewDealingCapacity(v))
}

// SetMDEntrySpotRate sets MDEntrySpotRate, Tag 1026.
func (m NoMDEntries) SetMDEntrySpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntrySpotRate(value, scale))
}

// SetMDEntryForwardPoints sets MDEntryForwardPoints, Tag 1027.
func (m NoMDEntries) SetMDEntryForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntryForwardPoints(value, scale))
}

// SetMDPriceLevel sets MDPriceLevel, Tag 1023.
func (m NoMDEntries) SetMDPriceLevel(v int) {
	m.Set(field.NewMDPriceLevel(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m NoMDEntries) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetSecondaryOrderID sets SecondaryOrderID, Tag 198.
func (m NoMDEntries) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

// SetOrdType sets OrdType, Tag 40.
func (m NoMDEntries) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

// SetMDSubBookType sets MDSubBookType, Tag 1173.
func (m NoMDEntries) SetMDSubBookType(v int) {
	m.Set(field.NewMDSubBookType(v))
}

// SetMarketDepth sets MarketDepth, Tag 264.
func (m NoMDEntries) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

// SetPriceType sets PriceType, Tag 423.
func (m NoMDEntries) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

// SetYieldType sets YieldType, Tag 235.
func (m NoMDEntries) SetYieldType(v enum.YieldType) {
	m.Set(field.NewYieldType(v))
}

// SetYield sets Yield, Tag 236.
func (m NoMDEntries) SetYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewYield(value, scale))
}

// SetYieldCalcDate sets YieldCalcDate, Tag 701.
func (m NoMDEntries) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

// SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696.
func (m NoMDEntries) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

// SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697.
func (m NoMDEntries) SetYieldRedemptionPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewYieldRedemptionPrice(value, scale))
}

// SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698.
func (m NoMDEntries) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

// SetSpread sets Spread, Tag 218.
func (m NoMDEntries) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

// SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220.
func (m NoMDEntries) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

// SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221.
func (m NoMDEntries) SetBenchmarkCurveName(v enum.BenchmarkCurveName) {
	m.Set(field.NewBenchmarkCurveName(v))
}

// SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222.
func (m NoMDEntries) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

// SetBenchmarkPrice sets BenchmarkPrice, Tag 662.
func (m NoMDEntries) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

// SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663.
func (m NoMDEntries) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

// SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699.
func (m NoMDEntries) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

// SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761.
func (m NoMDEntries) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

// SetNoOfSecSizes sets NoOfSecSizes, Tag 1177.
func (m NoMDEntries) SetNoOfSecSizes(f NoOfSecSizesRepeatingGroup) {
	m.SetGroup(f)
}

// SetLotType sets LotType, Tag 1093.
func (m NoMDEntries) SetLotType(v enum.LotType) {
	m.Set(field.NewLotType(v))
}

// SetSecurityTradingStatus sets SecurityTradingStatus, Tag 326.
func (m NoMDEntries) SetSecurityTradingStatus(v enum.SecurityTradingStatus) {
	m.Set(field.NewSecurityTradingStatus(v))
}

// SetHaltReasonInt sets HaltReasonInt, Tag 327.
func (m NoMDEntries) SetHaltReasonInt(v enum.HaltReasonInt) {
	m.Set(field.NewHaltReasonInt(v))
}

// SetTrdType sets TrdType, Tag 828.
func (m NoMDEntries) SetTrdType(v enum.TrdType) {
	m.Set(field.NewTrdType(v))
}

// SetMatchType sets MatchType, Tag 574.
func (m NoMDEntries) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

// SetTradeID sets TradeID, Tag 1003.
func (m NoMDEntries) SetTradeID(v string) {
	m.Set(field.NewTradeID(v))
}

// SetTransBkdTime sets TransBkdTime, Tag 483.
func (m NoMDEntries) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m NoMDEntries) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetNoStatsIndicators sets NoStatsIndicators, Tag 1175.
func (m NoMDEntries) SetNoStatsIndicators(f NoStatsIndicatorsRepeatingGroup) {
	m.SetGroup(f)
}

// SetSettlCurrency sets SettlCurrency, Tag 120.
func (m NoMDEntries) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

// SetNoRateSources sets NoRateSources, Tag 1445.
func (m NoMDEntries) SetNoRateSources(f NoRateSourcesRepeatingGroup) {
	m.SetGroup(f)
}

// SetFirstPx sets FirstPx, Tag 1025.
func (m NoMDEntries) SetFirstPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewFirstPx(value, scale))
}

// SetLastPx sets LastPx, Tag 31.
func (m NoMDEntries) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

// SetMDStreamID sets MDStreamID, Tag 1500.
func (m NoMDEntries) SetMDStreamID(v string) {
	m.Set(field.NewMDStreamID(v))
}

// GetMDUpdateAction gets MDUpdateAction, Tag 279.
func (m NoMDEntries) GetMDUpdateAction() (v enum.MDUpdateAction, err quickfix.MessageRejectError) {
	var f field.MDUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeleteReason gets DeleteReason, Tag 285.
func (m NoMDEntries) GetDeleteReason() (v enum.DeleteReason, err quickfix.MessageRejectError) {
	var f field.DeleteReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryType gets MDEntryType, Tag 269.
func (m NoMDEntries) GetMDEntryType() (v enum.MDEntryType, err quickfix.MessageRejectError) {
	var f field.MDEntryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryID gets MDEntryID, Tag 278.
func (m NoMDEntries) GetMDEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryRefID gets MDEntryRefID, Tag 280.
func (m NoMDEntries) GetMDEntryRefID() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m NoMDEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m NoMDEntries) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoMDEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoMDEntries) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m NoMDEntries) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m NoMDEntries) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m NoMDEntries) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m NoMDEntries) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m NoMDEntries) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m NoMDEntries) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m NoMDEntries) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m NoMDEntries) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m NoMDEntries) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m NoMDEntries) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m NoMDEntries) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m NoMDEntries) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m NoMDEntries) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m NoMDEntries) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m NoMDEntries) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m NoMDEntries) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m NoMDEntries) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m NoMDEntries) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m NoMDEntries) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m NoMDEntries) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m NoMDEntries) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m NoMDEntries) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m NoMDEntries) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m NoMDEntries) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m NoMDEntries) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m NoMDEntries) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m NoMDEntries) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m NoMDEntries) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m NoMDEntries) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m NoMDEntries) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m NoMDEntries) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m NoMDEntries) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m NoMDEntries) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m NoMDEntries) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m NoMDEntries) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864.
func (m NoMDEntries) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m NoMDEntries) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m NoMDEntries) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatus gets SecurityStatus, Tag 965.
func (m NoMDEntries) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966.
func (m NoMDEntries) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049.
func (m NoMDEntries) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeMultiplier gets StrikeMultiplier, Tag 967.
func (m NoMDEntries) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeValue gets StrikeValue, Tag 968.
func (m NoMDEntries) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrement gets MinPriceIncrement, Tag 969.
func (m NoMDEntries) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionLimit gets PositionLimit, Tag 970.
func (m NoMDEntries) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNTPositionLimit gets NTPositionLimit, Tag 971.
func (m NoMDEntries) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentParties gets NoInstrumentParties, Tag 1018.
func (m NoMDEntries) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnitOfMeasure gets UnitOfMeasure, Tag 996.
func (m NoMDEntries) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeUnit gets TimeUnit, Tag 997.
func (m NoMDEntries) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityTime gets MaturityTime, Tag 1079.
func (m NoMDEntries) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityGroup gets SecurityGroup, Tag 1151.
func (m NoMDEntries) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146.
func (m NoMDEntries) GetMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147.
func (m NoMDEntries) GetUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLLen gets SecurityXMLLen, Tag 1184.
func (m NoMDEntries) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXML gets SecurityXML, Tag 1185.
func (m NoMDEntries) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186.
func (m NoMDEntries) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetProductComplex gets ProductComplex, Tag 1227.
func (m NoMDEntries) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191.
func (m NoMDEntries) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192.
func (m NoMDEntries) GetPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlMethod gets SettlMethod, Tag 1193.
func (m NoMDEntries) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExerciseStyle gets ExerciseStyle, Tag 1194.
func (m NoMDEntries) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutAmount gets OptPayoutAmount, Tag 1195.
func (m NoMDEntries) GetOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196.
func (m NoMDEntries) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListMethod gets ListMethod, Tag 1198.
func (m NoMDEntries) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCapPrice gets CapPrice, Tag 1199.
func (m NoMDEntries) GetCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFloorPrice gets FloorPrice, Tag 1200.
func (m NoMDEntries) GetFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m NoMDEntries) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexibleIndicator gets FlexibleIndicator, Tag 1244.
func (m NoMDEntries) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242.
func (m NoMDEntries) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValuationMethod gets ValuationMethod, Tag 1197.
func (m NoMDEntries) GetValuationMethod() (v enum.ValuationMethod, err quickfix.MessageRejectError) {
	var f field.ValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435.
func (m NoMDEntries) GetContractMultiplierUnit() (v enum.ContractMultiplierUnit, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlowScheduleType gets FlowScheduleType, Tag 1439.
func (m NoMDEntries) GetFlowScheduleType() (v enum.FlowScheduleType, err quickfix.MessageRejectError) {
	var f field.FlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRestructuringType gets RestructuringType, Tag 1449.
func (m NoMDEntries) GetRestructuringType() (v enum.RestructuringType, err quickfix.MessageRejectError) {
	var f field.RestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSeniority gets Seniority, Tag 1450.
func (m NoMDEntries) GetSeniority() (v enum.Seniority, err quickfix.MessageRejectError) {
	var f field.SeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451.
func (m NoMDEntries) GetNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m NoMDEntries) GetOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAttachmentPoint gets AttachmentPoint, Tag 1457.
func (m NoMDEntries) GetAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDetachmentPoint gets DetachmentPoint, Tag 1458.
func (m NoMDEntries) GetDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478.
func (m NoMDEntries) GetStrikePriceDeterminationMethod() (v enum.StrikePriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479.
func (m NoMDEntries) GetStrikePriceBoundaryMethod() (v enum.StrikePriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480.
func (m NoMDEntries) GetStrikePriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m NoMDEntries) GetUnderlyingPriceDeterminationMethod() (v enum.UnderlyingPriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutType gets OptPayoutType, Tag 1482.
func (m NoMDEntries) GetOptPayoutType() (v enum.OptPayoutType, err quickfix.MessageRejectError) {
	var f field.OptPayoutTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEvents gets NoComplexEvents, Tag 1483.
func (m NoMDEntries) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoUnderlyings gets NoUnderlyings, Tag 711.
func (m NoMDEntries) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoLegs gets NoLegs, Tag 555.
func (m NoMDEntries) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetFinancialStatus gets FinancialStatus, Tag 291.
func (m NoMDEntries) GetFinancialStatus() (v enum.FinancialStatus, err quickfix.MessageRejectError) {
	var f field.FinancialStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCorporateAction gets CorporateAction, Tag 292.
func (m NoMDEntries) GetCorporateAction() (v enum.CorporateAction, err quickfix.MessageRejectError) {
	var f field.CorporateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryPx gets MDEntryPx, Tag 270.
func (m NoMDEntries) GetMDEntryPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntryPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m NoMDEntries) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntrySize gets MDEntrySize, Tag 271.
func (m NoMDEntries) GetMDEntrySize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntrySizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryDate gets MDEntryDate, Tag 272.
func (m NoMDEntries) GetMDEntryDate() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryTime gets MDEntryTime, Tag 273.
func (m NoMDEntries) GetMDEntryTime() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTickDirection gets TickDirection, Tag 274.
func (m NoMDEntries) GetTickDirection() (v enum.TickDirection, err quickfix.MessageRejectError) {
	var f field.TickDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDMkt gets MDMkt, Tag 275.
func (m NoMDEntries) GetMDMkt() (v string, err quickfix.MessageRejectError) {
	var f field.MDMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m NoMDEntries) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m NoMDEntries) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteCondition gets QuoteCondition, Tag 276.
func (m NoMDEntries) GetQuoteCondition() (v enum.QuoteCondition, err quickfix.MessageRejectError) {
	var f field.QuoteConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeCondition gets TradeCondition, Tag 277.
func (m NoMDEntries) GetTradeCondition() (v enum.TradeCondition, err quickfix.MessageRejectError) {
	var f field.TradeConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryOriginator gets MDEntryOriginator, Tag 282.
func (m NoMDEntries) GetMDEntryOriginator() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryOriginatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocationID gets LocationID, Tag 283.
func (m NoMDEntries) GetLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.LocationIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeskID gets DeskID, Tag 284.
func (m NoMDEntries) GetDeskID() (v string, err quickfix.MessageRejectError) {
	var f field.DeskIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOpenCloseSettlFlag gets OpenCloseSettlFlag, Tag 286.
func (m NoMDEntries) GetOpenCloseSettlFlag() (v enum.OpenCloseSettlFlag, err quickfix.MessageRejectError) {
	var f field.OpenCloseSettlFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeInForce gets TimeInForce, Tag 59.
func (m NoMDEntries) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpireDate gets ExpireDate, Tag 432.
func (m NoMDEntries) GetExpireDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExpireDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExpireTime gets ExpireTime, Tag 126.
func (m NoMDEntries) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinQty gets MinQty, Tag 110.
func (m NoMDEntries) GetMinQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExecInst gets ExecInst, Tag 18.
func (m NoMDEntries) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSellerDays gets SellerDays, Tag 287.
func (m NoMDEntries) GetSellerDays() (v int, err quickfix.MessageRejectError) {
	var f field.SellerDaysField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m NoMDEntries) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteEntryID gets QuoteEntryID, Tag 299.
func (m NoMDEntries) GetQuoteEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryBuyer gets MDEntryBuyer, Tag 288.
func (m NoMDEntries) GetMDEntryBuyer() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryBuyerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntrySeller gets MDEntrySeller, Tag 289.
func (m NoMDEntries) GetMDEntrySeller() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntrySellerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNumberOfOrders gets NumberOfOrders, Tag 346.
func (m NoMDEntries) GetNumberOfOrders() (v int, err quickfix.MessageRejectError) {
	var f field.NumberOfOrdersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryPositionNo gets MDEntryPositionNo, Tag 290.
func (m NoMDEntries) GetMDEntryPositionNo() (v int, err quickfix.MessageRejectError) {
	var f field.MDEntryPositionNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetScope gets Scope, Tag 546.
func (m NoMDEntries) GetScope() (v enum.Scope, err quickfix.MessageRejectError) {
	var f field.ScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceDelta gets PriceDelta, Tag 811.
func (m NoMDEntries) GetPriceDelta() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceDeltaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNetChgPrevDay gets NetChgPrevDay, Tag 451.
func (m NoMDEntries) GetNetChgPrevDay() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NetChgPrevDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m NoMDEntries) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m NoMDEntries) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m NoMDEntries) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderCapacity gets OrderCapacity, Tag 528.
func (m NoMDEntries) GetOrderCapacity() (v enum.OrderCapacity, err quickfix.MessageRejectError) {
	var f field.OrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDOriginType gets MDOriginType, Tag 1024.
func (m NoMDEntries) GetMDOriginType() (v enum.MDOriginType, err quickfix.MessageRejectError) {
	var f field.MDOriginTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHighPx gets HighPx, Tag 332.
func (m NoMDEntries) GetHighPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.HighPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLowPx gets LowPx, Tag 333.
func (m NoMDEntries) GetLowPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LowPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeVolume gets TradeVolume, Tag 1020.
func (m NoMDEntries) GetTradeVolume() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TradeVolumeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlType gets SettlType, Tag 63.
func (m NoMDEntries) GetSettlType() (v enum.SettlType, err quickfix.MessageRejectError) {
	var f field.SettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlDate gets SettlDate, Tag 64.
func (m NoMDEntries) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDQuoteType gets MDQuoteType, Tag 1070.
func (m NoMDEntries) GetMDQuoteType() (v enum.MDQuoteType, err quickfix.MessageRejectError) {
	var f field.MDQuoteTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRptSeq gets RptSeq, Tag 83.
func (m NoMDEntries) GetRptSeq() (v int, err quickfix.MessageRejectError) {
	var f field.RptSeqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDealingCapacity gets DealingCapacity, Tag 1048.
func (m NoMDEntries) GetDealingCapacity() (v enum.DealingCapacity, err quickfix.MessageRejectError) {
	var f field.DealingCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntrySpotRate gets MDEntrySpotRate, Tag 1026.
func (m NoMDEntries) GetMDEntrySpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntrySpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDEntryForwardPoints gets MDEntryForwardPoints, Tag 1027.
func (m NoMDEntries) GetMDEntryForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntryForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDPriceLevel gets MDPriceLevel, Tag 1023.
func (m NoMDEntries) GetMDPriceLevel() (v int, err quickfix.MessageRejectError) {
	var f field.MDPriceLevelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m NoMDEntries) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSecondaryOrderID gets SecondaryOrderID, Tag 198.
func (m NoMDEntries) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrdType gets OrdType, Tag 40.
func (m NoMDEntries) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDSubBookType gets MDSubBookType, Tag 1173.
func (m NoMDEntries) GetMDSubBookType() (v int, err quickfix.MessageRejectError) {
	var f field.MDSubBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarketDepth gets MarketDepth, Tag 264.
func (m NoMDEntries) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceType gets PriceType, Tag 423.
func (m NoMDEntries) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldType gets YieldType, Tag 235.
func (m NoMDEntries) GetYieldType() (v enum.YieldType, err quickfix.MessageRejectError) {
	var f field.YieldTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYield gets Yield, Tag 236.
func (m NoMDEntries) GetYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldCalcDate gets YieldCalcDate, Tag 701.
func (m NoMDEntries) GetYieldCalcDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldCalcDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696.
func (m NoMDEntries) GetYieldRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697.
func (m NoMDEntries) GetYieldRedemptionPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698.
func (m NoMDEntries) GetYieldRedemptionPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.YieldRedemptionPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSpread gets Spread, Tag 218.
func (m NoMDEntries) GetSpread() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SpreadField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220.
func (m NoMDEntries) GetBenchmarkCurveCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221.
func (m NoMDEntries) GetBenchmarkCurveName() (v enum.BenchmarkCurveName, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222.
func (m NoMDEntries) GetBenchmarkCurvePoint() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurvePointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPrice gets BenchmarkPrice, Tag 662.
func (m NoMDEntries) GetBenchmarkPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663.
func (m NoMDEntries) GetBenchmarkPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699.
func (m NoMDEntries) GetBenchmarkSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761.
func (m NoMDEntries) GetBenchmarkSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoOfSecSizes gets NoOfSecSizes, Tag 1177.
func (m NoMDEntries) GetNoOfSecSizes() (NoOfSecSizesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOfSecSizesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLotType gets LotType, Tag 1093.
func (m NoMDEntries) GetLotType() (v enum.LotType, err quickfix.MessageRejectError) {
	var f field.LotTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityTradingStatus gets SecurityTradingStatus, Tag 326.
func (m NoMDEntries) GetSecurityTradingStatus() (v enum.SecurityTradingStatus, err quickfix.MessageRejectError) {
	var f field.SecurityTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetHaltReasonInt gets HaltReasonInt, Tag 327.
func (m NoMDEntries) GetHaltReasonInt() (v enum.HaltReasonInt, err quickfix.MessageRejectError) {
	var f field.HaltReasonIntField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTrdType gets TrdType, Tag 828.
func (m NoMDEntries) GetTrdType() (v enum.TrdType, err quickfix.MessageRejectError) {
	var f field.TrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchType gets MatchType, Tag 574.
func (m NoMDEntries) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeID gets TradeID, Tag 1003.
func (m NoMDEntries) GetTradeID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransBkdTime gets TransBkdTime, Tag 483.
func (m NoMDEntries) GetTransBkdTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransBkdTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m NoMDEntries) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoStatsIndicators gets NoStatsIndicators, Tag 1175.
func (m NoMDEntries) GetNoStatsIndicators() (NoStatsIndicatorsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStatsIndicatorsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetSettlCurrency gets SettlCurrency, Tag 120.
func (m NoMDEntries) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRateSources gets NoRateSources, Tag 1445.
func (m NoMDEntries) GetNoRateSources() (NoRateSourcesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRateSourcesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetFirstPx gets FirstPx, Tag 1025.
func (m NoMDEntries) GetFirstPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FirstPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLastPx gets LastPx, Tag 31.
func (m NoMDEntries) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDStreamID gets MDStreamID, Tag 1500.
func (m NoMDEntries) GetMDStreamID() (v string, err quickfix.MessageRejectError) {
	var f field.MDStreamIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMDUpdateAction returns true if MDUpdateAction is present, Tag 279.
func (m NoMDEntries) HasMDUpdateAction() bool {
	return m.Has(tag.MDUpdateAction)
}

// HasDeleteReason returns true if DeleteReason is present, Tag 285.
func (m NoMDEntries) HasDeleteReason() bool {
	return m.Has(tag.DeleteReason)
}

// HasMDEntryType returns true if MDEntryType is present, Tag 269.
func (m NoMDEntries) HasMDEntryType() bool {
	return m.Has(tag.MDEntryType)
}

// HasMDEntryID returns true if MDEntryID is present, Tag 278.
func (m NoMDEntries) HasMDEntryID() bool {
	return m.Has(tag.MDEntryID)
}

// HasMDEntryRefID returns true if MDEntryRefID is present, Tag 280.
func (m NoMDEntries) HasMDEntryRefID() bool {
	return m.Has(tag.MDEntryRefID)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoMDEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m NoMDEntries) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoMDEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoMDEntries) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m NoMDEntries) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m NoMDEntries) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m NoMDEntries) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m NoMDEntries) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m NoMDEntries) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m NoMDEntries) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m NoMDEntries) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m NoMDEntries) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m NoMDEntries) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m NoMDEntries) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m NoMDEntries) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m NoMDEntries) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m NoMDEntries) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m NoMDEntries) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m NoMDEntries) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m NoMDEntries) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m NoMDEntries) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m NoMDEntries) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m NoMDEntries) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m NoMDEntries) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m NoMDEntries) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m NoMDEntries) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m NoMDEntries) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m NoMDEntries) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m NoMDEntries) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m NoMDEntries) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m NoMDEntries) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m NoMDEntries) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m NoMDEntries) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m NoMDEntries) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m NoMDEntries) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasPool returns true if Pool is present, Tag 691.
func (m NoMDEntries) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m NoMDEntries) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m NoMDEntries) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m NoMDEntries) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m NoMDEntries) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m NoMDEntries) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m NoMDEntries) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasSecurityStatus returns true if SecurityStatus is present, Tag 965.
func (m NoMDEntries) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

// HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966.
func (m NoMDEntries) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

// HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049.
func (m NoMDEntries) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

// HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967.
func (m NoMDEntries) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

// HasStrikeValue returns true if StrikeValue is present, Tag 968.
func (m NoMDEntries) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

// HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969.
func (m NoMDEntries) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

// HasPositionLimit returns true if PositionLimit is present, Tag 970.
func (m NoMDEntries) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

// HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971.
func (m NoMDEntries) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

// HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018.
func (m NoMDEntries) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

// HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996.
func (m NoMDEntries) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

// HasTimeUnit returns true if TimeUnit is present, Tag 997.
func (m NoMDEntries) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

// HasMaturityTime returns true if MaturityTime is present, Tag 1079.
func (m NoMDEntries) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

// HasSecurityGroup returns true if SecurityGroup is present, Tag 1151.
func (m NoMDEntries) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

// HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146.
func (m NoMDEntries) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

// HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147.
func (m NoMDEntries) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

// HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184.
func (m NoMDEntries) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

// HasSecurityXML returns true if SecurityXML is present, Tag 1185.
func (m NoMDEntries) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

// HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186.
func (m NoMDEntries) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

// HasProductComplex returns true if ProductComplex is present, Tag 1227.
func (m NoMDEntries) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

// HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191.
func (m NoMDEntries) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

// HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192.
func (m NoMDEntries) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

// HasSettlMethod returns true if SettlMethod is present, Tag 1193.
func (m NoMDEntries) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

// HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194.
func (m NoMDEntries) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

// HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195.
func (m NoMDEntries) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

// HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196.
func (m NoMDEntries) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

// HasListMethod returns true if ListMethod is present, Tag 1198.
func (m NoMDEntries) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

// HasCapPrice returns true if CapPrice is present, Tag 1199.
func (m NoMDEntries) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

// HasFloorPrice returns true if FloorPrice is present, Tag 1200.
func (m NoMDEntries) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m NoMDEntries) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244.
func (m NoMDEntries) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

// HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242.
func (m NoMDEntries) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

// HasValuationMethod returns true if ValuationMethod is present, Tag 1197.
func (m NoMDEntries) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

// HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435.
func (m NoMDEntries) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

// HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439.
func (m NoMDEntries) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

// HasRestructuringType returns true if RestructuringType is present, Tag 1449.
func (m NoMDEntries) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

// HasSeniority returns true if Seniority is present, Tag 1450.
func (m NoMDEntries) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

// HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451.
func (m NoMDEntries) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

// HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452.
func (m NoMDEntries) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

// HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457.
func (m NoMDEntries) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

// HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458.
func (m NoMDEntries) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

// HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478.
func (m NoMDEntries) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

// HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479.
func (m NoMDEntries) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

// HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480.
func (m NoMDEntries) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

// HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481.
func (m NoMDEntries) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

// HasOptPayoutType returns true if OptPayoutType is present, Tag 1482.
func (m NoMDEntries) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

// HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483.
func (m NoMDEntries) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
}

// HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711.
func (m NoMDEntries) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

// HasNoLegs returns true if NoLegs is present, Tag 555.
func (m NoMDEntries) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

// HasFinancialStatus returns true if FinancialStatus is present, Tag 291.
func (m NoMDEntries) HasFinancialStatus() bool {
	return m.Has(tag.FinancialStatus)
}

// HasCorporateAction returns true if CorporateAction is present, Tag 292.
func (m NoMDEntries) HasCorporateAction() bool {
	return m.Has(tag.CorporateAction)
}

// HasMDEntryPx returns true if MDEntryPx is present, Tag 270.
func (m NoMDEntries) HasMDEntryPx() bool {
	return m.Has(tag.MDEntryPx)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m NoMDEntries) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasMDEntrySize returns true if MDEntrySize is present, Tag 271.
func (m NoMDEntries) HasMDEntrySize() bool {
	return m.Has(tag.MDEntrySize)
}

// HasMDEntryDate returns true if MDEntryDate is present, Tag 272.
func (m NoMDEntries) HasMDEntryDate() bool {
	return m.Has(tag.MDEntryDate)
}

// HasMDEntryTime returns true if MDEntryTime is present, Tag 273.
func (m NoMDEntries) HasMDEntryTime() bool {
	return m.Has(tag.MDEntryTime)
}

// HasTickDirection returns true if TickDirection is present, Tag 274.
func (m NoMDEntries) HasTickDirection() bool {
	return m.Has(tag.TickDirection)
}

// HasMDMkt returns true if MDMkt is present, Tag 275.
func (m NoMDEntries) HasMDMkt() bool {
	return m.Has(tag.MDMkt)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m NoMDEntries) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m NoMDEntries) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasQuoteCondition returns true if QuoteCondition is present, Tag 276.
func (m NoMDEntries) HasQuoteCondition() bool {
	return m.Has(tag.QuoteCondition)
}

// HasTradeCondition returns true if TradeCondition is present, Tag 277.
func (m NoMDEntries) HasTradeCondition() bool {
	return m.Has(tag.TradeCondition)
}

// HasMDEntryOriginator returns true if MDEntryOriginator is present, Tag 282.
func (m NoMDEntries) HasMDEntryOriginator() bool {
	return m.Has(tag.MDEntryOriginator)
}

// HasLocationID returns true if LocationID is present, Tag 283.
func (m NoMDEntries) HasLocationID() bool {
	return m.Has(tag.LocationID)
}

// HasDeskID returns true if DeskID is present, Tag 284.
func (m NoMDEntries) HasDeskID() bool {
	return m.Has(tag.DeskID)
}

// HasOpenCloseSettlFlag returns true if OpenCloseSettlFlag is present, Tag 286.
func (m NoMDEntries) HasOpenCloseSettlFlag() bool {
	return m.Has(tag.OpenCloseSettlFlag)
}

// HasTimeInForce returns true if TimeInForce is present, Tag 59.
func (m NoMDEntries) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

// HasExpireDate returns true if ExpireDate is present, Tag 432.
func (m NoMDEntries) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

// HasExpireTime returns true if ExpireTime is present, Tag 126.
func (m NoMDEntries) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

// HasMinQty returns true if MinQty is present, Tag 110.
func (m NoMDEntries) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

// HasExecInst returns true if ExecInst is present, Tag 18.
func (m NoMDEntries) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

// HasSellerDays returns true if SellerDays is present, Tag 287.
func (m NoMDEntries) HasSellerDays() bool {
	return m.Has(tag.SellerDays)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m NoMDEntries) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasQuoteEntryID returns true if QuoteEntryID is present, Tag 299.
func (m NoMDEntries) HasQuoteEntryID() bool {
	return m.Has(tag.QuoteEntryID)
}

// HasMDEntryBuyer returns true if MDEntryBuyer is present, Tag 288.
func (m NoMDEntries) HasMDEntryBuyer() bool {
	return m.Has(tag.MDEntryBuyer)
}

// HasMDEntrySeller returns true if MDEntrySeller is present, Tag 289.
func (m NoMDEntries) HasMDEntrySeller() bool {
	return m.Has(tag.MDEntrySeller)
}

// HasNumberOfOrders returns true if NumberOfOrders is present, Tag 346.
func (m NoMDEntries) HasNumberOfOrders() bool {
	return m.Has(tag.NumberOfOrders)
}

// HasMDEntryPositionNo returns true if MDEntryPositionNo is present, Tag 290.
func (m NoMDEntries) HasMDEntryPositionNo() bool {
	return m.Has(tag.MDEntryPositionNo)
}

// HasScope returns true if Scope is present, Tag 546.
func (m NoMDEntries) HasScope() bool {
	return m.Has(tag.Scope)
}

// HasPriceDelta returns true if PriceDelta is present, Tag 811.
func (m NoMDEntries) HasPriceDelta() bool {
	return m.Has(tag.PriceDelta)
}

// HasNetChgPrevDay returns true if NetChgPrevDay is present, Tag 451.
func (m NoMDEntries) HasNetChgPrevDay() bool {
	return m.Has(tag.NetChgPrevDay)
}

// HasText returns true if Text is present, Tag 58.
func (m NoMDEntries) HasText() bool {
	return m.Has(tag.Text)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m NoMDEntries) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m NoMDEntries) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasOrderCapacity returns true if OrderCapacity is present, Tag 528.
func (m NoMDEntries) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

// HasMDOriginType returns true if MDOriginType is present, Tag 1024.
func (m NoMDEntries) HasMDOriginType() bool {
	return m.Has(tag.MDOriginType)
}

// HasHighPx returns true if HighPx is present, Tag 332.
func (m NoMDEntries) HasHighPx() bool {
	return m.Has(tag.HighPx)
}

// HasLowPx returns true if LowPx is present, Tag 333.
func (m NoMDEntries) HasLowPx() bool {
	return m.Has(tag.LowPx)
}

// HasTradeVolume returns true if TradeVolume is present, Tag 1020.
func (m NoMDEntries) HasTradeVolume() bool {
	return m.Has(tag.TradeVolume)
}

// HasSettlType returns true if SettlType is present, Tag 63.
func (m NoMDEntries) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

// HasSettlDate returns true if SettlDate is present, Tag 64.
func (m NoMDEntries) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

// HasMDQuoteType returns true if MDQuoteType is present, Tag 1070.
func (m NoMDEntries) HasMDQuoteType() bool {
	return m.Has(tag.MDQuoteType)
}

// HasRptSeq returns true if RptSeq is present, Tag 83.
func (m NoMDEntries) HasRptSeq() bool {
	return m.Has(tag.RptSeq)
}

// HasDealingCapacity returns true if DealingCapacity is present, Tag 1048.
func (m NoMDEntries) HasDealingCapacity() bool {
	return m.Has(tag.DealingCapacity)
}

// HasMDEntrySpotRate returns true if MDEntrySpotRate is present, Tag 1026.
func (m NoMDEntries) HasMDEntrySpotRate() bool {
	return m.Has(tag.MDEntrySpotRate)
}

// HasMDEntryForwardPoints returns true if MDEntryForwardPoints is present, Tag 1027.
func (m NoMDEntries) HasMDEntryForwardPoints() bool {
	return m.Has(tag.MDEntryForwardPoints)
}

// HasMDPriceLevel returns true if MDPriceLevel is present, Tag 1023.
func (m NoMDEntries) HasMDPriceLevel() bool {
	return m.Has(tag.MDPriceLevel)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m NoMDEntries) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198.
func (m NoMDEntries) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

// HasOrdType returns true if OrdType is present, Tag 40.
func (m NoMDEntries) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

// HasMDSubBookType returns true if MDSubBookType is present, Tag 1173.
func (m NoMDEntries) HasMDSubBookType() bool {
	return m.Has(tag.MDSubBookType)
}

// HasMarketDepth returns true if MarketDepth is present, Tag 264.
func (m NoMDEntries) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

// HasPriceType returns true if PriceType is present, Tag 423.
func (m NoMDEntries) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

// HasYieldType returns true if YieldType is present, Tag 235.
func (m NoMDEntries) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

// HasYield returns true if Yield is present, Tag 236.
func (m NoMDEntries) HasYield() bool {
	return m.Has(tag.Yield)
}

// HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701.
func (m NoMDEntries) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

// HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696.
func (m NoMDEntries) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

// HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697.
func (m NoMDEntries) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

// HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698.
func (m NoMDEntries) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

// HasSpread returns true if Spread is present, Tag 218.
func (m NoMDEntries) HasSpread() bool {
	return m.Has(tag.Spread)
}

// HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220.
func (m NoMDEntries) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

// HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221.
func (m NoMDEntries) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

// HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222.
func (m NoMDEntries) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

// HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662.
func (m NoMDEntries) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

// HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663.
func (m NoMDEntries) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

// HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699.
func (m NoMDEntries) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

// HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761.
func (m NoMDEntries) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

// HasNoOfSecSizes returns true if NoOfSecSizes is present, Tag 1177.
func (m NoMDEntries) HasNoOfSecSizes() bool {
	return m.Has(tag.NoOfSecSizes)
}

// HasLotType returns true if LotType is present, Tag 1093.
func (m NoMDEntries) HasLotType() bool {
	return m.Has(tag.LotType)
}

// HasSecurityTradingStatus returns true if SecurityTradingStatus is present, Tag 326.
func (m NoMDEntries) HasSecurityTradingStatus() bool {
	return m.Has(tag.SecurityTradingStatus)
}

// HasHaltReasonInt returns true if HaltReasonInt is present, Tag 327.
func (m NoMDEntries) HasHaltReasonInt() bool {
	return m.Has(tag.HaltReasonInt)
}

// HasTrdType returns true if TrdType is present, Tag 828.
func (m NoMDEntries) HasTrdType() bool {
	return m.Has(tag.TrdType)
}

// HasMatchType returns true if MatchType is present, Tag 574.
func (m NoMDEntries) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

// HasTradeID returns true if TradeID is present, Tag 1003.
func (m NoMDEntries) HasTradeID() bool {
	return m.Has(tag.TradeID)
}

// HasTransBkdTime returns true if TransBkdTime is present, Tag 483.
func (m NoMDEntries) HasTransBkdTime() bool {
	return m.Has(tag.TransBkdTime)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m NoMDEntries) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasNoStatsIndicators returns true if NoStatsIndicators is present, Tag 1175.
func (m NoMDEntries) HasNoStatsIndicators() bool {
	return m.Has(tag.NoStatsIndicators)
}

// HasSettlCurrency returns true if SettlCurrency is present, Tag 120.
func (m NoMDEntries) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

// HasNoRateSources returns true if NoRateSources is present, Tag 1445.
func (m NoMDEntries) HasNoRateSources() bool {
	return m.Has(tag.NoRateSources)
}

// HasFirstPx returns true if FirstPx is present, Tag 1025.
func (m NoMDEntries) HasFirstPx() bool {
	return m.Has(tag.FirstPx)
}

// HasLastPx returns true if LastPx is present, Tag 31.
func (m NoMDEntries) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

// HasMDStreamID returns true if MDStreamID is present, Tag 1500.
func (m NoMDEntries) HasMDStreamID() bool {
	return m.Has(tag.MDStreamID)
}

// NoSecurityAltID is a repeating group element, Tag 454.
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455.
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455.
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455.
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456.
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454.
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup.
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

// Add create and append a new NoSecurityAltID to this group.
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup.
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoEvents is a repeating group element, Tag 864.
type NoEvents struct {
	*quickfix.Group
}

// SetEventType sets EventType, Tag 865.
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

// SetEventDate sets EventDate, Tag 866.
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

// SetEventPx sets EventPx, Tag 867.
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

// SetEventText sets EventText, Tag 868.
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

// SetEventTime sets EventTime, Tag 1145.
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

// GetEventType gets EventType, Tag 865.
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventDate gets EventDate, Tag 866.
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventPx gets EventPx, Tag 867.
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventText gets EventText, Tag 868.
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventTime gets EventTime, Tag 1145.
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEventType returns true if EventType is present, Tag 865.
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

// HasEventDate returns true if EventDate is present, Tag 866.
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

// HasEventPx returns true if EventPx is present, Tag 867.
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

// HasEventText returns true if EventText is present, Tag 868.
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

// HasEventTime returns true if EventTime is present, Tag 1145.
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864.
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup.
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

// Add create and append a new NoEvents to this group.
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

// Get returns the ith NoEvents in the NoEventsRepeatinGroup.
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

// NoInstrumentParties is a repeating group element, Tag 1018.
type NoInstrumentParties struct {
	*quickfix.Group
}

// SetInstrumentPartyID sets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

// SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

// SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

// SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetInstrumentPartyID gets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019.
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

// HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050.
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

// HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051.
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

// HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052.
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

// NoInstrumentPartySubIDs is a repeating group element, Tag 1052.
type NoInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

// SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

// GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

// HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

// NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052.
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup.
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

// Add create and append a new NoInstrumentPartySubIDs to this group.
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

// Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup.
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018.
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup.
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoInstrumentParties to this group.
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

// Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup.
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoComplexEvents is a repeating group element, Tag 1483.
type NoComplexEvents struct {
	*quickfix.Group
}

// SetComplexEventType sets ComplexEventType, Tag 1484.
func (m NoComplexEvents) SetComplexEventType(v enum.ComplexEventType) {
	m.Set(field.NewComplexEventType(v))
}

// SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

// SetComplexEventPrice sets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

// SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v enum.ComplexEventPriceBoundaryMethod) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

// SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
}

// SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) SetComplexEventPriceTimeType(v enum.ComplexEventPriceTimeType) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

// SetComplexEventCondition sets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) SetComplexEventCondition(v enum.ComplexEventCondition) {
	m.Set(field.NewComplexEventCondition(v))
}

// SetNoComplexEventDates sets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventType gets ComplexEventType, Tag 1484.
func (m NoComplexEvents) GetComplexEventType() (v enum.ComplexEventType, err quickfix.MessageRejectError) {
	var f field.ComplexEventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) GetComplexOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexOptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPrice gets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) GetComplexEventPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (v enum.ComplexEventPriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) GetComplexEventPriceTimeType() (v enum.ComplexEventPriceTimeType, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceTimeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventCondition gets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) GetComplexEventCondition() (v enum.ComplexEventCondition, err quickfix.MessageRejectError) {
	var f field.ComplexEventConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventDates gets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventType returns true if ComplexEventType is present, Tag 1484.
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

// HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485.
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

// HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486.
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

// HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487.
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

// HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488.
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

// HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489.
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

// HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490.
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

// HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491.
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

// NoComplexEventDates is a repeating group element, Tag 1491.
type NoComplexEventDates struct {
	*quickfix.Group
}

// SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

// SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

// SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) GetComplexEventStartDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) GetComplexEventEndDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492.
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

// HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493.
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

// HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494.
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

// NoComplexEventTimes is a repeating group element, Tag 1494.
type NoComplexEventTimes struct {
	*quickfix.Group
}

// SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

// SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

// GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) GetComplexEventStartTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) GetComplexEventEndTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495.
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

// HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496.
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

// NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494.
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup.
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

// Add create and append a new NoComplexEventTimes to this group.
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

// Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup.
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

// NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491.
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup.
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
}

// Add create and append a new NoComplexEventDates to this group.
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

// Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup.
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

// NoComplexEventsRepeatingGroup is a repeating group, Tag 1483.
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup.
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
}

// Add create and append a new NoComplexEvents to this group.
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

// Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup.
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}

// NoUnderlyings is a repeating group element, Tag 711.
type NoUnderlyings struct {
	*quickfix.Group
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

// SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

// SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

// SetUnderlyingQty sets UnderlyingQty, Tag 879.
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

// SetUnderlyingPx sets UnderlyingPx, Tag 810.
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

// SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

// SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

// SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

// SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

// SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

// SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

// SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

// SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

// SetUnderlyingCashType sets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

// SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

// SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

// SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

// SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

// SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

// SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

// SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

// SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

// SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

// SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

// SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

// SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

// SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

// SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

// SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

// SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

// SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

// SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

// SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingQty gets UnderlyingQty, Tag 879.
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPx gets UnderlyingPx, Tag 810.
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashType gets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) GetUnderlyingContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) GetUnderlyingFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) GetUnderlyingRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) GetUnderlyingSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) GetUnderlyingNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) GetUnderlyingOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingOriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) GetUnderlyingAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) GetUnderlyingDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311.
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312.
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309.
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305.
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457.
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462.
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463.
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310.
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763.
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313.
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542.
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241.
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242.
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243.
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244.
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245.
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246.
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256.
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595.
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592.
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593.
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594.
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247.
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316.
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941.
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317.
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436.
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435.
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308.
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306.
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362.
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363.
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307.
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877.
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

// HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878.
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

// HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318.
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

// HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879.
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

// HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810.
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

// HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882.
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

// HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883.
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

// HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884.
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

// HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885.
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

// HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886.
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

// HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887.
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

// HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972.
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

// HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975.
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

// HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973.
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

// HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974.
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

// HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

// HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000.
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

// HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038.
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

// HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058.
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

// HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039.
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

// HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044.
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

// HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045.
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

// HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046.
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

// HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213.
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

// HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315.
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

// HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419.
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

// HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

// HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

// HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

// HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437.
func (m NoUnderlyings) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

// HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441.
func (m NoUnderlyings) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

// HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453.
func (m NoUnderlyings) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

// HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454.
func (m NoUnderlyings) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

// HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455.
func (m NoUnderlyings) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

// HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456.
func (m NoUnderlyings) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

// HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459.
func (m NoUnderlyings) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

// HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460.
func (m NoUnderlyings) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457.
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457.
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup.
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingStips is a repeating group element, Tag 887.
type NoUnderlyingStips struct {
	*quickfix.Group
}

// SetUnderlyingStipType sets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

// SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

// GetUnderlyingStipType gets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888.
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

// HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889.
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

// NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887.
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup.
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

// Add create and append a new NoUnderlyingStips to this group.
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

// Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup.
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentParties is a repeating group element, Tag 1058.
type NoUndlyInstrumentParties struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartyID sets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyID(v))
}

// SetUnderlyingInstrumentPartyIDSource sets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyIDSource(v))
}

// SetUnderlyingInstrumentPartyRole sets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.Set(field.NewUnderlyingInstrumentPartyRole(v))
}

// SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetUnderlyingInstrumentPartyID gets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUnderlyingInstrumentPartyID returns true if UnderlyingInstrumentPartyID is present, Tag 1059.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyID() bool {
	return m.Has(tag.UnderlyingInstrumentPartyID)
}

// HasUnderlyingInstrumentPartyIDSource returns true if UnderlyingInstrumentPartyIDSource is present, Tag 1060.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyIDSource() bool {
	return m.Has(tag.UnderlyingInstrumentPartyIDSource)
}

// HasUnderlyingInstrumentPartyRole returns true if UnderlyingInstrumentPartyRole is present, Tag 1061.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyRole() bool {
	return m.Has(tag.UnderlyingInstrumentPartyRole)
}

// HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062.
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

// NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062.
type NoUndlyInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartySubID sets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartySubID(v))
}

// SetUnderlyingInstrumentPartySubIDType sets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubIDType(v int) {
	m.Set(field.NewUnderlyingInstrumentPartySubIDType(v))
}

// GetUnderlyingInstrumentPartySubID gets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingInstrumentPartySubID returns true if UnderlyingInstrumentPartySubID is present, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubID() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubID)
}

// HasUnderlyingInstrumentPartySubIDType returns true if UnderlyingInstrumentPartySubIDType is present, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubIDType() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubIDType)
}

// NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062.
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup.
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartySubID), quickfix.GroupElement(tag.UnderlyingInstrumentPartySubIDType)})}
}

// Add create and append a new NoUndlyInstrumentPartySubIDs to this group.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

// Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058.
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup.
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoUndlyInstrumentParties to this group.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

// Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingsRepeatingGroup is a repeating group, Tag 711.
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup.
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint)})}
}

// Add create and append a new NoUnderlyings to this group.
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

// Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup.
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

// NoLegs is a repeating group element, Tag 555.
type NoLegs struct {
	*quickfix.Group
}

// SetLegSymbol sets LegSymbol, Tag 600.
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

// SetLegSymbolSfx sets LegSymbolSfx, Tag 601.
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

// SetLegSecurityID sets LegSecurityID, Tag 602.
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

// SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603.
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

// SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604.
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetLegProduct sets LegProduct, Tag 607.
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

// SetLegCFICode sets LegCFICode, Tag 608.
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

// SetLegSecurityType sets LegSecurityType, Tag 609.
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

// SetLegSecuritySubType sets LegSecuritySubType, Tag 764.
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

// SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610.
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

// SetLegMaturityDate sets LegMaturityDate, Tag 611.
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

// SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248.
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

// SetLegIssueDate sets LegIssueDate, Tag 249.
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

// SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

// SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251.
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

// SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252.
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

// SetLegFactor sets LegFactor, Tag 253.
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

// SetLegCreditRating sets LegCreditRating, Tag 257.
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

// SetLegInstrRegistry sets LegInstrRegistry, Tag 599.
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

// SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596.
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

// SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

// SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598.
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

// SetLegRedemptionDate sets LegRedemptionDate, Tag 254.
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

// SetLegStrikePrice sets LegStrikePrice, Tag 612.
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

// SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942.
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

// SetLegOptAttribute sets LegOptAttribute, Tag 613.
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

// SetLegContractMultiplier sets LegContractMultiplier, Tag 614.
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

// SetLegCouponRate sets LegCouponRate, Tag 615.
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

// SetLegSecurityExchange sets LegSecurityExchange, Tag 616.
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

// SetLegIssuer sets LegIssuer, Tag 617.
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

// SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

// SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619.
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

// SetLegSecurityDesc sets LegSecurityDesc, Tag 620.
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

// SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

// SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

// SetLegRatioQty sets LegRatioQty, Tag 623.
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

// SetLegSide sets LegSide, Tag 624.
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

// SetLegCurrency sets LegCurrency, Tag 556.
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

// SetLegPool sets LegPool, Tag 740.
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

// SetLegDatedDate sets LegDatedDate, Tag 739.
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

// SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955.
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

// SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956.
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

// SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999.
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

// SetLegTimeUnit sets LegTimeUnit, Tag 1001.
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

// SetLegOptionRatio sets LegOptionRatio, Tag 1017.
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

// SetLegPrice sets LegPrice, Tag 566.
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
}

// SetLegMaturityTime sets LegMaturityTime, Tag 1212.
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

// SetLegPutOrCall sets LegPutOrCall, Tag 1358.
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

// SetLegExerciseStyle sets LegExerciseStyle, Tag 1420.
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

// SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

// SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

// SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

// SetLegContractMultiplierUnit sets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) SetLegContractMultiplierUnit(v int) {
	m.Set(field.NewLegContractMultiplierUnit(v))
}

// SetLegFlowScheduleType sets LegFlowScheduleType, Tag 1440.
func (m NoLegs) SetLegFlowScheduleType(v int) {
	m.Set(field.NewLegFlowScheduleType(v))
}

// GetLegSymbol gets LegSymbol, Tag 600.
func (m NoLegs) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSymbolSfx gets LegSymbolSfx, Tag 601.
func (m NoLegs) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityID gets LegSecurityID, Tag 602.
func (m NoLegs) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603.
func (m NoLegs) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604.
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLegProduct gets LegProduct, Tag 607.
func (m NoLegs) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCFICode gets LegCFICode, Tag 608.
func (m NoLegs) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityType gets LegSecurityType, Tag 609.
func (m NoLegs) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecuritySubType gets LegSecuritySubType, Tag 764.
func (m NoLegs) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610.
func (m NoLegs) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityDate gets LegMaturityDate, Tag 611.
func (m NoLegs) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248.
func (m NoLegs) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssueDate gets LegIssueDate, Tag 249.
func (m NoLegs) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251.
func (m NoLegs) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252.
func (m NoLegs) GetLegRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFactor gets LegFactor, Tag 253.
func (m NoLegs) GetLegFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCreditRating gets LegCreditRating, Tag 257.
func (m NoLegs) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInstrRegistry gets LegInstrRegistry, Tag 599.
func (m NoLegs) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596.
func (m NoLegs) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598.
func (m NoLegs) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRedemptionDate gets LegRedemptionDate, Tag 254.
func (m NoLegs) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikePrice gets LegStrikePrice, Tag 612.
func (m NoLegs) GetLegStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942.
func (m NoLegs) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptAttribute gets LegOptAttribute, Tag 613.
func (m NoLegs) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplier gets LegContractMultiplier, Tag 614.
func (m NoLegs) GetLegContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponRate gets LegCouponRate, Tag 615.
func (m NoLegs) GetLegCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityExchange gets LegSecurityExchange, Tag 616.
func (m NoLegs) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssuer gets LegIssuer, Tag 617.
func (m NoLegs) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619.
func (m NoLegs) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityDesc gets LegSecurityDesc, Tag 620.
func (m NoLegs) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRatioQty gets LegRatioQty, Tag 623.
func (m NoLegs) GetLegRatioQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSide gets LegSide, Tag 624.
func (m NoLegs) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCurrency gets LegCurrency, Tag 556.
func (m NoLegs) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPool gets LegPool, Tag 740.
func (m NoLegs) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegDatedDate gets LegDatedDate, Tag 739.
func (m NoLegs) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955.
func (m NoLegs) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956.
func (m NoLegs) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999.
func (m NoLegs) GetLegUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegTimeUnit gets LegTimeUnit, Tag 1001.
func (m NoLegs) GetLegTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.LegTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptionRatio gets LegOptionRatio, Tag 1017.
func (m NoLegs) GetLegOptionRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegOptionRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPrice gets LegPrice, Tag 566.
func (m NoLegs) GetLegPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityTime gets LegMaturityTime, Tag 1212.
func (m NoLegs) GetLegMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPutOrCall gets LegPutOrCall, Tag 1358.
func (m NoLegs) GetLegPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.LegPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegExerciseStyle gets LegExerciseStyle, Tag 1420.
func (m NoLegs) GetLegExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.LegExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) GetLegUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) GetLegPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplierUnit gets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) GetLegContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFlowScheduleType gets LegFlowScheduleType, Tag 1440.
func (m NoLegs) GetLegFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.LegFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSymbol returns true if LegSymbol is present, Tag 600.
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

// HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601.
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

// HasLegSecurityID returns true if LegSecurityID is present, Tag 602.
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

// HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603.
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

// HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604.
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

// HasLegProduct returns true if LegProduct is present, Tag 607.
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

// HasLegCFICode returns true if LegCFICode is present, Tag 608.
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

// HasLegSecurityType returns true if LegSecurityType is present, Tag 609.
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

// HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764.
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

// HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610.
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

// HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611.
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

// HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248.
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

// HasLegIssueDate returns true if LegIssueDate is present, Tag 249.
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

// HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250.
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

// HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251.
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

// HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252.
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

// HasLegFactor returns true if LegFactor is present, Tag 253.
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

// HasLegCreditRating returns true if LegCreditRating is present, Tag 257.
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

// HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599.
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

// HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596.
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

// HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597.
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

// HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598.
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

// HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254.
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

// HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612.
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

// HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942.
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

// HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613.
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

// HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614.
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

// HasLegCouponRate returns true if LegCouponRate is present, Tag 615.
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

// HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616.
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

// HasLegIssuer returns true if LegIssuer is present, Tag 617.
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

// HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618.
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

// HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619.
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

// HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620.
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

// HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621.
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

// HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622.
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

// HasLegRatioQty returns true if LegRatioQty is present, Tag 623.
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

// HasLegSide returns true if LegSide is present, Tag 624.
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

// HasLegCurrency returns true if LegCurrency is present, Tag 556.
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

// HasLegPool returns true if LegPool is present, Tag 740.
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

// HasLegDatedDate returns true if LegDatedDate is present, Tag 739.
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

// HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955.
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

// HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956.
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

// HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999.
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

// HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001.
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

// HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017.
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

// HasLegPrice returns true if LegPrice is present, Tag 566.
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

// HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212.
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

// HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358.
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

// HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420.
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

// HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224.
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

// HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421.
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

// HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422.
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

// HasLegContractMultiplierUnit returns true if LegContractMultiplierUnit is present, Tag 1436.
func (m NoLegs) HasLegContractMultiplierUnit() bool {
	return m.Has(tag.LegContractMultiplierUnit)
}

// HasLegFlowScheduleType returns true if LegFlowScheduleType is present, Tag 1440.
func (m NoLegs) HasLegFlowScheduleType() bool {
	return m.Has(tag.LegFlowScheduleType)
}

// NoLegSecurityAltID is a repeating group element, Tag 604.
type NoLegSecurityAltID struct {
	*quickfix.Group
}

// SetLegSecurityAltID sets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

// SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

// GetLegSecurityAltID gets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605.
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

// HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606.
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

// NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604.
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup.
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

// Add create and append a new NoLegSecurityAltID to this group.
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

// Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup.
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoLegsRepeatingGroup is a repeating group, Tag 555.
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup.
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegContractMultiplierUnit), quickfix.GroupElement(tag.LegFlowScheduleType)})}
}

// Add create and append a new NoLegs to this group.
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

// Get returns the ith NoLegs in the NoLegsRepeatinGroup.
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoOfSecSizes is a repeating group element, Tag 1177.
type NoOfSecSizes struct {
	*quickfix.Group
}

// SetMDSecSizeType sets MDSecSizeType, Tag 1178.
func (m NoOfSecSizes) SetMDSecSizeType(v enum.MDSecSizeType) {
	m.Set(field.NewMDSecSizeType(v))
}

// SetMDSecSize sets MDSecSize, Tag 1179.
func (m NoOfSecSizes) SetMDSecSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDSecSize(value, scale))
}

// GetMDSecSizeType gets MDSecSizeType, Tag 1178.
func (m NoOfSecSizes) GetMDSecSizeType() (v enum.MDSecSizeType, err quickfix.MessageRejectError) {
	var f field.MDSecSizeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDSecSize gets MDSecSize, Tag 1179.
func (m NoOfSecSizes) GetMDSecSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDSecSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMDSecSizeType returns true if MDSecSizeType is present, Tag 1178.
func (m NoOfSecSizes) HasMDSecSizeType() bool {
	return m.Has(tag.MDSecSizeType)
}

// HasMDSecSize returns true if MDSecSize is present, Tag 1179.
func (m NoOfSecSizes) HasMDSecSize() bool {
	return m.Has(tag.MDSecSize)
}

// NoOfSecSizesRepeatingGroup is a repeating group, Tag 1177.
type NoOfSecSizesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoOfSecSizesRepeatingGroup returns an initialized, NoOfSecSizesRepeatingGroup.
func NewNoOfSecSizesRepeatingGroup() NoOfSecSizesRepeatingGroup {
	return NoOfSecSizesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOfSecSizes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDSecSizeType), quickfix.GroupElement(tag.MDSecSize)})}
}

// Add create and append a new NoOfSecSizes to this group.
func (m NoOfSecSizesRepeatingGroup) Add() NoOfSecSizes {
	g := m.RepeatingGroup.Add()
	return NoOfSecSizes{g}
}

// Get returns the ith NoOfSecSizes in the NoOfSecSizesRepeatinGroup.
func (m NoOfSecSizesRepeatingGroup) Get(i int) NoOfSecSizes {
	return NoOfSecSizes{m.RepeatingGroup.Get(i)}
}

// NoStatsIndicators is a repeating group element, Tag 1175.
type NoStatsIndicators struct {
	*quickfix.Group
}

// SetStatsType sets StatsType, Tag 1176.
func (m NoStatsIndicators) SetStatsType(v enum.StatsType) {
	m.Set(field.NewStatsType(v))
}

// GetStatsType gets StatsType, Tag 1176.
func (m NoStatsIndicators) GetStatsType() (v enum.StatsType, err quickfix.MessageRejectError) {
	var f field.StatsTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasStatsType returns true if StatsType is present, Tag 1176.
func (m NoStatsIndicators) HasStatsType() bool {
	return m.Has(tag.StatsType)
}

// NoStatsIndicatorsRepeatingGroup is a repeating group, Tag 1175.
type NoStatsIndicatorsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoStatsIndicatorsRepeatingGroup returns an initialized, NoStatsIndicatorsRepeatingGroup.
func NewNoStatsIndicatorsRepeatingGroup() NoStatsIndicatorsRepeatingGroup {
	return NoStatsIndicatorsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStatsIndicators,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StatsType)})}
}

// Add create and append a new NoStatsIndicators to this group.
func (m NoStatsIndicatorsRepeatingGroup) Add() NoStatsIndicators {
	g := m.RepeatingGroup.Add()
	return NoStatsIndicators{g}
}

// Get returns the ith NoStatsIndicators in the NoStatsIndicatorsRepeatinGroup.
func (m NoStatsIndicatorsRepeatingGroup) Get(i int) NoStatsIndicators {
	return NoStatsIndicators{m.RepeatingGroup.Get(i)}
}

// NoRateSources is a repeating group element, Tag 1445.
type NoRateSources struct {
	*quickfix.Group
}

// SetRateSource sets RateSource, Tag 1446.
func (m NoRateSources) SetRateSource(v enum.RateSource) {
	m.Set(field.NewRateSource(v))
}

// SetRateSourceType sets RateSourceType, Tag 1447.
func (m NoRateSources) SetRateSourceType(v enum.RateSourceType) {
	m.Set(field.NewRateSourceType(v))
}

// SetReferencePage sets ReferencePage, Tag 1448.
func (m NoRateSources) SetReferencePage(v string) {
	m.Set(field.NewReferencePage(v))
}

// GetRateSource gets RateSource, Tag 1446.
func (m NoRateSources) GetRateSource() (v enum.RateSource, err quickfix.MessageRejectError) {
	var f field.RateSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRateSourceType gets RateSourceType, Tag 1447.
func (m NoRateSources) GetRateSourceType() (v enum.RateSourceType, err quickfix.MessageRejectError) {
	var f field.RateSourceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetReferencePage gets ReferencePage, Tag 1448.
func (m NoRateSources) GetReferencePage() (v string, err quickfix.MessageRejectError) {
	var f field.ReferencePageField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRateSource returns true if RateSource is present, Tag 1446.
func (m NoRateSources) HasRateSource() bool {
	return m.Has(tag.RateSource)
}

// HasRateSourceType returns true if RateSourceType is present, Tag 1447.
func (m NoRateSources) HasRateSourceType() bool {
	return m.Has(tag.RateSourceType)
}

// HasReferencePage returns true if ReferencePage is present, Tag 1448.
func (m NoRateSources) HasReferencePage() bool {
	return m.Has(tag.ReferencePage)
}

// NoRateSourcesRepeatingGroup is a repeating group, Tag 1445.
type NoRateSourcesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRateSourcesRepeatingGroup returns an initialized, NoRateSourcesRepeatingGroup.
func NewNoRateSourcesRepeatingGroup() NoRateSourcesRepeatingGroup {
	return NoRateSourcesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRateSources,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RateSource), quickfix.GroupElement(tag.RateSourceType), quickfix.GroupElement(tag.ReferencePage)})}
}

// Add create and append a new NoRateSources to this group.
func (m NoRateSourcesRepeatingGroup) Add() NoRateSources {
	g := m.RepeatingGroup.Add()
	return NoRateSources{g}
}

// Get returns the ith NoRateSources in the NoRateSourcesRepeatinGroup.
func (m NoRateSourcesRepeatingGroup) Get(i int) NoRateSources {
	return NoRateSources{m.RepeatingGroup.Get(i)}
}

// NoMDEntriesRepeatingGroup is a repeating group, Tag 268.
type NoMDEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMDEntriesRepeatingGroup returns an initialized, NoMDEntriesRepeatingGroup.
func NewNoMDEntriesRepeatingGroup() NoMDEntriesRepeatingGroup {
	return NoMDEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDEntries,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDUpdateAction), quickfix.GroupElement(tag.DeleteReason), quickfix.GroupElement(tag.MDEntryType), quickfix.GroupElement(tag.MDEntryID), quickfix.GroupElement(tag.MDEntryRefID), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), NewNoEventsRepeatingGroup(), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.SecurityStatus), quickfix.GroupElement(tag.SettleOnOpenFlag), quickfix.GroupElement(tag.InstrmtAssignmentMethod), quickfix.GroupElement(tag.StrikeMultiplier), quickfix.GroupElement(tag.StrikeValue), quickfix.GroupElement(tag.MinPriceIncrement), quickfix.GroupElement(tag.PositionLimit), quickfix.GroupElement(tag.NTPositionLimit), NewNoInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnitOfMeasure), quickfix.GroupElement(tag.TimeUnit), quickfix.GroupElement(tag.MaturityTime), quickfix.GroupElement(tag.SecurityGroup), quickfix.GroupElement(tag.MinPriceIncrementAmount), quickfix.GroupElement(tag.UnitOfMeasureQty), quickfix.GroupElement(tag.SecurityXMLLen), quickfix.GroupElement(tag.SecurityXML), quickfix.GroupElement(tag.SecurityXMLSchema), quickfix.GroupElement(tag.ProductComplex), quickfix.GroupElement(tag.PriceUnitOfMeasure), quickfix.GroupElement(tag.PriceUnitOfMeasureQty), quickfix.GroupElement(tag.SettlMethod), quickfix.GroupElement(tag.ExerciseStyle), quickfix.GroupElement(tag.OptPayoutAmount), quickfix.GroupElement(tag.PriceQuoteMethod), quickfix.GroupElement(tag.ListMethod), quickfix.GroupElement(tag.CapPrice), quickfix.GroupElement(tag.FloorPrice), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.FlexibleIndicator), quickfix.GroupElement(tag.FlexProductEligibilityIndicator), quickfix.GroupElement(tag.ValuationMethod), quickfix.GroupElement(tag.ContractMultiplierUnit), quickfix.GroupElement(tag.FlowScheduleType), quickfix.GroupElement(tag.RestructuringType), quickfix.GroupElement(tag.Seniority), quickfix.GroupElement(tag.NotionalPercentageOutstanding), quickfix.GroupElement(tag.OriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.AttachmentPoint), quickfix.GroupElement(tag.DetachmentPoint), quickfix.GroupElement(tag.StrikePriceDeterminationMethod), quickfix.GroupElement(tag.StrikePriceBoundaryMethod), quickfix.GroupElement(tag.StrikePriceBoundaryPrecision), quickfix.GroupElement(tag.UnderlyingPriceDeterminationMethod), quickfix.GroupElement(tag.OptPayoutType), NewNoComplexEventsRepeatingGroup(), NewNoUnderlyingsRepeatingGroup(), NewNoLegsRepeatingGroup(), quickfix.GroupElement(tag.FinancialStatus), quickfix.GroupElement(tag.CorporateAction), quickfix.GroupElement(tag.MDEntryPx), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.MDEntrySize), quickfix.GroupElement(tag.MDEntryDate), quickfix.GroupElement(tag.MDEntryTime), quickfix.GroupElement(tag.TickDirection), quickfix.GroupElement(tag.MDMkt), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.QuoteCondition), quickfix.GroupElement(tag.TradeCondition), quickfix.GroupElement(tag.MDEntryOriginator), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID), quickfix.GroupElement(tag.OpenCloseSettlFlag), quickfix.GroupElement(tag.TimeInForce), quickfix.GroupElement(tag.ExpireDate), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.MinQty), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.SellerDays), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.QuoteEntryID), quickfix.GroupElement(tag.MDEntryBuyer), quickfix.GroupElement(tag.MDEntrySeller), quickfix.GroupElement(tag.NumberOfOrders), quickfix.GroupElement(tag.MDEntryPositionNo), quickfix.GroupElement(tag.Scope), quickfix.GroupElement(tag.PriceDelta), quickfix.GroupElement(tag.NetChgPrevDay), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.MDOriginType), quickfix.GroupElement(tag.HighPx), quickfix.GroupElement(tag.LowPx), quickfix.GroupElement(tag.TradeVolume), quickfix.GroupElement(tag.SettlType), quickfix.GroupElement(tag.SettlDate), quickfix.GroupElement(tag.MDQuoteType), quickfix.GroupElement(tag.RptSeq), quickfix.GroupElement(tag.DealingCapacity), quickfix.GroupElement(tag.MDEntrySpotRate), quickfix.GroupElement(tag.MDEntryForwardPoints), quickfix.GroupElement(tag.MDPriceLevel), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.MDSubBookType), quickfix.GroupElement(tag.MarketDepth), quickfix.GroupElement(tag.PriceType), quickfix.GroupElement(tag.YieldType), quickfix.GroupElement(tag.Yield), quickfix.GroupElement(tag.YieldCalcDate), quickfix.GroupElement(tag.YieldRedemptionDate), quickfix.GroupElement(tag.YieldRedemptionPrice), quickfix.GroupElement(tag.YieldRedemptionPriceType), quickfix.GroupElement(tag.Spread), quickfix.GroupElement(tag.BenchmarkCurveCurrency), quickfix.GroupElement(tag.BenchmarkCurveName), quickfix.GroupElement(tag.BenchmarkCurvePoint), quickfix.GroupElement(tag.BenchmarkPrice), quickfix.GroupElement(tag.BenchmarkPriceType), quickfix.GroupElement(tag.BenchmarkSecurityID), quickfix.GroupElement(tag.BenchmarkSecurityIDSource), NewNoOfSecSizesRepeatingGroup(), quickfix.GroupElement(tag.LotType), quickfix.GroupElement(tag.SecurityTradingStatus), quickfix.GroupElement(tag.HaltReasonInt), quickfix.GroupElement(tag.TrdType), quickfix.GroupElement(tag.MatchType), quickfix.GroupElement(tag.TradeID), quickfix.GroupElement(tag.TransBkdTime), quickfix.GroupElement(tag.TransactTime), NewNoStatsIndicatorsRepeatingGroup(), quickfix.GroupElement(tag.SettlCurrency), NewNoRateSourcesRepeatingGroup(), quickfix.GroupElement(tag.FirstPx), quickfix.GroupElement(tag.LastPx), quickfix.GroupElement(tag.MDStreamID)})}
}

// Add create and append a new NoMDEntries to this group.
func (m NoMDEntriesRepeatingGroup) Add() NoMDEntries {
	g := m.RepeatingGroup.Add()
	return NoMDEntries{g}
}

// Get returns the ith NoMDEntries in the NoMDEntriesRepeatinGroup.
func (m NoMDEntriesRepeatingGroup) Get(i int) NoMDEntries {
	return NoMDEntries{m.RepeatingGroup.Get(i)}
}
