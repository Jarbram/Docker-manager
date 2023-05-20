package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func StartContainer(containerName string) {
	// Ejecutar el comando "docker start" para encender el contenedor especificado
	cmd := exec.Command("docker", "start", containerName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error al encender el contenedor: %s", err)
	}

	// Imprimir mensaje de confirmación
	fmt.Println("Contenedor encendido exitosamente.")
}
