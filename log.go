package golog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func New() *Logger {
	return &Logger{
		IsDebug:     false,
		isFileExist: false,
		tm:          time.Now(),
	}
}

func (l *Logger) SetOutput(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		l.Errorln(err)
		return nil, err
	}
	defer l.file.Close()
	l.file = file
	l.isFileExist = true
	return file, nil
}

func (l *Logger) Info(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

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

func (l *Logger) httpDebug(r *http.Request, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	l.logLevel = "DEBUG"

	debug := &Debug{
		tm:           now,
		logLevel:     l.logLevel,
		errorMessage: err.Error(),
		method:       r.Method,
		status:       r.Response.Status,
		path:         r.URL.Path,
	}

	jsonData, err := json.MarshalIndent(debug, "", " ")
	if err != nil {
		l.Errorln(err)
		return
	}

	colorB := "\033[32m"
	colorE := "\033[0m"

	if l.isFileExist {
		if _, err := l.file.WriteString(l.logLevel + " [" + now + "] " + err.Error() + "\n"); err != nil {
			return
		}
	}
	if l.IsDebug {
		fmt.Println(colorB + string(jsonData) + colorE)
	}
}

func (l *Logger) SetJSONformat(data *LogFields) string {
	data.Data = make(map[string]string)
	jsonData, err := json.Marshal(data)
	if err != nil {
		l.Errorln(err)
		return ""
	}
	return string(jsonData)
}
