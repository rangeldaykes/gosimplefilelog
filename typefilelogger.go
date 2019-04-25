package gosimplefilelog

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// FileLogger provides output file to write logs
type FileLogger struct {
	// full paht name of file to write
	filename string
	// compose file name with year/month/day/hour
	datefile time.Time
	// path where the file will be created
	path string
	// pointer to logger std lib
	logger *log.Logger
	// pointer to filo on disk
	logFile *os.File
}

// NewFileLogger Creates a new logger to file
func NewFileLogger(nameoffile string, pathoffile string) *FileLogger {

	f := FileLogger{filename: "Log.txt"}

	if nameoffile != "" {
		f.filename = nameoffile
	}

	if pathoffile != "" {
		_, err := os.Stat(pathoffile)
		if !os.IsNotExist(err) {
			f.path = pathoffile
		} else {
			f.path = getPathDefault()
		}
	} else {
		f.path = getPathDefault()
	}

	f.start()
	return &f
}

// start initiate a new file per hour to write log
func (f *FileLogger) start() {
	f.datefile = time.Now()

	fileyear := f.datefile.Year()
	filemonth := f.datefile.Month()
	fileday := f.datefile.Day()
	filehour := f.datefile.Hour()

	filename := strconv.Itoa(fileyear) +
		strconv.Itoa(int(filemonth)) +
		strconv.Itoa(fileday) + "_" +
		strconv.Itoa(filehour) + "_" +
		f.filename

	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}

	file := filepath.Join(f.path, filename)

	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	f.logFile = logFile

	f.logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

}

// Log implements the interface ITypeLogger - write log to file
func (f *FileLogger) Log(level LogLevel, message string) {
	if f.logFile == nil {
		return
	}

	if time.Now().Hour() != f.datefile.Hour() {
		f.Close()
		f.start()
	}

	if f.logger != nil {
		f.logger.Println(message)
	}
}

// Close releases resources
func (f *FileLogger) Close() error {
	err := f.logFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func getPathCurrent() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return dir
}

func getPathDefault() string {
	path := filepath.Join(getPathCurrent(), "Log")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	return path
}
