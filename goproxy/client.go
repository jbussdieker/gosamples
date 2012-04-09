package main

import "log"
import "bytes"
import "net"
import "fmt"
import "strings"

var upstream_log *log.Logger

func read_content_length(headers string) int {
  var content_length int;

  //upstream_log.Printf("Response Headers: \n\n%v\n\n", headers)
  tmpstr := strings.SplitAfterN(headers, "Content-Length: ", 2)
  if (len(tmpstr) > 1) {
    fmt.Sscanf(tmpstr[1], "%d", &content_length);
  }

  return content_length + len(headers)
}

func concat(old1, old2 []byte) []byte {
   newslice := make([]byte, len(old1) + len(old2))
   copy(newslice, old1) 
   copy(newslice[len(old1):], old2)
   return newslice
}

func read_response(dc *net.TCPConn, c *net.TCPConn) []byte {
  var buffer = make([]byte, 0, TCP_BUFFER_SIZE)
  var fbuffer = make([]byte, 0)

  content_length := -1
  cursize := 0

  for {
    // Size the buffer
    buffer = buffer[0:TCP_BUFFER_SIZE]

    // Read from the connection
    size, err := c.Read(buffer)
    if (err != nil) {
      upstream_log.Printf("%v=>%v: Connection closed", dc.RemoteAddr(), c.RemoteAddr())
      return buffer
    }

    // Resize the buffer
    buffer = buffer[0:size]
    fbuffer = concat(fbuffer, buffer)
    cursize += size

    //if (content_length == -1) {
      split_buffer := bytes.SplitAfterN(fbuffer, []byte("\r\n\r\n"), 2)
      if (len(split_buffer) > 1) {
        content_length = read_content_length(string(split_buffer[0]))
        upstream_log.Printf("%v=>%v: Content-Length: %v", dc.RemoteAddr(), c.RemoteAddr(), content_length);
      }
    //}

    if (content_length == cursize) {
      break
    } else {
      upstream_log.Printf("Packet incomplete %v != %v", content_length, cursize)
    }
  }


  //upstream_log.Printf("%v=>%v: Response Received:\n%v", dc.RemoteAddr(), c.RemoteAddr(), string(buffer))

  return fbuffer
}

func send_request(dc *net.TCPConn, c *net.TCPConn, buffer []byte) {
  size, err := c.Write(buffer)
  if (err != nil) {
    upstream_log.Printf("%v: Connection closed", c.RemoteAddr())
    return
  }
  upstream_log.Printf("%v: Wrote %v bytes", c.RemoteAddr(), size)
}

func issue_upstream_request(dc *net.TCPConn, buffer []byte) []byte {
  //upstream_log.Printf("%v=>nil: Request Header: \n\n%v", dc.RemoteAddr(), string(buffer))

  c, err := net.DialTCP("tcp4", create_address("tcp4", "0.0.0.0", 0), create_address("tcp4", "localhost", 8080))
  check_err(err);

  upstream_log.Printf("%v=>%v: Connected", dc.RemoteAddr(), c.RemoteAddr())

  send_request(dc, c, buffer);
 
  upstream_log.Printf("%v=>%v: Request Sent", dc.RemoteAddr(), c.RemoteAddr())

  response := read_response(dc, c)

  return response
}

