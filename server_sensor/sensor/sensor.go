package main

import "log"

// estrutura para representar o servidor, do tipo pública
type Server struct {
	Ip   string
	Port string
}

// definem o ip e a porta do servidor
func (s *Server) SetIpAndPort(ip, port string) {
	s.Ip = ip
	s.Port = port
}

// cria a conexão tcp com os clientes, chamando a goroutine para tratá-los
func (s Server) Service() {
	log.Printf("[+]Servidor em execucao em %s %s \n", s.Ip, s.Port)
}
