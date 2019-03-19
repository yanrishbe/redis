package getData

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func OutputToClient(conn net.Conn, scannerStdin *bufio.Scanner){
	scannerConn := bufio.NewScanner(conn)
	for scannerConn.Scan() {
		log.Println("The server sends: " + scannerConn.Text())
		break
	}
	if errReadConn := scannerStdin.Err(); errReadConn != nil {
		log.Printf("Reading error: %T %+v", errReadConn, errReadConn)
		os.Exit(1)
	}
	fmt.Println("---")
}
