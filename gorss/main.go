package main

import "syscall"

func GetRSS() uint64 {
	rusage := &syscall.Rusage{}
	ret := syscall.Getrusage(0, rusage)
	if ret == nil && rusage.Maxrss > 0 {
		return uint64(rusage.Maxrss)
	}
	return 0
}

func main() {
	println(GetRSS())
}
