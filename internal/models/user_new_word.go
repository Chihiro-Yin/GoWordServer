package models

import "time"

type UserNewWord struct {
	ID         uint      `json:"id"`
	Nick       string    `json:"nick"`
	WordID     uint      `json:"word_id"`
	IsMastered bool      `json:"is_mastered"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"-"` // 软删除字段，不返回给前端
}
