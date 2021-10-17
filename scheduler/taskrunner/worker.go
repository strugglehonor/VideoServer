package taskrunner

import (
	"errors"
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/video_server/scheduler/dbops"
)

// video delete dispatch
func VideoDeleteDispatch(dc DataChan) error {
	recs, err := dbops.GetVideoDeleteRecord(3)
	if err != nil {
		return err
	}

	if len(recs) == 0 {
		log.Info("video_delete_record is null")
	}

	for _, video := range recs {
		dc <- video.VideoUUID
	}
	return nil
}

// delete video
func DeleteVideo(vid string) error {
	err := os.Remove(VIDEO_DIR + vid)
	if err != nil {
		return err
	}

	err = dbops.DeleteVideoDeleteRecord(vid)
	if err != nil {
		return err
	}

	return nil
}

// video delete excute
func VideoDeleteExcute(dc DataChan) error {
	var errMap sync.Map

forloop:
	for {
		select {
		case vid := <-dc:
			go func(vid interface{}) {
				vidstr := vid.(string)
				err := DeleteVideo(vidstr)
				errMap.Store(vidstr, err)
			}(vid)
		default:
			break forloop
		}
	}

	errMsg := ""
	errMap.Range(func(k, v interface{}) bool {
		errMsg = fmt.Sprintf("%s\nvid is %s video delete failed, error:%s",
			errMsg, k.(string), v.(error).Error())
		return true
	})

	if errMsg == "" {
		return nil
	}
	return errors.New(errMsg)
}
