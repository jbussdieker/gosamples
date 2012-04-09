package main

import "log"
import "net"
import "os"
import "fmt"

const TCP_BUFFER_SIZE int = 4096 * 4

func create_logger(name string) *log.Logger {
  full_name := fmt.Sprintf("(%v): ", name)
  return log.New(os.Stdout , full_name, 0)
}

func check_err(err os.Error) {
  if (err != nil) {
    main_log.Fatalf("ERROR: %v", err)
  }
}

func create_address(tcp string, ip string, port int) *net.TCPAddr {
  addr := fmt.Sprintf("%v:%v", ip, port)
  address, err := net.ResolveTCPAddr(tcp, addr)
  check_err(err)
  return address
}

