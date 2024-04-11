// Code generated from Pkl module `zerodoctor.shawarma.pkg.config`. DO NOT EDIT.
package action

import (
	"encoding"
	"fmt"
)

type Action string

const (
	Continue Action = "continue"
	Pause    Action = "pause"
	Stop     Action = "stop"
)

// String returns the string representation of Action
func (rcv Action) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Action)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Action.
func (rcv *Action) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "continue":
		*rcv = Continue
	case "pause":
		*rcv = Pause
	case "stop":
		*rcv = Stop
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Action`, str)
	}
	return nil
}
