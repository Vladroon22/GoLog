package golog

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type Logger struct {
	// file where stores logs
	file *os.File

	// current time of this log
	tm time.Time `json:"time"`

	mu sync.Mutex

	// check file's existing
	isFileExist bool `json:"fileIs"`

	// level of logging
	logLevel string `json:"level"`

	// enables to see http errors
	IsHttpDebug bool `json:"IsHttpDebug"`

	// for writting in file
	writer *bufio.Writer
}

type httpDebug struct {
	// current time of log
	tm string `json:"time"`

	// current log level
	logLevel string `json:"level"`

	// log message
	message any `json:"msg"`

	// http-method
	method string `json:"method"`

	// http status
	status string `json:"httpStatus"`

	// url path of request
	path string `json:"url"`
}
