package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// atuador requisita informações de luminosidade a cada 10s
type Actuator struct {
	ServerIp   string
	ServerPort string
}

// solicita, a cada 10s, informações de luminosidade ao servidor
func (a Actuator) Request() {
	log.Println("[ACTUADOR]")
	const MESSAGE string = "ASSINAR, LUMINOSIDADE"

	buffer := make([]byte, 2048)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", a.ServerIp, a.ServerPort))
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[+]Conectado ao servidor em: ", a.ServerIp, a.ServerPort)

	for {
		log.Println("[+]Solicitando informacoes de luminosidade...")

		_, err = conn.Write([]byte(MESSAGE + "\n"))
		if err != nil {
			log.Panicln(err)
		}
		_, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		log.Println("[+]Resposta do servidor: ", strings.TrimSpace(string(buffer)))
		log.Println("[+]Esperando 10s")
		time.Sleep(10 * time.Second)
	}
}

func main() {
	actuador := Actuator{ServerIp: "127.0.0.1", ServerPort: "9999"}
	actuador.Request()
}
