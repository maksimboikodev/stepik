/*
Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().

Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
 */

func task2(ch chan string, s string) {
for i := 1; i <= 5; i++ {
ch <- s + " "
}
}