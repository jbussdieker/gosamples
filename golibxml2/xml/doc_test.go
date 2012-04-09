package xml

import "testing"

func TestNew(t *testing.T) {
	New("1.0")
}

func TestNewFromString(t *testing.T) {
	doc := NewFromString("<html></html>", "", "UTF-8", 0)
	if doc.String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<html/>\n" {
		t.Fatal(doc.String())
		t.Fail()
	}
}

func TestString(t *testing.T) {
	if New("1.0").String() != "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" {
		t.Fail()
	}
	if New("9.9").String() != "<?xml version=\"9.9\" encoding=\"UTF-8\"?>\n" {
		t.Fail()
	}
}
