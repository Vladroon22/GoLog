package golog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type httpRequest *http.Request

func New() *Logger {
	return &Logger{
		IsHttpDebug: false,
		isFileExist: false,
		tm:          time.Now(),
	}
}

func NewWithJSON() *Logger {
	return &Logger{
		IsHttpDebug: true,
		isFileExist: false,
		tm:          time.Now(),
	}
}

// set log's data into file
func (l *Logger) SetOutput(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		l.Errorln(err)
		return nil, err
	}
	defer file.Close()
	l.file = file
	l.isFileExist = true
	return file, nil
}

func (l *Logger) Info(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "INFO"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Infoln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "INFO"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Infof(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "INFO"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Error(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "ERROR"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Errorln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "ERROR"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Errorf(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "ERROR"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Fatal(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "FATAL"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) Fatalln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "FATAL"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}

	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsHttpDebug {
		for _, item := range i {
			if req, ok := item.(*httpRequest); ok {
				l.httplog(*req)
			}
		}
	}

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)
	l.logLevel = "FATAL"
	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
			return
		}
	}
	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) httplog(r *http.Request) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	l.logLevel = "DEBUG"
	debug := &httpDebug{
		tm:       now,
		logLevel: l.logLevel,
		message:  r.Response.Request.Body,
		method:   r.Method,
		status:   r.Response.Status,
		path:     r.URL.Path,
	}

	jsonData, err := json.MarshalIndent(debug, "", " ")
	if err != nil {
		l.Errorln(err)
		return
	}

	colorB := "\033[32m"
	colorE := "\033[0m"

	if l.IsHttpDebug {
		if l.isFileExist {
			if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + string(jsonData) + "\n"); err != nil {
				return
			}
		}
		fmt.Println(colorB + string(jsonData) + colorE)
	}
}
