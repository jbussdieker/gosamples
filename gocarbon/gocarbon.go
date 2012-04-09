package gocarbon

import "net"
import "strconv"
import "time"
import "io"

type Connection interface {
	Write(metric string, value interface{})
}

type connection struct {
	host string
	port string
	prefix string
	conn interface{}
}

func DialTCP(host string, port string, prefix string) Connection {
	c := &connection{}
	c.host = host
	c.port = port
	c.prefix = prefix
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(c.host, c.port))
	if err != nil {
		panic(err.String())
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err.String())
	}
	c.conn = conn
	return c
}

func DialUDP(host string, port string, prefix string) Connection {
	c := &connection{}
	c.host = host
	c.port = port
	c.prefix = prefix
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(c.host, c.port))
	if err != nil {
		panic(err.String())
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err.String())
	}
	c.conn = conn
	return c
}

func (c *connection) Write(metric string, value interface{}) {
	var outvalue int
	switch v := value.(type) {
		case int:
			outvalue = v
			break
		default:
			println("Can't use")
	}
	outbuf := c.prefix + "." + metric + " " + strconv.Itoa(outvalue) + " " + strconv.Itoa64(time.Seconds()) + "\n"
	size, err := c.conn.(io.Writer).Write([]byte(outbuf))
	if err != nil {
		panic(err.String())
	}
	if size != len(outbuf) {
		panic("unset bytes")
	}
}

