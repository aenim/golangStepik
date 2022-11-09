/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

Входные данные
Программа получает на вход три натуральных числа: x, p, y.

Выходные данные
Программа должна вывести одно целое число.
*/

package main

import "fmt"

func main() {
	var x float32
	var y float32
	var p float32
	var procent float32
	var count int

	fmt.Scan(&x, &p, &y)
	procent = 1 + p/100

	for count = 0; x < float32(y); count++ {
		x = x * procent
		//fmt.Println(x)
	}
	fmt.Println(count)
}
