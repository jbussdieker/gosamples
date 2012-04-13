package main

import "os"

func main() {
	if len(os.Args) < 2 {
		println("Invalid args")
		println(os.Args[0], "aaaa")
	}

	
}
