package main

import (
	"github.com/yanrishbe/redis/main/server/entities"
	"github.com/yanrishbe/redis/main/server/flagsServer"
	"github.com/yanrishbe/redis/main/server/readWriteData"
	"log"
	"net"
	"strconv"
)

func main() {
	////////////////////////////////////HANDLING FLAGS//////////////////////////////////////////////////////////////////
	var (
		port        int
		mode        string
		defaultPort = 9090
		defaultMode = "RAM"
		usagePort   = "the flag is used to choose on which port the server should listen to connections"
		usageMode   = "the flag is used to choose where to store the data the server is going to process"
	)

	flagsServer.InitFlags(&port, defaultPort, &mode, defaultMode, usagePort, usageMode)

	errPort := flagsServer.ValidPort(port)
	addr := ":" + strconv.Itoa(port)

	if errPort != nil {
		log.Fatalln("incorrect host info")
	}

	//////////////////////////////LISTENING AND ACCEPTING CONNECTIONS///////////////////////////////////////////////////
	li, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(){
		err:=li.Close()
		if err!= nil {
			panic(err)
		}
	}()

	log.Printf("Server is running on %s\n", port)
	log.Println("Ready to accept connections")

	commands := make(chan entities.Command)
	go readWriteData.AnswerToClient(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error accepting connection %+#v", err)
		}
		go readWriteData.HandleConnection(conn, commands)
	}
}
