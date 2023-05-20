package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func ShowAvailablePorts() {
	// Ejecutar el comando "docker port" para obtener los puertos disponibles
	cmd := exec.Command("docker", "port")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error al obtener los puertos disponibles: %s", err)
	}

	// Imprimir los puertos disponibles en la salida estándar
	fmt.Println("Puertos disponibles:")
	fmt.Println(string(output))
}
