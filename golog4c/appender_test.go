package log4c

import "testing"

func TestNewAppender(t *testing.T) {
	if NewAppender("stdout") == nil {
		t.Fail()
	}
}

func TestGetAppender(t *testing.T) {
	if GetAppender("stdout") == nil {
		t.Fail()
	}
}

