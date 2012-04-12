package dns

import "testing"

func TestNewHeader(t *testing.T) {
	expected := []byte{219, 66, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0}
	h := NewHeader()
	h.ID = 0xdb42
	h.Query = true
	h.Recursion = true
	h.QDCOUNT = 1
	if string(h.Bytes()) != string(expected) {
		t.Error("Got:", h.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}

func TestParseHeader(t *testing.T) {
	//expected := []byte{219, 66, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0}
	//expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expected := []byte{157, 186, 38, 46, 237, 158, 76, 138, 117, 39, 157, 33}
	h := ParseHeader(expected)

	if string(h.Bytes()) != string(expected) {
		t.Error("Got:", h.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}

