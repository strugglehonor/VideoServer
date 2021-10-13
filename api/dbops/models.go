package dbops

import (
	"time"
)

// user uuid和id要分开
type User struct {
	ID       int    `json:"id"`
	UserUUID string `json:"user_uuid" gorm:"unique"`
	Username string `json:"username" gorm:"not null; size:32; unique"`
	Password string `json:"password" gorm:"not null; size:32"`
}

// video
type Video struct {
	ID        int       `json:"id"`
	VideoUUID string    `json:"video_uuid" gorm:"unique"`
	UserUUID  string    `json:"user_uuid"`
	VideoName string    `json:"video_name" gorm:"not null"`
	CreatedAt time.Time `json:"create_time"`
	UpdatedAt time.Time `json:"updated_at"`
}

// comment
type Comment struct {
	ID          int       `json:"id"`
	CommentUUID string    `json:"comment_uuid" gorm:"unique"`
	VideoUUID   string    `json:"video_uuid"`
	UserUUID    string    `json:"user_uuid"`
	Content     string    `json:"content" gorm:"not null"`
	CreatedAt   time.Time `json:"create_time"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// sessions
type Session struct {
	ID          int       `json:"id"`
	SessionUUID string    `json:"session_uuid" gorm:"unique"`
	UserName    string    `json:"username" gorm:"not null; unique"`
	CreatedAt   time.Time `json:"create_time"`
	ExpiredAt   int64     `json:"expired_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
