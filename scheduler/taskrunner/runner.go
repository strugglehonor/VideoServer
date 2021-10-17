package taskrunner

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type Runner struct {
	Control     ControlChan
	Error       ControlChan
	IsLongLived bool
	Data        DataChan
	Dispatch    fn
	Execute     fn
}

// return *Runner
func NewRunner(isLongLived bool, dispatch, execute fn, dataNum int) *Runner {
	return &Runner{
		Control:     make(ControlChan, 1),
		Error:       make(ControlChan),
		IsLongLived: isLongLived,
		Data:        make(DataChan, dataNum),
		Dispatch:    dispatch,
		Execute:     execute,
	}
}

// start dispatch (producer)
func (runner *Runner) StartDispatch() error {
	defer func() {
		if !runner.IsLongLived {
			close(runner.Control)
			close(runner.Data)
			close(runner.Error)
		}
	}()
	// forloop
	for {
		select {
		case c := <-runner.Control:
			if c == READY_TO_DISPATCH {
				err := runner.Dispatch(runner.Data)
				if err != nil {
					log.Error(err)
					runner.Error <- CLOSE
					return err
				}
				runner.Control <- READY_TO_EXECUTE
			}
			if c == READY_TO_EXECUTE {
				err := runner.Execute(runner.Data)
				if err != nil {
					log.Error(err)
					runner.Error <- CLOSE
					return err
				}
				runner.Control <- READY_TO_DISPATCH
			}
		case e := <-runner.Error:
			if e == CLOSE {
				log.Error("Runner.Error has CLOSE Val")
				return errors.New("Runner.Error has CLOSE Val")
			}

		}
	}
}

// start
func (runner *Runner) StartAll() {
	runner.Control <- READY_TO_DISPATCH
	runner.StartDispatch()
}
