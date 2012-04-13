package dns

////////////////////////////////////////////////////////////////////////////////
// Public functions
////////////////////////////////////////////////////////////////////////////////

func (dns *Connection) NewQuestion(rtype RecordType, domain string) *Message {
	return &Message{
		Header: &Header{
			ID: dns.cur_id,
			Query: true,
			OpCode: DNS_OPCODE_QUERY,
			Recursion: true,
			QuestionCount: 1,
		},
		Questions: []*Question{
			NewQuestion(domain, rtype, DNS_CLASS_IN),
		},
	}
}

