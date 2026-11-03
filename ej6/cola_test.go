package ej6

import (
	"os"
	"testing"
)

// nueva devuelve la versión de la cola que se quiere probar.
// Se elige con la variable de entorno VERSION:
//
//	go test ./ej6            -> versión correcta (debe PASAR)
//	VERSION=X go test ./ej6  -> copia X (debe FALLAR)
//	VERSION=Y go test ./ej6  -> copia Y (debe FALLAR)
//	VERSION=Z go test ./ej6  -> copia Z (debe FALLAR)
func nueva(version string) Cola {
	switch version {
	case "X":
		return &ColaX{}
	case "Y":
		return &ColaY{}
	case "Z":
		return &ColaZ{}
	}
	return &ColaPedidos{}
}

// caso es una fila de la tabla de pruebas.
type caso struct {
	nombre    string
	secuencia string // E = Encolar (1, 2, 3 en orden), D = Desencolar
	esperado  int    // lo que retorna la versión CORRECTA en el último D
	esperaErr bool   // true si el último D debe dar error
}

// Las secuencias son las de la hoja (P6.X, P6.Y, P6.Z).
var tabla = []caso{
	{"detecta Z: un solo pedido", "ED", 1, false},
	{"detecta Y: orden al traspasar", "EED", 1, false},
	{"detecta X: traspaso con salida llena", "EEDED", 2, false},
}

// ejecutar corre la secuencia y devuelve lo que retorna el ÚLTIMO D.
func ejecutar(c Cola, seq string) (int, error) {
	siguiente := 1 // se encolan 1, 2, 3 en ese orden
	var v int
	var err error
	for _, op := range seq {
		if op == 'E' {
			c.Encolar(siguiente)
			siguiente++
		} else {
			v, err = c.Desencolar()
		}
	}
	return v, err
}

// TestColaPedidos es la prueba con tabla: pasa con la correcta y
// falla con cada una de las tres copias.
func TestColaPedidos(t *testing.T) {
	version := os.Getenv("VERSION")
	for _, c := range tabla {
		t.Run(c.nombre, func(t *testing.T) {
			v, err := ejecutar(nueva(version), c.secuencia)
			if (err != nil) != c.esperaErr || (!c.esperaErr && v != c.esperado) {
				t.Errorf("versión %q, secuencia %s: obtuvo (%d, %v), esperado (%d, error=%v)",
					version, c.secuencia, v, err, c.esperado, c.esperaErr)
			}
		})
	}
}
