package dbops

import (
	"github.com/video_server/api/dbops"
)

var (
	videoDeleteRecord  Video_Delete_Record
	videoDeleteRecords []Video_Delete_Record
	video              dbops.Video
)

// add
func AddVideoDeleteRecord(videoName string) error {
	err := db.First(&video, "video_name = ?", videoName).Error
	if err != nil {
		return err
	}

	err = db.Create(&Video_Delete_Record{VideoUUID: video.VideoUUID, VideoName: videoName}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete
func DeleteVideoDeleteRecord(videoName string) error {
	err := db.First(&videoDeleteRecord, "video_name = ?", videoName).Error
	if err != nil {
		return err
	}

	err = db.Delete(&videoDeleteRecord).Error
	if err != nil {
		return err
	}

	return nil
}

// get
func GetVideoDeleteRecord(num int) ([]Video_Delete_Record, error) {
	err := db.Limit(num).Find(&videoDeleteRecords).Error
	if err != nil {
		return nil, err
	}
	return videoDeleteRecords, nil
}
