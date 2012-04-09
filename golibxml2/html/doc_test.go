package html

import "testing"

func TestNew(t *testing.T) {
	New("", "")
}

func TestNewFromString(t *testing.T) {
	doc := NewFromString("<html></html>", "", "UTF-8", 0)
	if doc.String() != "<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.0 Transitional//EN\" \"http://www.w3.org/TR/REC-html40/loose.dtd\">\n<html></html>\n" {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	if New("", "").String() != "<!DOCTYPE html PUBLIC \"\" \"\">\n\n" {
		t.Fail()
	}
	if New("1", "2").String() != "<!DOCTYPE html PUBLIC \"2\" \"1\">\n\n" {
		t.Fail()
	}
}
