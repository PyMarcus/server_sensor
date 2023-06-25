package main

import (
	"fmt"
	"os"

	"github.com/PyMarcus/server_sensor/actuator"
	"github.com/PyMarcus/server_sensor/sensor"
	"github.com/PyMarcus/server_sensor/server"
)

func main() {
	param := os.Args[1:]
	fmt.Print(param[1])
	if param[0] == "-server" {
		serv := server.Server{Ip: param[1], Port: param[2]}
		serv.Service()
	} else if param[0] == "-actuator" {
		actuador := actuator.Actuator{ServerIp: param[1], ServerPort: param[2]}
		actuador.Request()
	} else {
		senso := sensor.Sensor{ServerIp: param[1], ServerPort: param[2]}
		senso.Post()
	}
}
