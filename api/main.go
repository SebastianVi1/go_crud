package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var (
	alumnos  []Alumno
	numeroId = 1
)

type Alumno struct {
	Id       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Edad     int    `json:"edad"`
	Carrera  string `json:"carrera"`
	Promedio int    `json:"promedio"`
	Aprobado bool   `json:"aprobado"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func main() {

	cargarDatosPrueba()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /alumnos", obtenerAlumnos)
	mux.HandleFunc("GET /alumnos/{id}", obtenerAlumnoPorId)
	mux.HandleFunc("POST /alumnos", crearAlumno)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", enableCORS(mux))

}

func cargarDatosPrueba() {
	datos := []Alumno{
		{Nombre: "Ana", Edad: 20, Carrera: "ISC", Promedio: 90},

		{Nombre: "Luis", Edad: 21, Carrera: "ISC", Promedio: 90},
		{Nombre: "Maria", Edad: 22, Carrera: "ITIC", Promedio: 85},
		{Nombre: "Carlos", Edad: 19, Carrera: "Industrial", Promedio: 70},
		{Nombre: "Sofia", Edad: 23, Carrera: "Mecatronica", Promedio: 95},
	}

	for _, d := range datos {
		d.Id = numeroId
		d.Aprobado = d.Promedio >= 70
		numeroId++
		alumnos = append(alumnos, d)
	}

}

func obtenerAlumnos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alumnos)
}

func obtenerAlumnoPorId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idTexto := r.PathValue("id")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		http.Error(w, "Id invalido", http.StatusBadRequest)
		return
	}

	for _, a := range alumnos {
		if a.Id == id {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	http.Error(w, "Alumno no encontrado", http.StatusNotFound)
}

func crearAlumno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nuevo Alumno
	err := json.NewDecoder(r.Body).Decode(&nuevo)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(nuevo.Nombre) == "" {
		http.Error(w, "Nombre obligatorio", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(nuevo.Carrera) == "" {
		http.Error(w, "Carrera obligatoria", http.StatusBadRequest)
		return
	}
	if nuevo.Edad <= 0 {
		http.Error(w, "Edad invalida", http.StatusBadRequest)
		return
	}
	if nuevo.Promedio < 0 || nuevo.Promedio > 100 {
		http.Error(w, "Promedio invalido", http.StatusBadRequest)
		return
	}

	nuevo.Id = numeroId
	nuevo.Aprobado = nuevo.Promedio >= 70
	numeroId++

	alumnos = append(alumnos, nuevo)
	fmt.Println("Alumno creado", nuevo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}
