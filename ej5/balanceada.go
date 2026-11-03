package ej5

// EJERCICIO 5 - Responsable: Madelein - Revisa: Celery

func BalanceadaOriginal(formula string) bool {
	pares := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var p Pila[rune]
	for _, c := range formula {
		if c == '(' || c == '[' || c == '{' {
			p.Push(c)
		} else {
			tope, err := p.Pop()
			if err != nil || tope != pares[c] {
				return false
			}
		}
	}
	return true
}

func Balanceada(formula string) bool {
	pares := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var p Pila[rune]

	for _, c := range formula {
		if c == '(' || c == '[' || c == '{' {
			p.Push(c)

			// Solo los símbolos de cierre deben sacar elementos.
			// Así evitamos que las letras provoquen errores.
		} else if c == ')' || c == ']' || c == '}' {
			tope, err := p.Pop()

			// Comparamos el cierre con la apertura que le corresponde.
			// Si no coincide o no hay apertura, la fórmula es incorrecta.
			if err != nil || tope != pares[c] {
				return false
			}
		}
	}

	// Comprobamos que no queden símbolos abiertos.
	// Antes se aceptaban fórmulas incompletas como "((".
	return p.Vacia()
}
