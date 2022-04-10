package logruslogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// NewStructuredLogger ...
func NewStructuredLogger(path, types string) func(next http.Handler) http.Handler {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	if types == "file" {
		logger.SetOutput(file)
	} else {
		logger.SetOutput(os.Stdout)
	}

	return middleware.RequestLogger(&StructuredLogger{logger})
}

// StructuredLogger ...
type StructuredLogger struct {
	Logger *logrus.Logger
}

func jsonBody(r *http.Request) (res string) {
	// Log request body
	data := map[string]interface{}{}
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &data)
	if err == nil {
		if data["password"] != nil {
			data["password"] = nil
		}

		d, _ := json.Marshal(data)
		res = string(d)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return res
}

// NewLogEntry ...
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}

	reqID := middleware.GetReqID(r.Context())
	if reqID != "" {
		r.Header["req_id"] = []string{reqID}
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	entry.Logger = entry.Logger.WithFields(logrus.Fields{
		"req_id":      reqID,
		"http_scheme": scheme,
		"http_proto":  r.Proto,
		"http_method": r.Method,
		"remote_addr": r.RemoteAddr,
		"user_agent":  r.UserAgent(),
		"uri":         fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
		"req_body":    jsonBody(r),
		"tz":          time.Now().UTC().Format(time.RFC3339),
	})

	entry.Logger.Infoln("request started")

	return entry
}

// StructuredLoggerEntry ...
type StructuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

// Write ...
func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	if status < 300 {
		l.Logger.Infoln("request complete")
	} else if status < 500 {
		l.Logger.Warnln("request warning")
	} else {
		l.Logger.Errorln("request error")
	}
}

// Panic ...
func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}

// Helper methods used by the application to get the request-scoped
// logger entry and set additional fields between handlers.
//
// This is a useful pattern to use to set state on the entry as it
// passes through the handler chain, which at any point can be logged
// with a call to .Print(), .Info(), etc.

// GetLogEntry ...
func GetLogEntry(r *http.Request) logrus.FieldLogger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}

// LogEntrySetField ...
func LogEntrySetField(r *http.Request, key string, value interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithField(key, value)
	}
}

// LogEntrySetFields ...
func LogEntrySetFields(r *http.Request, fields map[string]interface{}) {
	if entry, ok := r.Context().Value(middleware.LogEntryCtxKey).(*StructuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
	}
}
