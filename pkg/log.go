package pkg

import (
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
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Errorlln(err)
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

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("INFO [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Print("INFO ["+now+"] ", txt)
}

func (l *Logger) Infoln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("INFO [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Println("INFO ["+now+"] ", txt)
}

func (l *Logger) Infof(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("INFO [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Printf(format, "INFO ["+now+"] ", txt)
}

func (l *Logger) Error(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("ERROR [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Print("ERROR ["+now+"] ", txt)
}

func (l *Logger) Errorlln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("ERROR [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Println("ERROR ["+now+"] ", txt)
}

func (l *Logger) Errorf(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("ERROR [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Printf("ERROR ["+now+"] ", txt)
}

func (l *Logger) Fatal(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("FATAL [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Print("FATAL ["+now+"] ", txt)
	os.Exit(1)
}

func (l *Logger) Fatalln(i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("FATAL [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Println("FATAL ["+now+"] ", txt)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, i ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.tm.Format(time.DateTime)
	txt := fmt.Sprint(i...)

	if l.isFileExist {
		var err error
		l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.Errorlln("Ошибка-открытия-json")
		}
		defer l.file.Close()

		if _, err := l.file.WriteString("FATAL [" + now + "] " + txt); err != nil {
			return
		}
	}

	fmt.Printf(format, "FATAL ["+now+"] ", txt)
	os.Exit(1)
}