package utils

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func NewUUID() string{
	u1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(u1[:])
}