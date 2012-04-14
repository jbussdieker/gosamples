package main

import "code.google.com/p/log4go"

func main() {
	logger := log4go.NewLogger()
	logger.Info("Test")
}
