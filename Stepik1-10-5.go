/* Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
Кратное значит остаток от деления равен нулю
Не кратное значит остаток от деления НЕ равен нулю

Входные данные

Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
*/

package main

import "fmt"

func main() {

	var n int
	var c int
	var d int

	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&d)

	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		} else {
			//fmt.Println(i, " ", i%c, " ", i%d)
			continue
		}
	}

}
