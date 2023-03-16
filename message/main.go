package message

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
