package checksnmp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Checksnmp(ip string, wg *sync.WaitGroup) {
	defer wg.Done()
	p := make([]byte, 2048)
	to_sent := []byte{0x30, 0x37, 0x02, 0x01, 0x01, 0x04, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0xa0, 0x2a, 0x02, 0x04, 0x60, 0x32, 0x86, 0x48, 0x02, 0x01, 0x00, 0x02, 0x01, 0x00, 0x30, 0x1c, 0x30, 0x0c, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x04, 0x00, 0x05, 0x00, 0x30, 0x0c, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x07, 0x00, 0x05, 0x00}
	conn, err := net.Dial("udp", ip+":161")
	conn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, string(to_sent))

	_, err = bufio.NewReader(conn).Read(p)
	if p[3] != 0 {
		log.Println("[*] UDP 161 can access the internet")
		return
	}
	log.Println("[-] UDP 161 May blocked")
	conn.Close()
	return
}
