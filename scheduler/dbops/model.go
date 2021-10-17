package dbops

type Video_Delete_Record struct {
	ID        int    `json:"id"`
	VideoUUID string `json:"video_uuid" gorm:"unique"`
	VideoName string `json:"video_name" gorm:"not null"`
}
