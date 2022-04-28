package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		buff, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Println(buff)

		cmd_args := strings.Split(string(buff)[:len(buff)-1], " ")
		cmd := exec.Command(cmd_args[0], cmd_args[1:]...)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Fprintf(c, string(stdout))
	}
}
