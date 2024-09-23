package golog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func New() *Logger {
	return &Logger{
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

	fmt.Print("INFO ["+now+"]", txt)
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

	fmt.Println("INFO ["+now+"]", txt)
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

	fmt.Printf(format, "INFO ["+now+"]", txt)
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

	fmt.Print("ERROR ["+now+"]", txt)
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

	fmt.Println("ERROR ["+now+"]", txt)
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

	fmt.Printf(format, "ERROR ["+now+"]", txt)
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

	fmt.Print("FATAL ["+now+"]", txt)
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

	fmt.Println("FATAL ["+now+"]", txt)
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

	fmt.Printf(format, "FATAL ["+now+"]", txt)
	os.Exit(1)
}

func (l *Logger) SetJSONformat(data *Logger) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		l.Errorln(err)
		return ""
	}
	return string(jsonData)
}
