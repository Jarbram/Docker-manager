package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func InspectContainer(containerName string) {
	// Ejecutar el comando "docker inspect" para obtener información del contenedor
	cmd := exec.Command("docker", "inspect", containerName)
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error al inspeccionar el contenedor: %s", err)
	}

	// Imprimir la información del contenedor en la salida estándar
	fmt.Println(string(output))
}
