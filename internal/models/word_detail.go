package models

import "time"

func (WordDetail) TableName() string {
	return "words"
}

type WordDetail struct {
	ID             uint      `json:"id"`
	Img            string    `json:"img"`
	Word           string    `json:"word"`
	Accent         string    `json:"accent"`
	MeanCn         string    `json:"mean"`
	Sound          string    `json:"sound"`
	CreatedAt      time.Time `json:"created_at"`
	DeformationImg string    `json:"deformation_img"`
	Sentence       string    `json:"sentence"`
	SentenceTrans  string    `json:"sentence_trans"`
	SentenceAudio  string    `json:"sentence_audio"`
	WordEtymas     string    `json:"word_etyma"`
}
