package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Alumno struct {
	Nombre  string
	Edad    int
	Carrera string
}

// Función auxiliar para validar que el string solo tenga letras y espacios
func esSoloLetras(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return s != ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// --- Inicializaciones previas ---
	alumno1 := Alumno{"Alberto", 21, "ITICS"}

	// Corregido: nombre del struct era Alumno, no Alumno2
	alumno2 := Alumno{
		Nombre:  "Abi",
		Edad:    21,
		Carrera: "Mecatronica",
	}

	var alumno3 Alumno
	alumno3.Nombre = "Kike"
	alumno3.Edad = 21
	alumno3.Carrera = "Civil"

	// --- d) Leer datos desde teclado con VALIDACIÓN ---
	var alumno4 Alumno

	// 1. Validar Nombre
	for {
		fmt.Print("Ingresa el nombre del alumno 4: ")
		nombreInput, _ := reader.ReadString('\n')
		nombreInput = strings.TrimSpace(nombreInput) // Quita el salto de línea

		if esSoloLetras(nombreInput) {
			alumno4.Nombre = nombreInput
			break
		}
		fmt.Println("Error: El nombre solo debe contener letras.")
	}

	// 2. Validar Carrera
	for {
		fmt.Print("Ingresa la carrera del alumno 4: ")
		carreraInput, _ := reader.ReadString('\n')
		carreraInput = strings.TrimSpace(carreraInput)

		if esSoloLetras(carreraInput) {
			alumno4.Carrera = carreraInput
			break
		}
		fmt.Println("Error: La carrera solo debe contener letras.")
	}

	// 3. Validar Edad (Solo números y > 0)
	for {
		fmt.Print("Ingresa la edad del alumno 4: ")
		edadInput, _ := reader.ReadString('\n')
		edadInput = strings.TrimSpace(edadInput)

		edadNum, err := strconv.Atoi(edadInput) // Intenta convertir a entero
		if err == nil && edadNum > 0 {
			alumno4.Edad = edadNum
			break
		}
		fmt.Println("Error: Ingresa una edad válida (número mayor a 0).")
	}

	fmt.Println("\n Alumno 1 ")
	fmt.Println("Nombre:", alumno1.Nombre, "| Edad:", alumno1.Edad, "| Carrera:", alumno1.Carrera)

}
