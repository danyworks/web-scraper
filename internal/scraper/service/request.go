package service

import (
	"strconv"
	"time"
)

type Request struct {
	Status  OPStatus          `json:"status"`
	Data    map[string]string `json:"data"`
	Timeout string            `json:"timeout"`
}

func (R *Request) GetTimeout() time.Duration {
	var timeout time.Duration
	sec, _ := strconv.Atoi(R.Timeout)
	timeout = time.Duration(sec) * time.Second
	return timeout
}
