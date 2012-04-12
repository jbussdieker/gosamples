package dns

import "testing"

func TestQuestion(t *testing.T) {
	expected := []byte{3, 119, 119, 119, 12, 110, 111, 114, 116, 104, 101, 97, 115, 116, 101, 114, 110, 3, 101, 100, 117, 0, 0, 1, 0, 1}
	q := Question{
		QNAME: "www.northeastern.edu",
		QTYPE: 1,
		QCLASS: 1,
	}
	if string(q.Bytes()) != string(expected) {
		t.Error("Got:", q.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}
