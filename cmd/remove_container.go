package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func RemoveContainer(containerName string) {
	// Ejecutar el comando "docker rm" para eliminar el contenedor especificado
	cmd := exec.Command("docker", "rm", containerName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error al eliminar el contenedor: %s", err)
	}

	// Imprimir mensaje de confirmación
	fmt.Println("Contenedor eliminado exitosamente.")
}
