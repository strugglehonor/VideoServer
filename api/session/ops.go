package session

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/video_server/api/dbops"
	"github.com/video_server/api/defs"
	"github.com/video_server/api/utils"
)

// key: username value: session
var sessionMap sync.Map

// create session
func NewSessionID(username string) (string, error) {
	id := utils.NewUUID()
	ctime := time.Now().UnixNano() / 1e6 // ms
	ttl := ctime + 30*60*1000            // set ttl 30min
	session := defs.Session{
		UserName:   username,
		ExpireTime: ttl,
		CreatedAt:  ctime,
	}
	sessionMap.Store(username, session)

	err := dbops.InsertSession(username, id, ttl)
	if err != nil {
		return "", err
	}

	return id, nil
}

// load session from DB
func LoadSession() {
	ssMap, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ssMap.Range(func(key, value interface{}) bool {
		sessionMap.Store(key, value)
		return true
	})
}

func DeleteSession(username string) error {
	err := dbops.DeleteSession(username)
	if err != nil {
		return err
	}

	sessionMap.Delete(username)
	return nil
}

// judge session is expired or not
func IsSessionExpired(username string) (bool, error) {
	o, ok := sessionMap.Load(username)
	if !ok {
		return false, fmt.Errorf("%s's session not exist", username)
	}

	session, ok := o.(defs.Session)
	if !ok {
		return false, fmt.Errorf("interface transfer failed")
	}

	// now time
	ntime := time.Now().UnixNano() / 1e9
	if ntime-session.CreatedAt > session.ExpireTime {
		// delete session
		DeleteSession(username)
		return true, nil
	}
	return false, nil
}
