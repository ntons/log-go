package timedrollingfile

import (
	"sync"
	"time"

	"github.com/ntons/lumberjack/v3"
)

// extract time from log content
// localTime: use local time or utc
type LogTimeFunc func(p []byte, localTime bool) time.Time

// for a short name
type Appender = TimedRollingFileAppender

// TimedRollingFileAppender based on modified lumberjack
type TimedRollingFileAppender struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`

	//
	BackupTimeFormat string `json:"backuptimeformat" yaml:"backuptimeformat"`

	// Filename define the rotation rule, filename change cause target rotate
	FilenameTimeFormat string `json:"filenametimeformat" yaml:"filenametimeformat"`

	// may be locked by logger
	DisableMutex bool `json:"disablemutex" yaml:"disablemutex"`

	// extract time from log content
	LogTime LogTimeFunc

	//
	mu sync.Mutex
	l  *lumberjack.Logger
}

func (a *TimedRollingFileAppender) Sync() error {
	if !a.DisableMutex {
		a.mu.Lock()
		defer a.mu.Unlock()
	}
	if a.l == nil {
		return nil
	}
	return a.l.Sync()
}

func (a *TimedRollingFileAppender) Close() error {
	if !a.DisableMutex {
		a.mu.Lock()
		defer a.mu.Unlock()
	}
	if a.l == nil {
		return nil
	}
	if err := a.l.Sync(); err != nil {
		return err
	}
	return a.l.Close()
}

func (a *TimedRollingFileAppender) Write(p []byte) (n int, err error) {
	if !a.DisableMutex {
		a.mu.Lock()
		defer a.mu.Unlock()
	}
	if a.l == nil {
		a.l = &lumberjack.Logger{
			MaxSize:          a.MaxSize,
			MaxBackups:       a.MaxBackups,
			LocalTime:        a.LocalTime,
			Compress:         a.Compress,
			BackupTimeFormat: a.BackupTimeFormat,
			DisableMutex:     true,
		}
	}
	var t time.Time
	if a.LogTime != nil {
		t = a.LogTime(p, a.LocalTime)
	} else {
		if a.LocalTime {
			t = time.Now()
		} else {
			t = time.Now().UTC()
		}
	}
	if filename := t.Format(a.FilenameTimeFormat); filename != a.l.Filename {
		a.l.Close()
		a.l.Filename = filename
	}
	return a.l.Write(p)
}
