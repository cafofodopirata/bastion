package http

import (
	log "github.com/sirupsen/logrus"
)

// logrusRunLogger implements the RunLogger interface using logrus logging
// package.
type logrusRunLogger struct{}

func (l *logrusRunLogger) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

func (l *logrusRunLogger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

// Printf logs a formatted message using logrus Printf method.
func (l *logrusRunLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Fatalf logs a formatted message and then terminates the program using logrus
// package's Fatalf method.
func (l *logrusRunLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (l *logrusRunLogger) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}
