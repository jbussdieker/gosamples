package log4c

import "testing"

func TestNewLogger(t *testing.T) {
	if NewCategory("test") == nil {
		t.Fail()
	}
}

func TestLog(t *testing.T) {
	Init()
	logger := NewCategory("stdout")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	logger.Log("Testing123")
	Free()
}

