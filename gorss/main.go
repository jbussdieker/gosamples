package main

import "syscall"

func main() {
	rusage := &syscall.Rusage{}
	ret := syscall.Getrusage(0, rusage)
	if ret == nil && rusage.Maxrss > 0 {
		println(uint64(rusage.Maxrss))
	}
}
