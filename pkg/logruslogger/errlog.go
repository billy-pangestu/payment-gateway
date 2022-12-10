package logruslogger

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// TOPIC for setting topic of log
	TOPIC = "payment-gateway-backend"
	// ErrorLevel ...
	ErrorLevel = log.ErrorLevel
	// WarnLevel ...
	WarnLevel = log.WarnLevel
	// InfoLevel ...
	InfoLevel = log.InfoLevel
)

// LogContext function for logging the context of echo
// c string context
// s string scope
func LogContext(c string, s string, cor ...interface{}) *log.Entry {
	topic, ok := os.LookupEnv("LOG_TOPIC")
	if !ok {
		topic = TOPIC
	}
	entry := log.WithFields(log.Fields{
		"topic":   topic,
		"context": c,
		"scope":   s,
		"tz":      time.Now().UTC().Format(time.RFC3339),
	})
	if len(cor) > 0 {
		if cor[0] != nil {
			entry = entry.WithFields(
				log.Fields{
					"req_id": fmt.Sprintf("%+v", cor[0]),
				})
		}
	}
	return entry
}

// Log function for returning entry type
// level log.Level
// message string message of log
// context string context of log
// scope string scope of log
func Log(level log.Level, message string, context string, scope string, corr ...interface{}) {
	log.SetFormatter(&log.JSONFormatter{})
	// append optional correlation id to logger
	var correlation interface{}
	if len(corr) > 0 {
		correlation = corr[0]
	}
	entry := LogContext(context, scope, correlation)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}
