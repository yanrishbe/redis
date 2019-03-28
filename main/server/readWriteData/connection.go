package readWriteData

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

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
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

		_, err := io.WriteString(conn, <-result+"\n")
		if err != nil {
			log.Println("Error writing data")
			return
		}
	}
}
