package logger

import (
	"os"
	"path/filepath"
	"time"
)

const fileformat = "2006010215.json"

type fileWriter struct {
	logDir  string
	tick    time.Duration
	current int
	file    *os.File
}

func newFileWriter(logDir string) *fileWriter {
	current := time.Now().Hour()
	path := filepath.Join(logDir, time.Now().Format(fileformat))

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &fileWriter{
		current: current,
		logDir:  logDir,
		file:    file,
	}
}

func (f *fileWriter) Write(p []byte) (n int, err error) {
	if time.Now().Hour() == f.current {
		return f.file.Write(p)
	}

	f.current = time.Now().Hour()
	path := filepath.Join(f.logDir, time.Now().Format(fileformat))

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}

	f.file = file
	return f.file.Write(p)
}
