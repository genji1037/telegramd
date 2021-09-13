package test

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
	"time"
)

const ()

var RequestForPQ = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x4A, 0x96, 0x70, 0x27, 0xC4, 0x7A, 0xE5, 0x51,
	0x14, 0x00, 0x00, 0x00, 0x78, 0x97, 0x46, 0x60,
	0x3E, 0x05, 0x49, 0x82, 0x8C, 0xCA, 0x27, 0xE9,
	0x66, 0xB3, 0x01, 0xA4, 0x8F, 0xEC, 0xE2, 0xFC}

func TestHandshakeTCP(t *testing.T) {
	address := "127.0.0.1:8800"
	conn, err := EstablishTCPConn(address)
	if err != nil {
		t.Fatalf("EstablishTCPConn %s failed: %v", address, err)
	}
	_, err = conn.Write(MTProtoTransportsAbridged(RequestForPQ))
	if err != nil {
		t.Fatalf("write conn failed: %v", err)
	}
	select {}
}

// refer https://core.telegram.org/mtproto/mtproto-transports
func MTProtoTransportsAbridged(payload []byte) []byte {
	length := make([]byte, 4)
	binary.LittleEndian.PutUint32(length, uint32(len(payload)/4))

	res := make([]byte, len(payload)+2)
	res[0] = 0xef
	res[1] = length[0]
	copy(res[2:], payload)
	return res
}

func EstablishTCPConn(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("dail %s failed: %v", address, err)
	}
	buf := make([]byte, 10240)
	go func() {
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("read TCP conn failed: %v\n", err)
				continue
			}
			fmt.Printf("%x\n", buf[:n])
		}
	}()
	return conn, nil
}

// refer https://core.telegram.org/mtproto/samples-auth_key
func ReqPQ() []byte {
	data := make([]byte, 40)
	// auth key is 0
	binary.LittleEndian.PutUint64(data[8:], uint64(time.Now().Unix()*(2<<32))) // message_id: Exact unixtime * 2^32
	binary.LittleEndian.PutUint32(data[16:], uint32(20))                       // Message body length
	copy(data[20:], []byte{0x78, 0x97, 0x46, 0x60})                            // req_pq constructor number from TL schema
	//todo
	return data
}
