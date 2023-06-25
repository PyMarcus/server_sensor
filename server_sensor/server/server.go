package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/fatih/color"
)

// Armazena as informações enviadas pelo cliente sensor do tipo chave, valor
var memory map[string]string = make(map[string]string)

// Estrutura para representar o servidor, do tipo pública
type Server struct {
	Ip   string
	Port string
}

// Definem o ip e a porta do servidor
func (s *Server) setIpAndPort(ip, port string) {
	s.Ip = ip
	s.Port = port
}

// Cria a conexão tcp com os clientes, chamando a goroutine para tratá-los
func (s Server) Service(ipAndPort ...string) {
	color.Set(color.FgYellow, color.BgBlack)
	log.Println("[SERVER]")
	if ipAndPort != nil {
		s.setIpAndPort(ipAndPort[0], ipAndPort[1])
	}
	log.Printf("[+]Servidor em execucao em %s:%s \n", s.Ip, s.Port)

	// socket em modo escuta
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Ip, s.Port))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[*]Aguardando conexoes...")

	for {
		// aceita conexoes com clientes
		conn, _ := listener.Accept()
		defer conn.Close()
		log.Println("[+]Conexao recebida de: ", conn.RemoteAddr())
		// inicia a goroutine para aceitar múltiplos clientes
		go s.handleConnection(conn)
	}

}

// Trata novas conexões
func (s *Server) handleConnection(conn net.Conn) {
	for {
		// le os dados recebidos
		buffer, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Printf("[-]ERROR: %v\n", err)
		}

		// mensagem recebida e tratada
		tbuffer := strings.TrimSpace(string(buffer))
		toSend, err := s.parse(tbuffer)

		if err != nil {
			// diz que foi criado
			_, err = conn.Write([]byte(toSend))
		} else {

			// salva retorna o valor solicitado
			if toSend == "" {
				toSend = "No content"
				conn.Write([]byte("No content"))

			} else {
				conn.Write([]byte(toSend))
			}
		}
		log.Println("[+]Resposta '", toSend, "' enviada para: ", conn.RemoteAddr(), "\n")
	}
}

/*
Trata a mensagem recebida.
Se for do tipo publicar,
salva no dicionario
*/
func (s Server) parse(message string) (toSend string, err error) {
	commands := strings.Split(message, ",")
	log.Println("[+]Tratando comando recebido: ", message)
	// se o tamanho da mensagem for de 2 palavraas, indica uma requisicao do cliente atuador
	if len(commands) == 2 {
		if strings.TrimSpace(commands[0]) == "ASSINAR" {
			return memory[strings.TrimSpace(commands[1])], nil
		}
		return "Falha! Comando inválido", errors.New("Comando inválido!")
	}
	// recebe mensagem do sensor para salvar nova informação
	if strings.TrimSpace(commands[0]) == "PUBLICAR" {
		memory[strings.TrimSpace(commands[1])] = strings.TrimSpace(commands[2])
	}

	return "created!", errors.New("Created!")
}
