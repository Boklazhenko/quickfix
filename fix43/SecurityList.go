package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityList struct {
	quickfixgo.Message
}

func (m *SecurityList) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) TotalNumSecurities() (*field.TotalNumSecurities, error) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}