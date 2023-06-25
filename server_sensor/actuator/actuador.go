package actuator

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// atuador requisita informações de luminosidade a cada 10s
type Actuator struct {
	ServerIp   string
	ServerPort string
}

// solicita, a cada 10s, informações de luminosidade ao servidor
func (a Actuator) Request() {
	color.Set(color.FgBlue, color.BgBlack)

	log.Println("[ACTUADOR]")
	const MESSAGE string = "ASSINAR, LUMINOSIDADE"

	buffer := make([]byte, 2048)

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", a.ServerIp, a.ServerPort))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[+]Conectado ao servidor em: ", a.ServerIp, a.ServerPort)

	reader(conn, MESSAGE, buffer)
}

func reader(conn net.Conn, MESSAGE string, buffer []byte) {
	for {
		log.Println("[+]Solicitando informacoes de luminosidade...")

		_, err := conn.Write([]byte(MESSAGE + "\n"))
		if err != nil {
			log.Panicln(err)
		}
		_, err = conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		log.Println("[+]Resposta do servidor: ", strings.TrimSpace(strings.Split(string(buffer), "\n")[0]))
		decision(strings.TrimSpace(strings.Split(string(buffer), "\x00")[0]))
		log.Println("[+]Esperando 10s\n")
		time.Sleep(10 * time.Second)
	}
}

func decision(number string) {
	if num, _ := strconv.Atoi(number); num >= 60 {
		log.Println("[+]Abrindo a cortina")
	} else {
		log.Println("[*]Fechando a cortina")
	}
}
