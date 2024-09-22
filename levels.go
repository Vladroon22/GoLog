package golog

import (
	"os"
	"sync"
	"time"
)

type Logger struct {
	file        *os.File
	tm          time.Time `json:"time"`
	mu          sync.Mutex
	isFileExist bool   `json:"fileIs"`
	logLevel    string `json:"level"`
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

type DataSetUP interface {
	SetOutput(string)
	SetJSONformat(JSONformat)
}
