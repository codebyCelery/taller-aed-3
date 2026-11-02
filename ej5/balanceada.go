// EJERCICIO 5 · Responsable: Madelein · Revisa: Celery
package ej5

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

        } else if c == ')' || c == ']' || c == '}' {
            tope, err := p.Pop()

            if err != nil || tope != pares[c] {
                return false
            }
        }
    }

    return p.IsEmpty()
}
