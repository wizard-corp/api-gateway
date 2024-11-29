package mylogger

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type logEntry struct {
	Time    time.Time    `json:"time"`
	Level   logrus.Level `json:"level"`
	Message string       `json:"message"`
	File    string       `json:"file"`
	Line    int          `json:"line"`
}

type csvFormatter struct {
	writer *csv.Writer
}

func (f *csvFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logEntry := logEntry{
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
		File:    filepath.Base(entry.Caller.File), // Extract filename
		Line:    entry.Caller.Line,
	}
	var record []string
	record = append(record, logEntry.Time.Format("2006-01-02 15:04:05")) // Format timestamp
	record = append(record, logEntry.Level.String())
	record = append(record, logEntry.Message)
	record = append(record, logEntry.File)
	record = append(record, fmt.Sprint(logEntry.Line))

	err := f.writer.Write(record)
	if err != nil {
		return nil, err
	}

	f.writer.Flush() // Ensure the record is written to the file
	return nil, nil
}

func newCSVFormatter(w io.Writer) *csvFormatter {
	formatter := &csvFormatter{writer: csv.NewWriter(w)}
	formatter.writer.Comma = ',' // Use comma as delimiter (optional)
	return formatter
}

func SetupLogger(filename string, level string) error {
	// Create or open the log file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Set up the logger
	logger := logrus.New()
	logger.SetReportCaller(true) // Enable line numbers and file names
	logger.SetOutput(file)       // Set log output to file

	// Create and set the custom CSV formatter
	formatter := newCSVFormatter(file)
	logger.SetFormatter(formatter)

	// Set the log level
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		parsedLevel = logrus.InfoLevel // Default to Info level if parsing fails
		logger.WithError(err).Printf("Invalid log level, defaulting to %s", parsedLevel)
	}
	logger.SetLevel(parsedLevel)
	logger.Printf("Log level set to %s", logger.GetLevel().String())

	return nil
}
