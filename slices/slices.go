package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Alumno struct {
	Nombre   string
	Edad     int
	Carrera  string
	Promedio int
	Aprobado bool
}

var reader = bufio.NewReader(os.Stdin)
var alumnos []Alumno

func main() {
	opcion := 0
	for {
		fmt.Println("\n==Menú de opciones")
		fmt.Println("1. Agregar alumno")
		fmt.Println("2. Mostrar todos los alumnos")
		fmt.Println("3. Buscar alumno por nombre")
		fmt.Println("4. Salir")
		fmt.Println("Selecciona la opcion")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			agregarAlumno()
		case 2:
			mostrarAlumnos()
		case 3:
			buscarAlumno()
		case 4:
			return
		default:
			fmt.Println("Opcion No valida")
		}
	}
}

func leerTexto() string {
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func agregarAlumno() {
	var edad int
	var promedio int

	fmt.Println("Nombre del Alumno:")
	nombre := leerTexto()
	fmt.Println("Edad del Alumno:")
	fmt.Scanln(&edad)
	fmt.Println("Carrera del Alumno:")
	carrera := leerTexto()
	fmt.Println("Promedio del alumno:")
	fmt.Scanln(&promedio)

	aprobado := promedio >= 60
	nuevo := Alumno{
		Nombre:   nombre,
		Edad:     edad,
		Carrera:  carrera,
		Promedio: promedio,
		Aprobado: aprobado,
	}

	alumnos = append(alumnos, nuevo)
	fmt.Println("Alumno agregado correctamente")
}

func mostrarAlumnos() {
	if len(alumnos) == 0 {
		fmt.Println("No hay alumnos registrados")
		return
	}
	fmt.Println("\n=====Lista de Alumnos=====")
	for i, a := range alumnos {
		fmt.Println("Alumno", i+1)
		fmt.Println("  Nombre:", a.Nombre)
		fmt.Println("  Edad:", a.Edad)
		fmt.Println("  Carrera:", a.Carrera)
		fmt.Println("  Promedio:", a.Promedio)
		fmt.Println("  Aprobado:", a.Aprobado)
		fmt.Println()
	}
}

func buscarAlumno() {
	fmt.Println("Ingresa el nombre a buscar:")
	nombre := leerTexto()
	encontrado := false
	for i, a := range alumnos {
		if strings.EqualFold(a.Nombre, nombre) {
			fmt.Println("Alumno encontrado (posición", i+1, ")")
			fmt.Println("  Nombre:", a.Nombre)
			fmt.Println("  Edad:", a.Edad)
			fmt.Println("  Carrera:", a.Carrera)
			fmt.Println("  Promedio:", a.Promedio)
			fmt.Println("  Aprobado:", a.Aprobado)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Println("No se encontró ningún alumno con ese nombre")
	}
}
