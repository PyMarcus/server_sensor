package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const MESSAGE string = "PUBLICAR, LUMINOSIDADE, 50"

type Sensor struct {
	ServerIp   string
	ServerPort string
}

// gera um valor aleatório, de 0 a 100, simulando a detecção de luminosidade do ambiente
func detectLuminosity() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(101)
}

// envia para o servidor a luminosidade detectada, a cada 5s
func (s Sensor) Post() {
	log.Println("[SENSOR]")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.ServerIp, s.ServerPort))
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[+]Conectado ao servidor em ", s.ServerIp, ":", s.ServerPort)

	for {
		log.Println("[+]Detectando luminosidade...")
		lumen := detectLuminosity()
		log.Println("[+]Detectado o valor de ", lumen, "lm.")
		_, err := conn.Write([]byte(MESSAGE + "\n"))

		if err != nil {
			log.Panicln(err)
		}
		log.Println("[+]Valor enviado ao servidor")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	sensor := Sensor{ServerIp: "127.0.0.1", ServerPort: "9999"}
	sensor.Post()
}
