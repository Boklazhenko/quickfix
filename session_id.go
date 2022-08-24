package quickfix

import (
	"bytes"
	"fmt"
)

// SessionID is a unique identifer of a Session
type SessionID struct {
	BeginString, TargetCompID, TargetSubID, OnBehalfOfCompID, TargetLocationID, SenderCompID, SenderSubID, SenderLocationID, Qualifier string
	DropCopy                                                                                                                           bool
}

//IsFIXT returns true if the SessionID has a FIXT BeginString
func (s SessionID) IsFIXT() bool {
	return s.BeginString == BeginStringFIXT11
}

func appendOptional(b *bytes.Buffer, delim, v string) {
	if len(v) == 0 {
		return
	}

	b.WriteString(delim)
	b.WriteString(v)
}

func (s SessionID) String() string {
	b := new(bytes.Buffer)
	b.WriteString(s.BeginString)
	b.WriteString(":")
	b.WriteString(fmt.Sprintf("[drop_copy:%v]", s.DropCopy))
	b.WriteString(":")
	b.WriteString(s.SenderCompID)

	appendOptional(b, "/", s.SenderSubID)
	appendOptional(b, "/", s.SenderLocationID)

	b.WriteString("->")
	b.WriteString(s.TargetCompID)

	appendOptional(b, "/", s.TargetSubID)
	appendOptional(b, "/", s.TargetLocationID)
	appendOptional(b, "/", s.OnBehalfOfCompID)

	appendOptional(b, ":", s.Qualifier)
	return b.String()
}
