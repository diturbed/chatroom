package main

import (
	"fmt"
	"net"
)

var usermap = make(map[string]*net.TCPConn)

func main() {
	tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	listener, _ := net.ListenTCP("tcp", tcpaddr)
	for {
		conn, _ := listener.AcceptTCP()
		go func(conn *net.TCPConn) {
			from := ""
			for {
				buf := make([]byte, 1024)
				cnt, _ := conn.Read(buf)
				from = string(buf[:cnt])
				usermap[from] = conn
				_, _ = conn.Write([]byte("y"))
				if cnt > 0 {
					break
				}
			}
			to := ""
			for {
				buf := make([]byte, 1024)
				cnt, _ := conn.Read(buf)
				to = string(buf[:cnt])
				_, _ = conn.Write([]byte("y"))
				if cnt > 0 {
					break
				}
			}
			fmt.Println(usermap)

			for {
				buf := make([]byte, 1024)
				cnt, _ := conn.Read(buf)
				// fmt.Println(string(buf[:cnt]), " from ", from, " to ", to)
				go func() {
					tmp := []byte("from" + to)
					tmp = append(tmp, buf[:cnt]...)
					if sendconn, ok := usermap[to]; ok {
						_, _ = sendconn.Write(tmp)
					}
				}()
			}
		}(conn)
	}
}
