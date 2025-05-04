package golog

import (
	"bufio"
	"os"
	"sync"
)

type Logger struct {
	// file where stores logs
	file *os.File

	// for concurrent access to file
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

type Logging interface {
	// information logs
	Info(i ...any)
	Infoln(i ...any)
	Infof(format string, i ...any)

	// error logs
	Error(i ...any)
	Errorln(i ...any)
	Errorf(format string, i ...any)

	// it's using - exit(1)
	Fatalf(format string, i ...any)
	Fatalln(i ...any)
	Fatal(i ...any)
}

type Info interface {
	Info(i ...any)
	Infoln(i ...any)
	Infof(format string, i ...any)
}

type ErrorInfo interface {
	Error(i ...any)
	Errorln(i ...any)
	Errorf(format string, i ...any)
}

type FatalInfo interface {
	Fatalf(format string, i ...any)
	Fatalln(i ...any)
	Fatal(i ...any)
}
