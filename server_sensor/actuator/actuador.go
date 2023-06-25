package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// atuador requisita informações de luminosidade a cada 10s
type Actuator struct {
	ServerIp   string
	ServerPort string
}

// solicita, a cada 10s, informações de luminosidade ao servidor
func (a Actuator) Request() {
	const message string = "ASSINAR, LUMINOSIDADE"

	buffer := make([]byte, 2048)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", a.ServerIp, a.ServerPort))
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[+]Conectado ao servidor em: ", a.ServerIp, a.ServerPort)

	for {
		log.Println("[+]Esperando 10s")
		time.Sleep(10 * time.Second)
		_, err = conn.Write([]byte(message + "\n"))
		log.Println("[+]Solicitando informacoes de luminosidade...")
		if err != nil {
			log.Panicln(err)
		}
		response, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		log.Println("[+]Resposta do servidor: ", string(response))
	}
}

func main() {
	actuador := Actuator{ServerIp: "127.0.0.1", ServerPort: "9999"}
	actuador.Request()
}
