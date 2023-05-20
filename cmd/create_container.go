package cmd

import (
	"bufio"
	"dockerManager/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CreateContainer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre del contenedor: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Imagen base: ")
	image, _ := reader.ReadString('\n')
	image = strings.TrimSpace(image)

	utils.ShowAvailablePorts()
	fmt.Print("Puerto a exponer: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	fmt.Print("Comando a ejecutar: ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	// Construir el comando para crear el contenedor
	dockerRunCommand := fmt.Sprintf("docker run -d --name %s -p %s %s %s", name, port, image, command)
	// Ejecutar el comando para crear el contenedor
	cmd := exec.Command("sh", "-c", dockerRunCommand)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error al crear el contenedor: %s", err)
	}

	fmt.Println("Contenedor creado exitosamente.")
}
