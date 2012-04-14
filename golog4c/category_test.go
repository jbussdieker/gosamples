package log4c

import "testing"

func TestNewCategory(t *testing.T) {
	if NewCategory("test") == nil {
		t.Fail()
	}
}

