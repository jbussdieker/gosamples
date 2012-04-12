package dns

import "bytes"
import "net"
import "fmt"

const HEADER_LENGTH = 12

type Dns struct {
	server string
	cur_id uint16
	net.Conn
}

type DnsPacket struct {
	*Header
	*Question
	*Answer
}

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

func (dns *Dns) Send(packet *DnsPacket) (resp[] byte, err error) {
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
	dns.Question, buf = ParseQuestion(buf[HEADER_LENGTH:])

	if dns.Header.Query == false {
		dns.Answer, buf = ParseAnswer(buf)
	}
	println(len(buf))
	return dns
}

func (dns *Dns) NewSimpleQuestion(domain string) *DnsPacket {
	return &DnsPacket{
		Header: &Header{
			ID: dns.cur_id,
			Query: true,
			OpCode: 0,
			Recursion: true,
			QDCOUNT: 1,
		},
		Question: &Question{
			QNAME: domain,
			QTYPE: 1,
			QCLASS: 1,
		},
		Answer: nil,
	}
}

func (packet *DnsPacket) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(packet.Header.Bytes())
	if packet.Question != nil {
		buf.Write(packet.Question.Bytes())
	}
	if packet.Answer != nil {
		buf.Write(packet.Answer.Bytes())
	}
	return buf.Bytes()
}

func (packet *DnsPacket) String() (str string) {
	str += packet.Header.String()
	if packet.Question != nil {
		str += packet.Question.String()
	}
	if packet.Answer != nil {
		str += packet.Answer.String()
	}
	return 
}

