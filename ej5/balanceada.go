// EJERCICIO 5 · Responsable: Madelein · Revisa: Celery
package ej5

// BalanceadaOriginal es el código del enunciado, con DOS defectos.
// NO lo borren ni lo arreglen aquí: la prueba debe fallar con este.
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

// ======================= AQUÍ ESCRIBES =======================
// Una función llamada Balanceada(formula string) bool: es BalanceadaOriginal
// con los DOS defectos corregidos (los que anotaste en P5.L1 y P5.L2 de la hoja).
// Encima de cada línea que cambies, escribe un comentario // que diga
// qué defecto arregla y por qué. Ese comentario vale 20 % del ejercicio.
// =============================================================
