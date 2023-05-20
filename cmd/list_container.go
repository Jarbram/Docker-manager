package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func ListContainers() {
	// Ejecutar el comando "docker ps" para obtener la lista de contenedores
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}\t{{.Status}}")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error al obtener la lista de contenedores: %s", err)
	}

	// Imprimir la lista de contenedores en la salida estándar
	fmt.Println(string(output))
}
