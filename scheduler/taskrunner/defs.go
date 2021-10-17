package taskrunner

const (
	CLOSE             = "e"
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "c"

	VIDEO_DIR = "/video/"

	DELETE_INTERVAL = 2
)

type ControlChan chan string

type DataChan chan interface{}

type fn func(dc DataChan) error
