package model

type Message struct {
	From    string
	Time    string
	Message string
}

type Chat []*Message

type OutGoing struct {
	Type int    `json:"type"`
	Me   string `json:"me"`
	Msg  string `json:"msg"`
}

type InComing struct {
	Type   int    `json:"type"`
	Active int    `json:"active"`
	From   string `json:"from"`
	Me     string `json:"me"`
	Msg    string `json:"msg"`
}
