/*Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных
На вход подается номер билета - одно шестизначное  число.
Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:
613244
Sample Output:
YES
*/
package main

import "fmt"

func main() {
	var a, sum1, sum2 int
	c := 0
	fmt.Scan(&a)
	for i := a; a != 0; i-- {
		b := a % 10
		a = a / 10
		c++
		if c > 3 {
			sum1 += b
		} else {
			sum2 += b
		}
	}
	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}