package getData

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func OutputToClient(conn net.Conn, out io.Writer) io.Writer {
	scannerConn := bufio.NewScanner(conn)
	for scannerConn.Scan() {
		_, errWrite :=out.Write([]byte("The server sends: " + scannerConn.Text()))
		if errWrite != nil {
			log.Println("Writing to file error")
		}
		log.Println("The server sends: " + scannerConn.Text())
		break
	}
	if errWriteConn := scannerConn.Err(); errWriteConn != nil {
		log.Printf("Writing to stdout error: %T %+v", errWriteConn, errWriteConn)
		os.Exit(1)
	}
	fmt.Println("---")
	return out
}
