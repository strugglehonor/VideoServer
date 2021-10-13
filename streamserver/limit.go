package main

import (
	log "github.com/sirupsen/logrus"
)

//limit rate: bucket token算法
type ConnLimiter struct {
	limitVal int
	bucket   chan int
}

// new ConnLimiter
func NewConnLimiter(limitVal int) *ConnLimiter {
	return &ConnLimiter{
		limitVal: limitVal,
		bucket:   make(chan int, limitVal),
	}
}

// GetConnLimiter
func (c *ConnLimiter) GetConnLimiter() bool {
	if c.limitVal < len(c.bucket) {
		log.Fatal("Reached the rate limitation.")
		return false
	}

	c.bucket <- 1
	return true
}

// ReleaseConnLimit
func (c *ConnLimiter) ReleaseConnLimit() {
	c1 := <-c.bucket
	log.Info("Connection is comming: %d", c1)
}
