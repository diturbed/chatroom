package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var from string
var to string

func main() {
	fmt.Println("your nick name:")
	fmt.Scanf("%s\n", &from)
	fmt.Println("talk to:")
	fmt.Scanf("%s\n", &to)

	buf := make([]byte, 1024)
	conn, _ := net.Dial("tcp", "127.0.0.1:8888")
	for {
		_, _ = conn.Write([]byte(from))
		cnt, _ := conn.Read(buf)
		if string(buf[:cnt]) == "y" {
			break
		}
	}
	for {
		_, _ = conn.Write([]byte(to))
		cnt, _ := conn.Read(buf)
		if string(buf[:cnt]) == "y" {
			break
		}
	}
	for {
		go func() {
			r := bufio.NewReader(os.Stdin)
			for {
				msg, _, _ := r.ReadLine()
				_, _ = conn.Write(msg)
			}
			// io.Copy(conn, os.Stdin)
		}()
		cnt, _ := conn.Read(buf)
		fmt.Println(string(buf[:cnt]))
	}
}
