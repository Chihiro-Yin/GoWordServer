package models

import (
	"time"

	"gorm.io/gorm"
)

type UserNewWord struct {
	ID         uint           `json:"id"`
	Nick       string         `json:"nick"`
	WordID     uint           `json:"word_id"`
	IsMastered bool           `json:"is_mastered"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"` // 软删除字段，不返回给前端
}
type UserNewWordVO struct {
	ID         uint      `json:"id"`
	WordID     uint      `json:"word_id"`
	Nick       string    `json:"nick"`
	IsMastered bool      `json:"is_mastered"`
	Img        string    `json:"img"`
	Word       string    `json:"word"`
	Phonetic   string    `json:"phonetic"`
	Mean       string    `json:"mean"`
	Sound      string    `json:"sound"`
	CreatedAt  time.Time `json:"created_at"`
}
type AddNewWordRequest struct {
	WordID uint `json:"word_id"`
}
