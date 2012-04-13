package main

import "os"
import "fmt"
import . "godns"

func main() {
	if len(os.Args) < 3 {
		println("Invalid args")
		println(os.Args[0], " [dns server] [domain]")
	}

	conn, err := NewConnection(os.Args[1], 53)
	if err != nil {
		println("Error connecting:", err.Error())
		os.Exit(1)
	}
	packet := conn.NewQuestion(DNS_RECORD_TYPE_TXT, os.Args[2])
	println(packet.String())
	resp, err := conn.Send(packet)
	if err != nil {
		println("Error sending:", err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.String())
}
