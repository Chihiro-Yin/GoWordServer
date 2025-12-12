package models

import "time"

type Word struct {
	ID        int       `json:"id"`
	Img       string    `json:"img"`
	Word      string    `json:"word"`
	Phonetic  string    `json:"phonetic"`
	Mean      string    `json:"mean"`
	Sound     string    `json:"sound"`
	CreatedAt time.Time `json:"created_at"`
}
