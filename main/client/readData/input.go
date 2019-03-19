package readData

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func InputClient(conn net.Conn) {

	scannerStdin := bufio.NewScanner(os.Stdin)
	fmt.Print("Command to send: ")
	for scannerStdin.Scan() {
		text := scannerStdin.Text()
		if errReadConn := scannerStdin.Err(); errReadConn != nil {
			log.Printf("Reading error: %T %+v", errReadConn, errReadConn)
			os.Exit(1)
		}
		if match := strings.EqualFold("stop", text); match {
			log.Println("Disconnecting from the server...")
			os.Exit(1)
		}
		fmt.Println("---")
		// send to server
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			log.Fatalln("The server is offline, try to reconnect")
		}
		log.Println("The server receives: " + text)
		break
	}
}