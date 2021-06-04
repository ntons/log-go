package timedrollingfile

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const Scheme = "trf"

func Open(s string) (l *TimedRollingFileAppender, err error) {
	u, err := url.Parse(s)
	if err != nil {
		return
	}
	return OpenURL(u)
}

func OpenURL(u *url.URL) (l *TimedRollingFileAppender, err error) {
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	for k, v := range q {
		q[strings.ToLower(k)] = v
	}
	if u.Path == "" {
		return nil, fmt.Errorf("bad filepath %q", u.Path)
	}
	l = &TimedRollingFileAppender{
		FilenameTimeFormat: u.Path,
	}
	if s := q.Get("maxsize"); s != "" {
		if l.MaxSize, err = strconv.Atoi(s); err != nil {
			return nil, fmt.Errorf("bad maxsize %q", s)
		}
	}
	if s := q.Get("maxbackups"); s != "" {
		if l.MaxBackups, err = strconv.Atoi(s); err != nil {
			return nil, fmt.Errorf("bad maxbackups %q", s)
		}
	}
	if s := q.Get("localtime"); s == "1" || s == "true" {
		l.LocalTime = true
	}
	if s := q.Get("compress"); s == "1" || s == "true" {
		l.Compress = true
	}
	if s := q.Get("backuptimeformat"); s != "" {
		l.BackupTimeFormat = s
	}
	return
}
