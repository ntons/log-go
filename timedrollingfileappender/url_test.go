package timedrollingfile

import (
	"fmt"
	"net/url"
	"testing"
)

func ExpectEqual(name string, expected, actually interface{}) error {
	if expected != actually {
		return fmt.Errorf(
			"%q is expected to be %q but actually %q", name, expected, actually)
	}
	return nil
}

func TestOpen(t *testing.T) {
	const (
		FilenameTimeFormat = "/tmp/timedrollingfile/2006/01/0215.log"
		MaxSize            = 1024
		MaxBackups         = 10
		LocalTime          = true
		Compress           = false
		BackupTimeFormat   = "0405.000"
	)
	s := fmt.Sprintf(
		"%s://%v?maxsize=%v&maxbackups=%v&localtime=%t&compress=%t&backupTimeFormat=%v",
		Scheme, FilenameTimeFormat, MaxSize, MaxBackups, LocalTime,
		Compress, url.QueryEscape(BackupTimeFormat))
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("url.Parse fail: %q", err)
	}
	l, err := Open(u)
	if err != nil {
		t.Fatalf("Open fail: %q", err)
	}
	if err := ExpectEqual(
		"FilenameTimeFormat",
		FilenameTimeFormat,
		l.FilenameTimeFormat); err != nil {
		t.Error(err.Error())
	}
	if err := ExpectEqual("MaxSize", MaxSize, l.MaxSize); err != nil {
		t.Error(err.Error())
	}
	if err := ExpectEqual("MaxBackups", MaxBackups, l.MaxBackups); err != nil {
		t.Error(err.Error())
	}
	if err := ExpectEqual("LocalTime", LocalTime, l.LocalTime); err != nil {
		t.Error(err.Error())
	}
	if err := ExpectEqual("Compress", Compress, l.Compress); err != nil {
		t.Error(err.Error())
	}
	if err := ExpectEqual("BackupTimeFormat", BackupTimeFormat, l.BackupTimeFormat); err != nil {
		t.Error(err.Error())
	}
}
