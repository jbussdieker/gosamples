package dns

import "bytes"
import "strings"
import "fmt"
import "encoding/binary"

type Question struct {
	QNAME string
	QTYPE RecordType
	QCLASS uint16
}

func ParseQuestion(buffer *bytes.Buffer) (q *Question) {
	q = &Question{}
	q.QNAME = readString(buffer)
	binary.Read(buffer, binary.BigEndian, &q.QTYPE)
	binary.Read(buffer, binary.BigEndian, &q.QCLASS)
	return q
}

func (q *Question) Bytes() []byte {
	buf := new(bytes.Buffer)
	parts := strings.Split(q.QNAME, ".")
	for _, part := range parts {
		write8(buf, uint8(len(part)))
		buf.Write([]byte(part))
	}
	write8(buf, 0)
	write16(buf, uint16(q.QTYPE))
	write16(buf, q.QCLASS)
	return buf.Bytes()
}

func (q *Question) String() (str string) {
	str += fmt.Sprintf("             QNAME: %s\n", q.QNAME)
	str += fmt.Sprintf("             QTYPE: %d\n", q.QTYPE)
	str += fmt.Sprintf("            QCLASS: %d\n", q.QCLASS)
	return
}
