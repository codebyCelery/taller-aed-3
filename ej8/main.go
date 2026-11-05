package main

import "fmt"
func main() {
	var raiz *Nodo

	notas := []int{14, 11, 17, 11, 9, 14, 20, 14, 12, 20}

	// Insertar notas
	for _, v := range notas {
		raiz = insertar(raiz, v)
	}

	fmt.Println("Arbol inicial:")
	inorden(raiz)
	fmt.Println()

	fmt.Println("Cantidad de 14:", cantidad(raiz, 14))
	fmt.Println("Cantidad de 11:", cantidad(raiz, 11))
	fmt.Println("Cantidad de 20:", cantidad(raiz, 20))

	// Eliminar notas
	eliminaciones := []int{14, 11, 17, 14, 14}

	for _, v := range eliminaciones {
		raiz = eliminar(raiz, v)
	}

	fmt.Println("\nArbol final:")
	inorden(raiz)
	fmt.Println()

	fmt.Println("Cantidad de 14:", cantidad(raiz, 14))
	fmt.Println("Cantidad de 11:", cantidad(raiz, 11))
	fmt.Println("Cantidad de 20:", cantidad(raiz, 20))
}
