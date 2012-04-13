package dns

import "bytes"
import "encoding/binary"
import "fmt"

type Header struct {
	ID uint16 // WORD: Message ID
	Query bool // BIT 7: 0:Query, 1:Response
	OpCode uint8 // BIT 6-3: 1: Standard Query
	Authoritative bool // BIT 2: Authoratative answer 1:true 0:false
	Truncated bool // BIT 1: 1:message truncated, 0:normal
	Recursion bool // BIT 0: 1:request recursion, 0:no recursion
	RecursionSupported bool // BIT 7: 1:recursion supported, 0:not
	Reserved uint8 // BIT 6-4: Reserved 0
	ResponseCode uint8 // BIT 3-0: Response code
	QDCOUNT uint16 // WORD: Number of queries
	ANCOUNT uint16 // WORD:
	NSCOUNT uint16 // WORD:
	ARCOUNT uint16 // WORD:
}

func NewHeader() *Header {
	return &Header{}
}

func ParseHeader(buffer *bytes.Buffer) *Header {
	h := &Header{}
	binary.Read(buffer, binary.BigEndian, &h.ID)

	var byte3 uint8
	binary.Read(buffer, binary.BigEndian, &byte3)
	h.parse_byte3(byte3)

	var byte4 uint8
	binary.Read(buffer, binary.BigEndian, &byte4)
	h.parse_byte4(byte4)

	binary.Read(buffer, binary.BigEndian, &h.QDCOUNT)
	binary.Read(buffer, binary.BigEndian, &h.ANCOUNT)
	binary.Read(buffer, binary.BigEndian, &h.NSCOUNT)
	binary.Read(buffer, binary.BigEndian, &h.ARCOUNT)
	return h
}

func write16(buf *bytes.Buffer, value uint16) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic("write16 failed")
	}
}

func write8(buf *bytes.Buffer, value uint8) {
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		panic("write16 failed")
	}
}

func (h *Header) parse_byte3(byte3 byte) {
	if (byte3 & 0x80) == 0 {
		h.Query = true
	}
	h.OpCode = (byte3 >> 3) & 0xF
	if (byte3 >> 2) & 1 == 1 {
		h.Authoritative = true
	}
	if (byte3 >> 1) & 1 == 1 {
		h.Truncated = true
	}
	if byte3 & 1 == 1 {
		h.Recursion = true
	}
}

func (h *Header) byte3() byte {
	var byte3 uint8
	if !h.Query {
		byte3 |= 0x80
	}
	byte3 |= (h.OpCode & 0xF) << 3
	if h.Authoritative {
		byte3 |= 0x04
	}
	if h.Truncated {
		byte3 |= 0x02
	}
	if h.Recursion {
		byte3 |= 0x01
	}
	return byte3
}

func (h *Header) parse_byte4(byte4 byte) {
	if (byte4 & 0x80) != 0 {
		h.RecursionSupported = true
	}
	h.Reserved = (byte4 >> 4) & 0x7
	h.ResponseCode = byte4 & 0x0F
}

func (h *Header) byte4() byte {
	var byte4 uint8
	if h.RecursionSupported {
		byte4 |= 0x80
	}
	byte4 |= (h.Reserved & 0x7) << 4 
	byte4 |= (h.ResponseCode & 0x0F)
	return byte4
}

func (h *Header) Bytes() []byte {
	buf := new(bytes.Buffer)
	write16(buf, h.ID)
	write8(buf, h.byte3())
	write8(buf, h.byte4())
	write16(buf, h.QDCOUNT)
	write16(buf, h.ANCOUNT)
	write16(buf, h.NSCOUNT)
	write16(buf, h.ARCOUNT)
	return buf.Bytes()
}

func (h *Header) String() (str string) {
	str += fmt.Sprintf("                ID: %d\n", h.ID)
	str += fmt.Sprintf("             Query: %v\n", h.Query)
	str += fmt.Sprintf("            OpCode: %d\n", h.OpCode)
	str += fmt.Sprintf("     Authoritative: %v\n", h.Authoritative)
	str += fmt.Sprintf("         Truncated: %v\n", h.Truncated)
	str += fmt.Sprintf("         Recursion: %v\n", h.Recursion)
	str += fmt.Sprintf("RecursionSupported: %v\n", h.RecursionSupported)
	str += fmt.Sprintf("          Reserved: %v\n", h.Reserved)
	str += fmt.Sprintf("      ResponseCode: %v\n", h.ResponseCode)
	str += fmt.Sprintf("           QDCOUNT: %v\n", h.QDCOUNT)
	str += fmt.Sprintf("           ANCOUNT: %v\n", h.ANCOUNT)
	str += fmt.Sprintf("           NSCOUNT: %v\n", h.NSCOUNT)
	str += fmt.Sprintf("           ARCOUNT: %v\n", h.ARCOUNT)
	return
}
