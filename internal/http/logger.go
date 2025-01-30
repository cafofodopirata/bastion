package http

import log "github.com/sirupsen/logrus"

// logrusRunLogger implements the RunLogger interface using logrus logging
// package.
type logrusRunLogger struct{}

// Printf logs a formatted message using logrus Printf method.
func (l *logrusRunLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Fatalf logs a formatted message and then terminates the program using logrus
// package's Fatalf method.
func (l *logrusRunLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
