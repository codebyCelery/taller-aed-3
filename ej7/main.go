package main

import "fmt"

func main() {
	var n Navegador

	// (a) Visitar p1, p2, ..., p7  -> P7.1
	for i := 1; i <= 7; i++ {
		n.Visitar(fmt.Sprintf("p%d", i))
	}
	a, f := n.Estado()
	fmt.Println("P7.1 después de visitar p1..p7")
	fmt.Println("     atras   :", a)
	fmt.Println("     adelante:", f)

	// Atras, Atras -> P7.2
	n.Atras()
	n.Atras()
	a, f = n.Estado()
	fmt.Println("P7.2 después de Atras, Atras")
	fmt.Println("     atras   :", a)
	fmt.Println("     adelante:", f)

	// Visitar p8 -> P7.3
	n.Visitar("p8")
	a, f = n.Estado()
	fmt.Println("P7.3 después de visitar p8")
	fmt.Println("     atras   :", a)
	fmt.Println("     adelante:", f)

	// Cuatro Atras seguidos -> P7.4 (página actual después de cada uno)
	fmt.Println("P7.4 cuatro Atras seguidos:")
	for i := 1; i <= 4; i++ {
		n.Atras()
		fmt.Printf("     Atras #%d -> actual: %s\n", i, n.Actual())
	}
}
