package dns

import "testing"
import "fmt"

func testNewDns(t *testing.T) (dns *Dns) {
	dns = NewDns("localhost", 53)
	if dns == nil {
		t.Fail()
	}
	return
}

func TestNewDns(t *testing.T) {
	testNewDns(t)
}

func TestNewSimpleQuery(t *testing.T) {
	expected := []byte{0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 6, 109, 108, 111, 99, 97, 108, 4, 106, 111, 115, 104, 3, 99, 111, 109, 0, 0 ,1, 0, 1}
	dns := NewDns("localhost", 53)
	packet := dns.NewQuestion(DNS_RECORD_TYPE_A, "mlocal.josh.com")
	if string(packet.Bytes()) != string(expected) {
		t.Error("Got:     ", packet.Bytes())
		t.Error("Expected:", expected)
		t.Fail()
	}
}
/*
func TestNewSimpleQueryIRL(t *testing.T) {
	dns := NewDns("8.8.8.8", 53)
	packet := dns.NewSimpleQuestion("mague.com")
	println(packet.String())
	resp, err := dns.Send(packet)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	println(resp)
	resp_packet := ParsePacket(resp)
	println(resp_packet.String())
}
*/
func TestNewTextQueryIRL(t *testing.T) {
	dns := NewDns("8.8.8.8", 53)
	packet := dns.NewQuestion(DNS_RECORD_TYPE_TXT, "mlocal.josh.com")
	println(packet.String())
	resp, err := dns.Send(packet)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(resp)
	h := ParseHeader(resp)
	println(h.String())
	//resp_packet := ParsePacket(resp)
	//println(resp_packet.String())
}

