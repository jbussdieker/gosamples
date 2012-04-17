package main

import "testing"

const TEST_CERT   = "/home/jbussdieker/.ec2/cert-62VJ7DAYEMXEINZQDU5UVMOB6IOXOSHC.pem"
const TEST_KEY    = "/home/jbussdieker/.ec2/pk-62VJ7DAYEMXEINZQDU5UVMOB6IOXOSHC.pem"
const TEST_REGION = "us-west-1"

func newAws(t *testing.T) *Aws {
	aws := NewAws(TEST_CERT, TEST_KEY, TEST_REGION)
	if aws == nil {
		t.Fail()
	}
	return aws
}

func TestNewAws(t *testing.T) {
	newAws(t)
}

func TestReadKeyPairs(t *testing.T) {
	aws := newAws(t)
	keys, err := aws.DescribeKeyPairs()
	if err != nil {
		t.Fatal(err)
	}
	if keys == nil {
		t.Fatal("nil keys")
	}
	t.Log(keys)
}

