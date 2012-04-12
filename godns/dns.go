package dns

import "bytes"
import "net"
import "fmt"
//import "os"

const HEADER_LENGTH = 12

type Error error
//type Error os.Error

type Dns struct {
	server string
	cur_id uint16
	net.Conn
}

type DnsPacket struct {
	*Header
	questions[] *Question
	*Answer
}

const (
	DNS_OPCODE_QUERY = iota // RFC1035
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
	DNS_RECORD_TYPE_A = 1
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

func NewDns(server string, port int) *Dns {
	return &Dns{server: fmt.Sprint(server, ":", port), cur_id: 1}
}

func create_address(tcp string, addr string) *net.UDPAddr {
	address, err := net.ResolveUDPAddr(tcp, addr)
	if err != nil {
		panic("resolve failed")
	}
	return address
}

func (dns *Dns) Send(packet *DnsPacket) (resp[] byte, err Error) {
	conn, err := net.Dial("udp", dns.server)
	dns.Conn = conn
	if err != nil {
		return
	}

	_, err = dns.Write(packet.Bytes())
	if err != nil {
		return
	}

	resp = make([]byte, 2000)
	s, err := dns.Read(resp)
	if err != nil {
		return
	}
	resp = resp[0:s]
	return
}

func ParsePacket(buf []byte) *DnsPacket {
	dns := &DnsPacket{
		Header: ParseHeader(buf),
	}
	dns.questions = make([]*Question, dns.Header.QDCOUNT)
	for i := 0; i < int(dns.Header.QDCOUNT); i++ {
		dns.questions[i], buf = ParseQuestion(buf[HEADER_LENGTH:])
	}

	if dns.Header.Query == false {
		dns.Answer, buf = ParseAnswer(buf)
	}
	println(len(buf))
	return dns
}

func (dns *Dns) NewQuestion(rtype RecordType, domain string) *DnsPacket {
	return &DnsPacket{
		Header: &Header{
			ID: dns.cur_id,
			Query: true,
			OpCode: DNS_OPCODE_QUERY,
			Recursion: true,
			QDCOUNT: 1,
		},
		questions: []*Question{
			{
				QNAME: domain,
				QTYPE: rtype,
				QCLASS: DNS_CLASS_IN,
			},
		},
		Answer: nil,
	}
}

func (packet *DnsPacket) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(packet.Header.Bytes())
	for i := 0; i < int(packet.Header.QDCOUNT); i++ {
		buf.Write(packet.questions[i].Bytes())
	}
	if packet.Answer != nil {
		buf.Write(packet.Answer.Bytes())
	}
	return buf.Bytes()
}

func (packet *DnsPacket) String() (str string) {
	str += packet.Header.String()
	for i := 0; i < int(packet.Header.QDCOUNT); i++ {
		str += packet.questions[i].String()
	}
	if packet.Answer != nil {
		str += packet.Answer.String()
	}
	return 
}

