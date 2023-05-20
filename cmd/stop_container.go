package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func StopContainer(containerName string) {
	// Ejecutar el comando "docker stop" para apagar el contenedor especificado
	cmd := exec.Command("docker", "stop", containerName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error al apagar el contenedor: %s", err)
	}

	// Imprimir mensaje de confirmación
	fmt.Println("Contenedor apagado exitosamente.")
}
