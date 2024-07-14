package message

import (
	"encoding/json"
	"io"
)

//
//

// Session struct representing a session.
type Message struct {
	Message  string `json:"message"`  // message
	Source   string `json:"source"`   // source of the message
	Priority int    `json:"priority"` // priority of the message
	Activity int64  `json:"active"`   // int time (unix time) of last time activated
	Created  int64  `json:"created"`  // int time (unix time) of creation time
}

type Messages []Message

func (M Messages) Read(handle io.Reader) error {
	return json.NewDecoder(handle).Decode(&M)
}

func (M Messages) Write(handle io.Writer) error {
	json.Marshal(M)
	return nil
}

// NewMessage returns a new message.
func NewMessage() *Message {
	return &Message{}
}