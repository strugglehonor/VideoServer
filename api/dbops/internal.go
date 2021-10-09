package dbops

import (
	"fmt"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

// var sessionMap session.sessionMap

// insert session into db
func InsertSession(username string, sessionUUID string, expireTime int64) error{
	err := db.Create(&Session{
		SessionUUID: sessionUUID,
		UserName: username,
		ExpiredAt: expireTime,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete session
func DeleteSession(username string) error {
	var session Session
	err := db.First(session, "username = ?", username).Error
	if err != nil {
		return err
	}

	err = db.Delete(&session).Error
	if err != nil {
		return err
	}

	return nil
}

// retrieve all session
func RetrieveAllSessions() (sync.Map, error) {
	sessions := []Session{}
	sessionMap := sync.Map{}
	res := db.Find(&sessions)
	if res.Error != nil {
		return sessionMap, res.Error
	}
	if res.RowsAffected == 0 {
		return sessionMap, fmt.Errorf("session not found in DB")
	}

	for _, ss := range sessions {
		sessionMap.Store(ss.UserName, ss)
	}
	return sessionMap, nil
}

// retrieve session by sid
func RetrieveSession(sid string) (Session, error) {
	var session Session
	err := db.First(&session, "session_uuid = ?", sid).Error
	if err != nil {
		return Session{}, err
	}
	return session, nil
}