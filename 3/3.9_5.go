/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
 */
package main

import "sync"

// вы уже внутри main()
var wg sync.WaitGroup
// wg := &sync.WaitGroup{}
// wg := new(sync.WaitGroup)

for i := 0; i < 10; i++ {
wg.Add(1)
go func() {
defer wg.Done()
work()
}()
}
wg.Wait()