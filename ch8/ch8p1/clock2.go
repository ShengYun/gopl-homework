// Exercise 8.1: Modify clock2 to accept a port number, and write a program,
// clockwall, that acts as a client of several clock servers at once,
// reading the times from each one and displaying the results in a table,
// akin to the wall of clocks seen in some business offices. If you have access
// to geographically distributed computers, run instances remotely;
// otherwise run local instances on different ports with fake time zones.

// Click here to view code image

// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var portVar int

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.IntVar(&portVar, "port", 8000, "port number of localhost")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(portVar))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
