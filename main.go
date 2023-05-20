package main

import (
	"bufio"
	"dockerManager/cmd"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar el menú de opciones
		fmt.Println("=== Menú ===")
		fmt.Println("1. Obtener lista de contenedores")
		fmt.Println("2. Crear contenedor")
		fmt.Println("3. Eliminar contenedor")
		fmt.Println("4. Encender contenedor")
		fmt.Println("5. Apagar contenedor")
		fmt.Println("6. Inspeccionar contenedor")
		fmt.Println("7. Salir")
		fmt.Print("Ingrese la opción: ")

		// Leer la opción seleccionada por el usuario
		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)

		switch optionStr {
		case "1":
			cmd.ListContainers() // Llamar a la función para obtener la lista de contenedores
		case "2":
			cmd.CreateContainer() // Llamar a la función para crear un contenedor
		case "3":
			fmt.Print("Ingrese el nombre del contenedor a eliminar: ")
			containerName, _ := reader.ReadString('\n')
			containerName = strings.TrimSpace(containerName)
			cmd.RemoveContainer(containerName) // Llamar a la función para eliminar un contenedor
		case "4":
			fmt.Print("Ingrese el nombre del contenedor a encender: ")
			containerName, _ := reader.ReadString('\n')
			containerName = strings.TrimSpace(containerName)
			cmd.StartContainer(containerName) // Llamar a la función para encender un contenedor
		case "5":
			fmt.Print("Ingrese el nombre del contenedor a apagar: ")
			containerName, _ := reader.ReadString('\n')
			containerName = strings.TrimSpace(containerName)
			cmd.StopContainer(containerName) // Llamar a la función para apagar un contenedor
		case "6":
			fmt.Print("Ingrese el nombre del contenedor a inspeccionar: ")
			containerName, _ := reader.ReadString('\n')
			containerName = strings.TrimSpace(containerName)
			cmd.InspectContainer(containerName) // Llamar a la función para inspeccionar un contenedor
		case "7":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida. Por favor, seleccione una opción válida.")
		}
	}
}
