package model

type Statistics struct {
	ID        string `json:"id"`
	BlogID    string `json:"blogId"`
	StartedAt string `json:"startedAt"`
	EndedAt   string `json:"endedAt"`
	Uuid      string `json:"userId"`
}
