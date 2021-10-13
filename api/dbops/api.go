package dbops

import (
	// "database/sql"  go原生sql封装
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/video_server/api/utils"
)

var (
	video   Video
	user    User
	comment Comment
)

// Add User
func AddUserCredential(username string, password string) error {
	useruuid := utils.NewUUID()
	err := db.Create(&User{Username: username, Password: password, UserUUID: useruuid}).Error
	// 待补充：err要再封装一层
	if err != nil {
		return err
	}
	return nil
}

// Get Password By UserName
func GetUserCredential(username string) (string, error) {
	// recs := db.First(&user, "username = ?", username)
	err := db.First(user, "username = ?", username).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

// Delete User
func DeleteUserCredential(username string) error {
	var user User
	err := db.First(user, "username = ?", username).Error
	if err != nil {
		return err
	}

	err = db.Delete(&user).Error
	if err != nil {
		return err
	}
	
	return nil
}

// Add New Video to DB
func AddNewVideo(userName string, videoName string) (*Video, error) {
	err := db.First(user, "username = ?", userName).Error
	if err != nil {
		return nil, err
	}

	vid := utils.NewUUID()
	err = db.Create(&Video{
		VideoUUID: vid,
		UserUUID: user.UserUUID,
		VideoName: videoName,
	}).Error
	if err != nil {
		return nil, err
	}

	return &Video{VideoUUID: vid, UserUUID: user.UserUUID, VideoName: videoName}, nil
}

// Get Video by video id
func GetVideo(vid string) (*Video, error) {
	err := db.Take(video, "video_uuid = ?", vid).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

// Get Video list by username, use paginate model
func GetVideoPaginate(username string, page, limit int) ([]Video, error) {
	offset := (page - 1) * limit
	recs := []Video{}
	query := db.Where("1=1")
	if username != "" {
		query = db.Where("username = ?", username)
	}

	err := query.Debug().Order("updated_at").Limit(limit).Offset(offset).Find(&recs).Error
	if err != nil {
		return nil, fmt.Errorf("get video paginate from db failed, %w", err)
	}
	return recs, nil
}

// Create Comment
func NewComment(videoUUID, userUUID, content string) error {
	commentUUID := utils.NewUUID()
	err := db.Create(&Comment{
		CommentUUID: commentUUID,
		VideoUUID: videoUUID,
		UserUUID: userUUID,
		Content: content,
	}).Error
	if err != nil {
		return  err
	}
	return nil
}

// List Comment
func ListCommentPaginate(videoname, username string, page, limit int) ([]Comment, error) {
	err := db.First(user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	err = db.First(video, "video_name = ?", videoname).Error
	if err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	recs := []Comment{}
	query := db.Where("1=1")

	if username != "" {
		query = query.Where("user_uuid = ?", user.UserUUID)
	}
	if videoname != "" {
		query = query.Where("video_uuid = ?", video.VideoUUID)
	}
	err = query.Debug().Order("updated_at").Limit(limit).Offset(offset).Find(&recs).Error
	if err != nil {
		return nil, err
	}

	return recs, nil
}