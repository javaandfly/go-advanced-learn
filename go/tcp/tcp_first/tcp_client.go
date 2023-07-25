package main

import (
	"fmt"
	"net"
)

// tcp/client/main.go

// 客户端
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	// inputReader := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		msg := `1`
		conn.Write([]byte(msg))
		// time.Sleep(time.Second * 1)
	}
}
