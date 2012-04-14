package log4c

import "testing"

func TestInit(t *testing.T) {
	if Init() != 0 {
		t.Fail()
	}
}

func TestFirst(t *testing.T) {
	Init()
	//appender := GetAppender("stdout")
	category := NewCategory("six13log.log.app.application1")

	//category.SetAppender(appender)
	//category.SetPriority(ERROR)
	category.Log("Testing logger\n")
	Free()
}

