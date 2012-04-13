package dns

import "bytes"
import "strings"
import "encoding/binary"

////////////////////////////////////////////////////////////////////////////////
// Types
////////////////////////////////////////////////////////////////////////////////

type ClassType uint16

const (
	DNS_CLASS_IN ClassType = 1
)

type RecordType uint16

const (
	DNS_RECORD_TYPE_A  RecordType = 1
	DNS_RECORD_TYPE_NS            = iota
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

type Question struct {
	Name  string
	Type  RecordType
	Class ClassType
}

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func ParseQuestion(buffer *bytes.Buffer) (q *Question) {
	q = &Question{}
	q.Name = readDnsString(buffer)
	binary.Read(buffer, binary.BigEndian, &q.Type)
	binary.Read(buffer, binary.BigEndian, &q.Class)
	return q
}

////////////////////////////////////////////////////////////////////////////////
// Method functions
////////////////////////////////////////////////////////////////////////////////

func (q *Question) Bytes() []byte {
	buf := new(bytes.Buffer)
	parts := strings.Split(q.Name, ".")
	for _, part := range parts {
		write8(buf, uint8(len(part)))
		buf.Write([]byte(part))
	}
	write8(buf, 0)
	write16(buf, uint16(q.Type))
	write16(buf, uint16(q.Class))
	return buf.Bytes()
}
