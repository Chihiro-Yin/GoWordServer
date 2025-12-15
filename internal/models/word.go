package models

import "time"

type Word struct {
	ID        uint      `json:"id"`
	Img       string    `json:"img"`
	Word      string    `json:"word"`
	Accent    string    `json:"accent"`
	MeanCn    string    `json:"mean"`
	Sound     string    `json:"sound"`
	CreatedAt time.Time `json:"created_at"`
}
