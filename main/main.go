package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"regexp"
	"sort"
	"strings"
)

type command struct {
	fields []string
	result chan string
}

type customError struct {
	custError string
}

func (e customError) Error() string {
	return fmt.Sprintf("%v", e.custError)
}

var data = make(map[string]string)


func handleConnection(conn net.Conn, commands chan command) {

	defer func() {
		//conn.Close()
		err:=conn.Close()
		if err!=nil{
			log.Fatalln("Error closing a connection")
		}
		log.Println("Connection closed")//не fatal потому что оборвутся другие коннекшены
	}()

	log.Println("Connection from", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)//возвращает интерфейс scanner

	for scanner.Scan() { //bufio.scan возвращает булевское значение
		ln := scanner.Text()//из байтов в текст
		fs := strings.Fields(ln)// парсит строку, пробелы пропускает

		for _,val:= range fs {
			if match:=strings.EqualFold("stop", val); match{
				return
			}
		}

		result := make(chan string)
		commands <- command{
			fields: fs,
			result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}
}

func storage(cmd chan command) {

	for cmd := range cmd {
		if len(cmd.fields) < 1{
			cmd.result <- "Please input a command you'd like to execute\n"
			continue
		}

		if len(cmd.fields) < 2 {
			cmd.result <- "Expected at least 2 arguments\n"
			continue
		}

		if len(cmd.fields) > 3 {
			cmd.result <- "Too many arguments"
			continue
		}

		fmt.Println("Command:", cmd.fields)


		// Executing commands
		switch strings.ToLower(cmd.fields[0]){

		// GET <KEY>
		case "get":
			_,state:=data[cmd.fields[1]]
			if !state {
				cmd.result <- "state:" + " " + "absent"
			} else {
				cmd.result <-  "value:" + " " + data[cmd.fields[1]] + "\n" + "state:" + " " + "present"
			}

			// SET <KEY> <VALUE>
		case "set":
			match,_:=regexp.MatchString("^[\\w]+$",cmd.fields[1])
			if !match {
				cmd.result <- "Incorrect key, please, try again"
				continue
			} else if len(cmd.fields) != 3 {
				cmd.result <- "Expected value"
				continue
			}
			data[cmd.fields[1]] = cmd.fields[2]
			cmd.result <- ""

			// DEL <KEY>
		case "del":
			_, state:=data[cmd.fields[1]]
			if !state{
				cmd.result <- "state:" + " " + "ignored"
			} else {
				delete(data, cmd.fields[1])
				cmd.result <- "state:" + " " + "absent"
			}

			// KEYS <PATTERN>
		case "keys":
			keys:=make([]string,0)
			keyString:=""

			if strings.Contains(cmd.fields[1],"*"){
				cmd.fields[1]=strings.TrimRight(cmd.fields[1],"*")

				for key := range data {
					if strings.HasPrefix(key,cmd.fields[1]){
						keys=append(keys, key)
					}
				}

				sort.Strings(keys)
				result:=strings.Join(keys,", ")

				if len(keys)==0{
					cmd.result <- "There are no keys matching the pattern"
				} else {
					cmd.result <- result
				}

			} else {

				for key := range data{
					if key == cmd.fields[1]{
						keyString +=key
					}
				}

				if len(keyString)==0{
					cmd.result <- "There are no keys matching the pattern"
				} else {
					cmd.result <- keyString
				}
			}

		default:
			cmd.result <- "Invalid command " + cmd.fields[0]
		}
	}
}

func main() {
	////////////////////////////////////HANDLING FLAGS//////////////////////////////////////////////////////////////////
	var (
		port string
		mode string
		defaultPort = "9090"
		defaultMode = "RAM"
		usagePort = "the flag is used to choose on which port the server should listen to connections"
		usageMode = "the flag is used to choose where to store the data the server is going to process"
	)

	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, "shorthand for --port")
	flag.StringVar(&mode, "mode", defaultMode, usageMode)
	flag.StringVar(&mode, "m", defaultMode,"shorthand for --mode")
	flag.Parse()

	match,err:=regexp.MatchString("^[\\d]+$",port)
	if !match || err!=nil {
		log.Fatalln(func ()error {
			return customError{"Incorrect port info"}
		}())
	}

	port=":"+port
	//////////////////////////////LISTENING AND ACCEPTING CONNECTIONS///////////////////////////////////////////////////
	li, err := net.Listen("tcp", port)//set up a server and return Listener interface
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	log.Printf("Server is running on %s\n", port)
	log.Println("Ready to accept connections")

	//The return value of net.Listen() is of the net.Conn type, which implements the io.Reader and io.Writer interfaces
	//net.Listen for telling a program to accept net.connections + act as a server

	commands := make(chan command)
	go storage(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Printf("Error accepting connection %+#v", err)
		}
		go handleConnection(conn, commands)//start a new goroutine each time it has to serve a TCP client
	}
}
