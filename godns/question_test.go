package dns

import "testing"

func TestQuestion1(t *testing.T) {
	expected := []byte{3, 119, 119, 119, 12, 110, 111, 114, 116, 104, 101, 97, 115, 116, 101, 114, 110, 3, 101, 100, 117, 0, 0, 1, 0, 1}
	q := NewQuestion("www.northeastern.edu", RECORD_TYPE_A, CLASS_IN)
	if string(q.Bytes()) != string(expected) {
		t.Error("Got:", q.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}

func TestQuestion2(t *testing.T) {
	expected := []byte{6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0, 0, 16, 0, 1}
	q := NewQuestion("google.com", RECORD_TYPE_TXT, CLASS_IN)
	if string(q.Bytes()) != string(expected) {
		t.Error("Got:", q.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}

