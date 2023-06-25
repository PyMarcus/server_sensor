package sensor

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/fatih/color"
)

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
	color.Set(color.FgGreen, color.BgBlack)

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
		MESSAGE := fmt.Sprintf("PUBLICAR, LUMINOSIDADE, %s", strconv.Itoa(lumen))
		log.Println("ENVIANDO: ", MESSAGE)
		_, err := conn.Write([]byte(MESSAGE + "\n"))

		if err != nil {
			log.Panicln(err)
		}
		log.Println("[+]Valor enviado ao servidor\n")
		time.Sleep(5 * time.Second)
	}
}
