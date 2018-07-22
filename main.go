package main

import (
	"github.com/tarm/serial"
	"log"
	"strings"
	"time"
)

func main() {
	c := &serial.Config{Name: "/dev/rfcomm0", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("> info")
	n, err := s.Write([]byte("info\n"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	for {
		time.Sleep(time.Millisecond * 500)

		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		sbuf := string(buf[:n])
		log.Printf(sbuf)

		if strings.Contains(sbuf, "result") {
			break
		}
	}
}
