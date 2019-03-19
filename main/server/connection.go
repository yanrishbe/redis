package server

import (
	"bufio"
	"github.com/yanrishbe/redis/main/server/entities"
	"io"
	"log"
	"net"
	"strings"
)

func HandleConnection(conn net.Conn, commands chan entities.Command) {

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln("Error closing a connection")
		}
		log.Println("Connection is closed")
	}()

	log.Println("Connection from", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn) //returns scanner interface

	for scanner.Scan() { //bufio.scan returns bool value
		ln := scanner.Text()
		fs := strings.Fields(ln)

		for _, val := range fs {
			if match := strings.EqualFold("stop", val); match {
				return
			}
		}

		result := make(chan string)
		commands <- entities.Command{
			Fields: fs,
			Result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}
}
