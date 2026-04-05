/*
En un sistema academico basico, se necesita registrar y administrar
informacion de varios alumnos.
El sistema no maneja un solo alumno, sino varios, por lo que no es suficiente
crear variables individuales.
Se requiere una estructura que represente a un alumno (struct) y una lista que
almacene multiples alumnos (slice).
Un slice es una lista dinamica. Puede crecer conforme se agregan alumnos.

Mostrar el menu:
1. Agregar alumno
2. Mostrar todos los alumnos
3. Buscar alumno por nombre
4. Salir
*/
//CRUD son las 4 operaciones basicas para manejar datos en

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Alumno struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	Carrera  string `json:"carrera"`
	Promedio int    `json:"promedio"`
	Aprobado bool   `json:"aprobado"`
}

const archivoJSON = "alumnos.json"

var numeroId = 1
var alumnos = []Alumno{}
var reader = bufio.NewReader(os.Stdin)

func main() {
	cargarjson()

	opcion := 0
	for {
		fmt.Println("\n==========Menu==========")
		fmt.Println("1. Agregar alumno")
		fmt.Println("2. Mostrar todos los alumnos")
		fmt.Println("3. Buscar alumno por nombre")
		fmt.Println("4. Actualizar alumno")
		fmt.Println("5. Eliminar alumno")
		fmt.Println("6. Salir")
		fmt.Println("Seleccione una opcion:")
		fmt.Scan(&opcion)
		leerTexto()

		switch opcion {
		case 1:
			agregarAlumno()
			guardarjson()

		case 2:
			mostrarAlumnos()
		case 3:
			buscarAlumno()
		case 4:
			actualizarAlumno()
			guardarjson()
		case 5:
			eliminarAlumno()
			guardarjson()
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion no valida, intente de nuevo.")
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
	var aprobado bool
	fmt.Println("Ingrese el nombre del alumno:")
	nombre := leerTexto()
	if nombre == "" {
		fmt.Println("Error: No ingreso ningun nombre.")
		return
	}

	fmt.Println("Ingrese la edad del alumno:")
	fmt.Scan(&edad)
	// Validación: la edad no debe ser negativa.
	// Si el usuario ingresa un valor negativo, informamos el error
	// y abortamos la operación de agregado del alumno.
	if edad < 0 {
		fmt.Println("Error: la edad no puede ser negativa. Operación cancelada.")
		return
	}
	leerTexto() // Limpiar buffer

	fmt.Println("Ingrese la carrera del alumno:")
	carrera := leerTexto()
	if carrera == "" {
		fmt.Println("Error: No ingreso ninguna carrera.")
		return
	}

	fmt.Println("Ingrese el promedio del alumno:")
	fmt.Scan(&promedio)
	if promedio >= 70 {
		aprobado = true
	}
	nuevo := Alumno{
		Id:       numeroId,
		Nombre:   nombre,
		Edad:     edad,
		Carrera:  carrera,
		Promedio: promedio,
		Aprobado: aprobado,
	}
	numeroId++
	alumnos = append(alumnos, nuevo)
	fmt.Println("Alumno agregado exitosamente.")
	fmt.Println(alumnos)
}

func mostrarAlumnos() {
	if len(alumnos) == 0 {
		fmt.Println("No hay alumnos registrados.")
		return
	}

	fmt.Println("\n----Lista de Alumnos----")
	for i, a := range alumnos {
		fmt.Println("ID:", a.Id)
		fmt.Println("Alumno", i+1)
		fmt.Println("Nombre:", a.Nombre)
		fmt.Println("Edad:", a.Edad)
		fmt.Println("Carrera:", a.Carrera)
		fmt.Println("Promedio:", a.Promedio)
		fmt.Println("-----------------------")
	}
}

func buscarAlumno() {
	fmt.Println("Ingrese el nombre del alumno a buscar:")
	nombreBusqueda := leerTexto()

	if nombreBusqueda == "" {
		fmt.Println("Error: No ingreso nada para buscar.")
		return
	}

	encontrado := false
	for _, a := range alumnos {
		if strings.Contains(strings.ToLower(a.Nombre), strings.ToLower(nombreBusqueda)) {
			fmt.Println("Alumno encontrado:")
			fmt.Printf("Nombre: %s, Carrera: %s, Promedio: %d\n", a.Nombre, a.Carrera, a.Promedio)
			encontrado = true
		}
	}

	if !encontrado {
		fmt.Println("No se encontro ningun alumno con ese nombre o no es correcto.")
	}
}
func actualizarAlumno() {
	if len(alumnos) == 0 {
		fmt.Println("No hay alumnos registrados.")
		return
	}

	var id int
	fmt.Println("\nIngrese el ID del alumno: ")
	fmt.Scan(&id)
	leerTexto()

	for i := range alumnos {
		if alumnos[i].Id == id {
			var nuevaEdad int
			var nuevoPromedio int

			fmt.Println("Nuevo nombre:")
			nuevoNombre := leerTexto()

			fmt.Print("Nueva edad:")
			fmt.Scan(&nuevaEdad)
			leerTexto()

			fmt.Print("Nueva Carrera:")
			nuevaCarrera := leerTexto()

			fmt.Print("Nuevo promedio: ")
			fmt.Scanln(&nuevoPromedio)

			alumnos[i].Nombre = nuevoNombre
			alumnos[i].Edad = nuevaEdad
			alumnos[i].Carrera = nuevaCarrera
			alumnos[i].Promedio = nuevoPromedio
			alumnos[i].Aprobado = nuevoPromedio >= 70

			fmt.Println("Alumno actualizado exitosamente.")
			return
		}
	}
}
func eliminarAlumno() {
	if len(alumnos) == 0 {
		fmt.Println("No hay alumnos registrados.")
		return
	}
	var id int
	fmt.Println("Ingrese el ID del alumno a eliminar:")
	fmt.Scanln(&id)
	for i, a := range alumnos {
		if a.Id == id {
			alumnos = append(alumnos[:i], alumnos[i+1:]...)
			fmt.Println("Alumno eliminado exitosamente.")
			return
		}
	}
	fmt.Println("Alumno no encontrado.")
}

func guardarjson() {
	archivo, err := os.Create(archivoJSON)
	if err != nil {
		fmt.Println("Error al guardar archivo: ", err)
		return
	}
	defer archivo.Close()

	encoder := json.NewEncoder(archivo)
	encoder.SetIndent("", " ")
	err = encoder.Encode(alumnos)
	if err != nil {
		fmt.Println("Error al escribir JSON:", err)
	}
}

func cargarjson() {
	archivo, err := os.Open(archivoJSON)
	if err != nil {
		return
	}
	defer archivo.Close()

	decoder := json.NewDecoder(archivo)
	err = decoder.Decode(&alumnos)
	if err != nil {
		fmt.Println("Error al leer JSON:", err)
		return
	}

	maxID := 0
	for _, a := range alumnos {
		if a.Id > maxID {
			maxID = a.Id
		}
	}
	numeroId = maxID + 1

	fmt.Println("Datos cargados desde", archivoJSON)

}
