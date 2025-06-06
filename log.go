package golog

import (
	"bufio"
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
	}
}

func NewWithJSON() *Logger {
	return &Logger{
		IsHttpDebug: true,
		isFileExist: false,
	}
}

// set log's data into file
func (l *Logger) SetOutput(filename string) (*os.File, error) {
	var err error
	l.file, err = os.Create(filename)
	if err != nil {
		l.Errorln(err)
		return nil, err
	}
	l.isFileExist = true
	l.writer = bufio.NewWriter(l.file)
	return l.file, nil
}

func writeToFile(l *Logger, loglevel, now, txt string) {
	if !l.isFileExist {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	defer l.writer.Flush()

	l.logLevel = loglevel
	if _, err := l.writer.WriteString(l.logLevel + " [" + now + "] " + txt + "\n"); err != nil {
		fmt.Println("Failed to write to log file:", err)
		return
	}
}

func httpLog(l *Logger, i ...any) {
	if !l.IsHttpDebug {
		return
	}

	for _, item := range i {
		if req, ok := item.(*httpRequest); ok {
			l.httplog(*req)
		}
	}
}

func (l *Logger) Info(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "INFO", now, txt)

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Infoln(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "INFO", now, txt)

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Infof(format string, i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "INFO", now, txt)

	colorB := "\033[34m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Error(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "ERROR", now, txt)

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Errorln(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "ERROR", now, txt)

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Errorf(format string, i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "ERROR", now, txt)

	colorB := "\033[33m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
}

func (l *Logger) Fatal(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "ERROR", now, txt)

	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Print(colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) Fatalln(i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "FATAL", now, txt)

	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Println(colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, i ...any) {
	go httpLog(l, i...)

	now := time.Now().Format(time.DateTime)
	txt := fmt.Sprint(i...)

	writeToFile(l, "FATAL", now, txt)

	colorB := "\033[31m"
	colorE := "\033[0m"
	fmt.Printf(format, colorB+l.logLevel+colorE+" ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) httplog(r *http.Request) {
	now := time.Now().Format(time.DateTime)
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

	if l.isFileExist {
		l.mu.Lock()
		if _, err := l.writer.WriteString(l.logLevel + " [" + now + "] " + string(jsonData) + "\n"); err != nil {
			return
		}
		l.mu.Unlock()
	}
	fmt.Println(colorB + string(jsonData) + colorE)
}
