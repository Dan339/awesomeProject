package main

import (
	"fmt"
	"sync"
)

// 注意不是用指针的时候，会产生死锁问题
// todo panic: sync: negative WaitGroup counter
func core(group *sync.WaitGroup) {
	fmt.Println("hello ")
	group.Done()
}

//func main() {
//	//go core()
//	//i := 0
//	//for i < 10 {
//	//	i++
//	//	fmt.Println(i)
//	//}
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go core(&wg)
//	wg.Wait()
//
//}
