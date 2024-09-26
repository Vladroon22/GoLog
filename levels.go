package golog

import (
	"net/http"
	"os"
	"sync"
	"time"
)

type Logger struct {
	file        *os.File   `json:"-"`
	tm          time.Time  `json:"time"`
	mu          sync.Mutex `json:"-"`
	isFileExist bool       `json:"fileIs"`
	logLevel    string     `json:"level"`
	IsDebug     bool       `json:"IsDebug"`
}

type Debug struct {
	tm           string `json:"time"`
	logLevel     string `json:"level"`
	errorMessage any    `json:"msg"`
	method       string `json:"method"`
	status       string `json:"httpStatus"`
	path         string `json:"url"`
}

type LogFields struct {
	Data map[string]string
}

type FatalErrors interface {
	Fatal(...any)
	Fatalln(...any)
	Fatalf(string, ...any)
}

type Errors interface {
	Error(...any)
	Errorlln(...any)
	Errorf(string, ...any)
}

type Regular interface {
	Info(...any)
	Infoln(...any)
	Infof(string, ...any)
}

type DebugFunc interface {
	httpDebug(*http.Request, error)
}

type DataSetUP interface {
	SetOutput(string)
	SetJSONformat(*LogFields)
}
