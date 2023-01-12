/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
*/

package main

import "fmt"

func main() {
	var first_number int
	//var second_number int

	//вычленяем цифры из числа
	for fmt.Scan(&first_number); first_number >= 1; first_number /= 10 {
		//fmt.Print(first_number%10, " ")
	}
	fmt.Printf("%b", 1564846156116818117)

}

// вывод делается в обратном порядке а надо в прямом
