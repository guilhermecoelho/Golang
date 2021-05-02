package exemplo

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//import exemplo "go/exemplo"

func ChatSimple() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server %s", err.Error())
	}

	defer listener.Close()
	log.Printf("start server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection %s", err.Error())
		}

		go readMsg(conn)

	}
}

func readMsg(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		if strings.Compare(msg, "quit") == 1 {
			fmt.Printf("close connection")
			conn.Write([]byte("good bye" + "\n"))
			conn.Close()
			break
		}

		fmt.Printf(msg)
		conn.Write([]byte(msg + "\n"))
	}
}
