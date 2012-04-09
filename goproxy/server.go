package main

import "log"
import "net"

var server_log *log.Logger

func send_response(c *net.TCPConn, buffer []byte) {
  size, err := c.Write(buffer)
  if (err != nil) {
    server_log.Printf("%v: Connection closed", c.RemoteAddr())
    return
  }
  server_log.Printf("%v: Wrote %v bytes", c.RemoteAddr(), size)
  c.Close()
}

func handle_request(c *net.TCPConn) {
  var buffer = make([]byte, 0, TCP_BUFFER_SIZE)

  for {
    // Make the buffer TCP_BUFFER_SIZE bytes
    buffer = buffer[0:TCP_BUFFER_SIZE]

    // Read from the connection
    size, err := c.Read(buffer)
    if (err != nil) {
      server_log.Printf("%v: Connection closed", c.RemoteAddr())
      return
    }

    // Resize the buffer to the size of the data
    buffer = buffer[0:size]

    server_log.Printf("%v: Request received", c.RemoteAddr())

    // Issue the upstream request
    response := issue_upstream_request(c, buffer)

    // Send the response
    send_response(c, response)
  }
}

func server(port int) {
  l := create_listener(port)

  // Accept connections
  for {
    c, err := l.AcceptTCP()
    check_err(err)

    server_log.Printf("%v: Connection received", c.RemoteAddr())
    go handle_request(c)
  }
}

func create_listener(port int) *net.TCPListener {
  // Create a listening socket
  main_log.Printf("Starting server")
  la := create_address("tcp4", "0.0.0.0", port)
  l, err := net.ListenTCP("tcp4", la)
  check_err(err)
  return l
}

