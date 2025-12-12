package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Nick      string    `json:"nick"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type RegisterRequest struct {
	Nick     string `json:"nick" validate:"required,min=2,max=50"` // 昵称：必填，2-50字符
	Email    string `json:"email" validate:"email"`                // 邮箱：选填
	Password string `json:"password" validate:"required,min=6"`    // 密码：必填，至少6位
}
type LoginRequest struct {
	Nick     string `json:"nick" validate:"required,min=2,max=50"` // 昵称：必填，2-50字符
	Password string `json:"password" validate:"required"`          // 密码：必填
}
