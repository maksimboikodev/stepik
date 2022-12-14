/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

Входные данные
Сначала задано число NN — количество элементов в последовательности (1\leq N\leq1001≤N≤100). Далее через пробел записаны NN чисел — элементы последовательности. Последовательность состоит из целых чисел.
Выходные данные
Необходимо вывести единственное число - количество положительных элементов в последовательности.

Sample Input:
5
1 2 3 -1 -4
Sample Output:
3
*/
package main

import "fmt"

func main() {
	var n, sum, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&k)
		a[i] = k
		if a[i] >= 0 {
			sum++
		}
	}
	fmt.Print(sum)
}
