package gocarbon

import "testing"

const TEST_HOST = "localhost"
const TEST_PORT = "2003"
const TEST_PREFIX = "josh"

func TestDialTCP(t *testing.T) {
	conn := DialTCP(TEST_HOST, TEST_PORT, TEST_PREFIX)
	if conn == nil {
		t.Fatal("Nil connection")
	}
}

func TestDialUDP(t *testing.T) {
	conn := DialTCP(TEST_HOST, TEST_PORT, TEST_PREFIX)
	if conn == nil {
		t.Fatal("Nil connection")
	}
}

func TestWriteTCP(t *testing.T) {
	conn := DialTCP(TEST_HOST, TEST_PORT, TEST_PREFIX)
	if conn == nil {
		t.Fatal("Nil connection")
	}
	conn.Write("gocarbon", 5)
}

func TestWriteUDP(t *testing.T) {
	conn := DialUDP(TEST_HOST, TEST_PORT, TEST_PREFIX)
	if conn == nil {
		t.Fatal("Nil connection")
	}
	conn.Write("gocarbon", 5)
}

