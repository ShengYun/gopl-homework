// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var cityAddrMap = make(map[string]string)

var supportedCities = map[string]bool{
	"NewYork": true,
	"Tokyo":   true,
	"London":  true,
}

func getTime(cityName string) {
	if _, ok := cityAddrMap[cityName]; !ok {
		errorMsg := fmt.Sprintf("clockwall: %s does not exist in map", cityName)
		log.Println(errors.New(errorMsg))
		return
	}

	addr := cityAddrMap[cityName]
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	copyTimeAndCityName(os.Stdout, conn, cityName)
}

func copyTimeAndCityName(dst io.Writer, src io.Reader, cityName string) {

	// need to read more about scanner
	s := bufio.NewScanner(src)
	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(dst, "%s: %s\n", cityName, s.Text())
	}
}

func populateMap(input string) {
	s := strings.Split(input, "=")
	if _, ok := supportedCities[s[0]]; ok {
		cityAddrMap[s[0]] = s[1]
	}
}

func main() {
	done := make(chan struct{})

	argsWithoutProg := os.Args[1:]
	// if len(argsWithoutProg) != 3 {
	// 	log.Fatal(errors.New("clockwall: Not enough command line arguments"))
	// }

	for _, input := range argsWithoutProg {
		populateMap(input)
	}

	for cityName := range cityAddrMap {
		go getTime(cityName)
	}

	<-done
	// for i := 0; i < len(argsWithoutProg); i++ {
	// 	go getTime(argsWithoutProg[i])
	// }

}
