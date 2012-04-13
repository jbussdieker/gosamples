package dns

import "bytes"
import "net"
import "fmt"
//import "os"

type Error error
//type Error os.Error

type Dns struct {
	server string
	cur_id uint16
	net.Conn
}

type DnsPacket struct {
	*Header
	Questions[] *Question
	Answers[] *Answer
	Nameservers[] *Answer
	Additionals[] *Answer
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

func (dns *Dns) Send(packet *DnsPacket) (resp *DnsPacket, err Error) {
	conn, err := net.Dial("udp", dns.server)
	dns.Conn = conn
	if err != nil {
		return
	}

	_, err = dns.Write(packet.Bytes())
	if err != nil {
		return
	}

	buf := make([]byte, 2000)
	s, err := dns.Read(buf)
	if err != nil {
		return
	}
	buf = buf[0:s]
	resp = ParsePacket(buf)
	return
}

func ParsePacket(buf []byte) *DnsPacket {
	buffer := bytes.NewBuffer(buf)

	dns := &DnsPacket{
		Header: ParseHeader(buffer),
	}
	dns.Questions = make([]*Question, dns.Header.QuestionCount)
	for i := 0; i < int(dns.Header.QuestionCount); i++ {
		dns.Questions[i] = ParseQuestion(buffer)
	}
	dns.Answers = make([]*Answer, dns.Header.AnswerCount)
	for i := 0; i < int(dns.Header.AnswerCount); i++ {
		dns.Answers[i] = ParseAnswer(buffer)
	}
	dns.Nameservers = make([]*Answer, dns.Header.NameserverCount)
	for i := 0; i < int(dns.Header.NameserverCount); i++ {
		dns.Nameservers[i] = ParseAnswer(buffer)
	}
	dns.Additionals = make([]*Answer, dns.Header.AdditionalCount)
	for i := 0; i < int(dns.Header.AdditionalCount); i++ {
		dns.Additionals[i] = ParseAnswer(buffer)
	}
	if buffer.Len() > 0 {
		println("ERROR UNPARSED BYTES:", buffer.Len())
	}
	return dns
}

func (dns *Dns) NewQuestion(rtype RecordType, domain string) *DnsPacket {
	return &DnsPacket{
		Header: &Header{
			ID: dns.cur_id,
			Query: true,
			OpCode: DNS_OPCODE_QUERY,
			Recursion: true,
			QuestionCount: 1,
		},
		Questions: []*Question{
			{
				QNAME: domain,
				QTYPE: rtype,
				QCLASS: DNS_CLASS_IN,
			},
		},
	}
}

func (packet *DnsPacket) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.Write(packet.Header.Bytes())
	for i := 0; i < int(packet.Header.QuestionCount); i++ {
		buf.Write(packet.Questions[i].Bytes())
	}
	for i := 0; i < int(packet.Header.AnswerCount); i++ {
		buf.Write(packet.Answers[i].Bytes())
	}
	for i := 0; i < int(packet.Header.NameserverCount); i++ {
		buf.Write(packet.Nameservers[i].Bytes())
	}
	for i := 0; i < int(packet.Header.AdditionalCount); i++ {
		buf.Write(packet.Additionals[i].Bytes())
	}
	return buf.Bytes()
}

func (packet *DnsPacket) String() (str string) {
	if packet.Header.Query {
		str += "========== DNS Query ==========\n"
	} else {
		str += "========= DNS Response ========\n"
	}
	str += packet.Header.String()
	if int(packet.Header.QuestionCount) > 0 { str += "======= Questions =======\n" }
	for i := 0; i < int(packet.Header.QuestionCount); i++ {
		str += packet.Questions[i].String()
	}
	if int(packet.Header.AnswerCount) > 0 { str += "======== Answers ========\n" }
	for i := 0; i < int(packet.Header.AnswerCount); i++ {
		str += packet.Answers[i].String()
	}
	if int(packet.Header.NameserverCount) > 0 { str += "====== Nameservers ======\n" }
	for i := 0; i < int(packet.Header.NameserverCount); i++ {
		str += packet.Nameservers[i].String()
	}
	if int(packet.Header.AdditionalCount) > 0 { str += "====== Additionals ======\n" }
	for i := 0; i < int(packet.Header.AdditionalCount); i++ {
		str += packet.Additionals[i].String()
	}
	return 
}

