package main

import (
	"net"
	"strings"
)
const (
	// конец сообщения
	END1 = "12345678"
)
// соединения 
var conections = make(map[net.Conn]bool) 

func main() {
	// стартуем сервер
	serv, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// заакрываем соединение
    defer serv.Close()
    // открываем соединенине в цикле
	for {
		con, err :=  serv.Accept()
		if err != nil {
			break
		}
		go Hundle(con)
	}
} 

func Hundle(conn net.Conn) {
	conections[conn] = true
	var (
		massage string
		//буфер
		buffer = make([]byte, 256)
	)
	close: for {
    for {
		// обнуляем строчку
		massage = ""
	len, err := conn.Read(buffer)
		if err != nil  || len == 0 { 
			break close
			panic(err)
		}
		//читаем буфер до конца	
		massage = string(buffer[:len])
		if  strings.HasSuffix(massage, END1) {
			massage = strings.TrimSuffix(massage,END1)
			break
		}
	}
	// перебираем соединения 
	for c := range conections {
		if c == conn { continue }
		c.Write([]byte(strings.ToUpper(massage) + END1))
	}
}
delete(conections,conn)
}
