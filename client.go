package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	// конец сообщения
	END = "123456789"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	defer conn.Close()
	conn.Write([]byte(InputString() + END))

	go ClientWriter(conn)
	ClientReader(conn)
}

// функция ввода строки
func InputString() string {
	read, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(read, "\n", "", -1)
}

// ввод сообщения от клиента
func ClientWriter(conn net.Conn) {
	for {
		conn.Write([]byte(InputString() + END))
	}
}

// функция вывода
func ClientReader(conn net.Conn) {
	var (
		massage string
		// буфер
		buffer = make([]byte, 512)
	)
end:
	for {
		massage = ""
		// цикл чтения
		for {
			len, err := conn.Read(buffer)
			if err != nil || len == 0 {
				break end
				panic(err)
			}
			// читаем до конца строоки
			massage = string(buffer[:len])
			if strings.HasSuffix(massage, END) {
				massage = strings.TrimSuffix(massage, END)
				break
			}
		}
		fmt.Print(massage, "\n")
	}
}
