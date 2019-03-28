package main

import (
	"github.com/yanrishbe/redis/main/client/flagsClient"
	"github.com/yanrishbe/redis/main/client/getData"
	"github.com/yanrishbe/redis/main/client/readData"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	var (
		port        int
		host        string
		defaultPort = 9090
		defaultHost = "127.0.0.1"
		usagePort   = "the flag is used to choose to which port the client should connect to the server"
		usageHost   = "the flag is used to choose the host to connect to the server"
	)

	flagsClient.InitFlags(&port, defaultPort, &host, defaultHost, usagePort, usageHost)
	matchHost, errHost := flagsClient.ValidHost(host)

	errPort := flagsClient.ValidPort(port)
	addr := host + ":" + strconv.Itoa(port)

	if errPort != nil || !matchHost || errHost != nil {
		log.Fatalln("incorrect host info")
	}

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
		readData.InputClient(conn)
		getData.OutputToClient(conn, &os.File{})
	}
}
