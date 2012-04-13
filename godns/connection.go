package dns

import "bytes"
import "net"
import "fmt"
//import "os"

type Error error
//type Error os.Error

type Connection struct {
	cur_id uint16
	net.Conn
}

type Message struct {
	*Header
	Questions[] *Question
	Answers[] *Answer
	Nameservers[] *Answer
	Additionals[] *Answer
}

type OpCodeType uint8
const (
	DNS_OPCODE_QUERY OpCodeType = iota // RFC1035
	DNS_OPCODE_IQUERY // RFC3425 (Obsolete)
	DNS_OPCODE_STATUS// RFC1035
	DNS_OPCODE_UNASSIGNED
	DNS_OPCODE_NOTIFY// RFC1996
	DNS_OPCODE_UPDATE// RFC2136
)

const (
	DNS_RCODE_NOERROR = iota
	DNS_RCODE_FORMAT_ERROR
	DNS_RCODE_SERVER_FAILURE
	DNS_RCODE_NON_EXISTANT_DOMAIN
	DNS_RCODE_NOT_IMPLEMENTED
	DNS_RCODE_QUERY_REFUSED
)

const (
	DNS_CLASS_IN = 1
)

type RecordType uint16
const (
	DNS_RECORD_TYPE_A RecordType = 1
	DNS_RECORD_TYPE_NS = iota
	DNS_RECORD_TYPE_MD
	DNS_RECORD_TYPE_MF
	DNS_RECORD_TYPE_CNAME
	DNS_RECORD_TYPE_SOA
	DNS_RECORD_TYPE_MB
	DNS_RECORD_TYPE_MG
	DNS_RECORD_TYPE_MR
	DNS_RECORD_TYPE_NULL
	DNS_RECORD_TYPE_WKS
	DNS_RECORD_TYPE_PTR
	DNS_RECORD_TYPE_HINFO
	DNS_RECORD_TYPE_MINFO
	DNS_RECORD_TYPE_MX
	DNS_RECORD_TYPE_TXT
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

