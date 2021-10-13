package dbops

import (
	"testing"
)

var tmpvid string

func TestUserMain(m *testing.M) {
	m.Run()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("AddUser", testAddUserCredential)
	t.Run("GetUser", testGetUserCredential)
	t.Run("DeleteUser", testDeleteUserCredential)
}

func TestVideoWorkFlow(t *testing.T) {
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("GetVideoPaginate", testGetVideoPaginate)
}

// func TestCommentWorkFlow

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("ableSu", "123456789")
	if err != nil {
		t.Errorf("AddUserCredential Error: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	passwd, err := GetUserCredential("ableSu")
	if err != nil || passwd != "123456789" {
		t.Errorf("GetUserCredential Error: %v", err)
	}
}

func testDeleteUserCredential(t *testing.T) {
	err := DeleteUserCredential("ableSu")
	if err != nil {
		t.Errorf("DeleteUserCredential Error:%v", err)
	}
}

func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo("ableSu", "video1")
	if err != nil {
		t.Errorf("AddNewVideo Error: %v", err)
	}
	tmpvid = video.VideoUUID
}

func testGetVideo(t *testing.T) {
	_, err := GetVideo(tmpvid)
	if err != nil {
		t.Errorf("GetVideo Error: %v", err)
	}
}

func testGetVideoPaginate(t *testing.T) {
	page, limit := 3, 5
	_, err := GetVideoPaginate("ableSu", page, limit)
	if err != nil {
		t.Errorf("GetVideoPaginate Error: %v", err)
	}
}
