package main

import (
	"flag"
	"github.com/yanrishbe/redis/main/server"
	"github.com/yanrishbe/redis/main/server/entities"
	"github.com/yanrishbe/redis/main/server/readWriteData"
	"log"
	"net"
	"regexp"
)

func main() {
	////////////////////////////////////HANDLING FLAGS//////////////////////////////////////////////////////////////////
	var (
		port        string
		mode        string
		defaultPort = "9090"
		defaultMode = "RAM"
		usagePort   = "the flag is used to choose on which port the server should listen to connections"
		usageMode   = "the flag is used to choose where to store the data the server is going to process"
	)

	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, "shorthand for --port")
	flag.StringVar(&mode, "mode", defaultMode, usageMode)
	flag.StringVar(&mode, "m", defaultMode, "shorthand for --mode")
	flag.Parse()

	pattern := "^(6553[0-5]|655[0-2]\\d|65[0-4](\\d){2}|6[0-4](\\d){3}|[1-5](\\d){4}|[1-9](\\d){0,3})$"
	match, err := regexp.MatchString(pattern, port)
	if !match || err != nil {
		log.Fatalln("Incorrect port info")
	}

	port = ":" + port
	//////////////////////////////LISTENING AND ACCEPTING CONNECTIONS///////////////////////////////////////////////////
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	log.Printf("Server is running on %s\n", port)
	log.Println("Ready to accept connections")

	commands := make(chan entities.Command)
	go readWriteData.AnswerToClient(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error accepting connection %+#v", err)
		}
		go server.HandleConnection(conn, commands)
	}
}
