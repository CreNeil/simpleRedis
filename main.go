package main

import (
	"fmt"
	"net"
	"time"
)

/**

send()
receive()
*/

type SimpleConn struct {
	TCPConn *net.TCPConn
	Buff    []byte
}

func getBytes(str string) []byte {
	return []byte(str)
}

func assembly(args []string) string {
	var res string
	for _, val := range args {
		res = res + val + "\r\n"
	}
	return res
}

func (sc *SimpleConn) Send(command string, key string, args ...string) (int, error) {
	sendStr := []string{command, key}
	for _, val := range args {
		sendStr = append(sendStr, val)
	}
	return sc.TCPConn.Write(getBytes(assembly(sendStr)))
}

func (sc *SimpleConn) Receive() {
	for {
		n, err := sc.TCPConn.Read(sc.Buff)
		if err != nil {
			fmt.Println("connection")
		}
		fmt.Println(" server receive " + string(sc.Buff[:n]))
	}
}

func main() {
	//dialer := &net.Dialer{
	//	Timeout:   time.Second * 30,
	//	KeepAlive: time.Minute * 5,
	//}
	//conn, err := dialer.Dial("tcp","localhost:6379")

	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:6379")

	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	simpleConn := &SimpleConn{
		TCPConn: tcpConn,
		Buff:    make([]byte, 128),
	}
	go simpleConn.Receive()

	if err != nil {
		return
	}
	//command := "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"
	//command = "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n"
	_, _ = simpleConn.Send("*3\r\n$3\r\nSET", "$5\r\nmykey", "$7\r\nmyvalue")
	_, _ = simpleConn.Send("*2\r\n$3\r\nGET", "$5\r\nmykey")

	time.Sleep(10000)

}
