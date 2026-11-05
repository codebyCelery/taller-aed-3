package main

import "fmt"

func main() {
	var raiz *Nodo

	// Inserta las notas del enunciado, en orden.
	for _, v := range []int{14, 11, 17, 11, 9, 14, 20, 14, 12, 20} {
		raiz = insertar(raiz, v)
	}

	// Elimina las notas del enunciado, en orden.
	for _, v := range []int{14, 11, 17, 14, 14} {
		raiz = eliminar(raiz, v)
	}

	// Imprime el inorden final: debe coincidir con P8.IN de la hoja.
	fmt.Print("Inorden final: ")
	inorden(raiz)
	fmt.Println()
}
