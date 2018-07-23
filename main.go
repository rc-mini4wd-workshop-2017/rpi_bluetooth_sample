package main

import (
	"github.com/tarm/serial"
	"log"
	"os"
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
	_, err = s.Write([]byte("info\n"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	message := ""
	for {
		time.Sleep(time.Millisecond * 500)

		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		message = message + string(buf[:n])

		if strings.Contains(message, "result") {
			log.Print(message)
			os.Exit(0)
		}
	}
	os.Exit(1)
}
