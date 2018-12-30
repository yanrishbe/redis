package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

type customClientError struct {
	custClientError string
}

func (e customClientError) Error() string {
	return fmt.Sprintf("%v", e.custClientError)
}

func main() {

	var (
		port        string
		host        string
		defaultPort = "9090"
		defaultHost = "127.0.0.1"
		usagePort   = "the flag is used to choose to which port the client should connect to the server"
		usageHost   = "the flag is used to choose the host to connect to the server"
	)

	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, "shorthand for --port")
	flag.StringVar(&host, "host", defaultHost, usageHost)
	flag.StringVar(&host, "h", defaultHost, "shorthand for --host")
	flag.Parse()

	patternPort := "^(6553[0-5]|655[0-2]\\d|65[0-4](\\d){2}|6[0-4](\\d){3}|[1-5](\\d){4}|[1-9](\\d){0,3})$"
	matchPort, errPort := regexp.MatchString(patternPort, port)
	if !matchPort || errPort != nil {
		log.Fatalln(func() error {
			return customClientError{"Incorrect port info"}
		}())
	}
	patternHost := "^((\\d|[1-9]\\d|1(\\d){2}|2[0-4]\\d|25[0-5])\\.){3}(\\d|[1-9]\\d|1[(\\d){2}|2[0-4]\\d|25[0-5])$"
	matchHost, errHost := regexp.MatchString(patternHost, host)
	if !matchHost || errHost != nil {
		log.Fatalln(func() error {
			return customClientError{"Incorrect host info"}
		}())
	}

	port = ":" + port
	addr := host + port

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to establish connection, dialling error %+#v", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln("Error closing a connection")
		}
		log.Println("Connection is closed")
	}()

	for {
		// read in input from stdin
		scannerStdin := bufio.NewScanner(os.Stdin)
		fmt.Print("Command to send: ")
		for scannerStdin.Scan() {
			text := scannerStdin.Text()
			if match := strings.EqualFold("stop", text); match {
				log.Println("Disconnecting from the server...")
				return
			}
			fmt.Println("---")
			// send to server
			//_, errWrite := fmt.Fprintf(conn, text+"\n")///////?????which better????/////////////////////////////////
			_, errWrite := conn.Write([]byte(text+"\n"))
			if errWrite != nil {
				log.Fatalln("The server is offline, try to reconnect")
			}
			log.Println("The server receives: " + text)
			break
		}
		// listen for reply
		scannerConn := bufio.NewScanner(conn)
		for scannerConn.Scan() {
			log.Println("The server sends: " + scannerConn.Text())
			break
		}
		if errReadConn := scannerStdin.Err(); errReadConn != nil {
			log.Printf("Reading error: %T %+v", errReadConn, errReadConn)
			return
		}
		fmt.Println("---")
	}
}
