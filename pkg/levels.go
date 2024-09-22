package pkg

import (
	"os"
	"sync"
	"time"
)

type Logger struct {
	file        *os.File
	tm          time.Time
	mu          sync.Mutex
	isFileExist bool
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
}
