package model

type Conversation struct {
	ID       string    `json:"id"`
	Messages []Message `json:"messages"`
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
