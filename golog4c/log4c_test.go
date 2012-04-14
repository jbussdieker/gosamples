package log4c

import "testing"

func TestNew(t *testing.T) {
	if New() == nil {
		t.Fail()
	}
}

func TestFirst(t *testing.T) {
	basic := GetLayout("dated")
	stderr := GetAppender("stderr")
	stderr.SetLayout(basic)
	category := NewCategory("test")
	category.SetAppender(stderr)
	category.Logf("Test log")
}

