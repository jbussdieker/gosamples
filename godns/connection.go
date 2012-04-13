package dns

import "net"
import "fmt"
//import "os"

type Error error
//type Error os.Error

type Connection struct {
	cur_id uint16
	net.Conn
}

const (
	DNS_RCODE_NOERROR = iota
	DNS_RCODE_FORMAT_ERROR
	DNS_RCODE_SERVER_FAILURE
	DNS_RCODE_NON_EXISTANT_DOMAIN
	DNS_RCODE_NOT_IMPLEMENTED
	DNS_RCODE_QUERY_REFUSED
)

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func NewConnection(server string, port int) (conn *Connection, err error) {
	udpconn, err := net.Dial("udp", fmt.Sprint(server, ":", port))
	if err != nil {
		return
	}
	conn = &Connection{
		cur_id: 1, 
		Conn: udpconn,
	}
	return
}

func (conn *Connection) Send(message *Message) (resp *Message, err Error) {
	_, err = conn.Write(message.Bytes())
	if err != nil {
		return
	}

	buf := make([]byte, 2000)
	s, err := conn.Read(buf)
	if err != nil {
		return
	}
	buf = buf[0:s]
	resp, err = ParseMessage(buf)
	return
}

