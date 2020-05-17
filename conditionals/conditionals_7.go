/*
Если число больше нуля, выведите сначала фразу "Positive number. The next number for the number ", затем введённое число, затем слово " is ", окруженное пробелами, затем формулу для следующего за введенным числа, наконец, знак точки без пробела. Для предыдущего числа пишем фразу "Positive number. The previous number for the number ", а затем всё вышеперечисленное.

Если число меньше нуля, выведите сначала фразу "Negative number. The next number for the number ", а затем делаем всё то же самое (то же - и для предыдущего числа). В данном задании для отрицательного n числа, для следующего (next number) числа нужно вывести n + 1. А для предыдущего n - 1. Например, если на вход подается число -317 то нужно вывести:

Negative number. The next number for the number -317 is -316.
Negative number. The previous number for the number -317 is -318.

Если число равно нулю - ничего выводить не нужно.
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		fmt.Print("Positive number. The next number for the number ", a, " is ", a+1, ".\n")
		fmt.Print("Positive number. The previous number for the number ", a, " is ", a-1, ".\n")
	} else if a < 0 {
		fmt.Print("Negative number. The next number for the number ", a, " is ", a+1, ".\n")
		fmt.Print("Negative number. The previous number for the number ", a, " is ", a-1, ".\n")
	} else {
		return
	}

}
