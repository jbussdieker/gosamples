package log4c

import "testing"

func TestGetLayout(t *testing.T) {
	if GetLayout("dated") == nil {
		t.Fail()
	}
}

