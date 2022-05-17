package main

// 标准库的包
import (
	"fmt"
)

//func main必须需要在名叫main的package下
func main() {
	//fmt.Println("hello go go go")
	//test()
	//constTest()
	//coreReturn()
	//zhizhen()
	//array()
	//sliceTest()
	slice2Test()
}

// 单独声明
var name string

const length int = 1

// 批量声明
var (
	age  int
	high int
)

func test() {
	age = 100
	name = "go"
	high = 999
	fmt.Printf("name: %s\n", name)
	fmt.Println("age+high", age+high)

	c := 7 // :=  只能用在函数体内声明

	fmt.Println("c= ,", c)
	fmt.Println("length, ", length)
}

const constName = "abc"

// iota只能在const中使用，const第一行 默认iota为0，按行递增加1
const (
	a, b = 1, 2
	c    = iota     // iota = 1
	d    = iota * 2 // iota = 2
	e    = iota + 3 // iota = 3
)

func constTest() {
	fmt.Println("constName, ", constName)

	fmt.Println("a,", a, "b,", b)
	fmt.Println("c, ", c, "d,", d, "e,", e)
}

func coreReturn() {
	returnTest(1, "000")
	fmt.Println("-----------  test2 -----")
	m, n := returnTest1(2, "100")
	fmt.Println("mc:", m, "nc", n)
}

// 返回值的测试中
func returnTest(a int, b string) int {
	fmt.Println("a,", a, "b,", b)
	return 100
}

func returnTest1(a int, b string) (m, n int) {
	fmt.Println("a,", a, "b,", b)

	// m,n 作用域在func中
	fmt.Println("m,", m, "n,", n)

	m = 1000
	n = 2000

	return
}

func swap(za *int, zb *int) {
	var swap int
	swap = *za
	fmt.Println("za", *za)
	*za = *zb
	*zb = swap
}

func zhizhen() {
	defer fmt.Println("end1")
	// 栈形式存储
	// defer 在return之后调用
	defer fmt.Println("end2")
	aa := 10
	bb := 20
	swap(&aa, &bb)
	fmt.Println("&aa,", &aa, "&bb", &bb)
	fmt.Println("aa, ", aa, "bb,", bb)
}

func arrayFunc(count []int) {
	// _匿名可忽略的值
	for _, value := range count {
		fmt.Println(value)
	}

	count[0] = 10 // 动态数组 切片的模式可以更改数据的值
}

func array() {
	count := []int{1, 2, 3} // 动态数组切片 slice

	arrayFunc(count)
	//for i := 0; i < len(count); i++ {
	//	fmt.Println(count[i])
	//}

	fmt.Println("=======")
	////range写法
	for index, value := range count {
		fmt.Println(index, value)
	}
}

func sliceTest() {
	// 方式1
	//slice1 := []int{1, 2, 3}
	// 方式2
	//var slice1 = []int{}
	//slice1 = make([]int, 3)
	// 方式3
	var slice1 []int
	fmt.Println("len = %d, slice = %v\n", len(slice1), slice1)

	if slice1 == nil {
		fmt.Println(" null")
	} else {
		fmt.Println("slice = %v\n", slice1)
	}

}

func slice2Test() {
	var slince2 []int = make([]int, 3, 5)
	fmt.Println("len = %d ,cap = %d, slice = %v", len(slince2), cap(slince2), slince2)

	slince2 = append(slince2, 1)
	fmt.Println("len = %d ,cap = %d, slice = %v", len(slince2), cap(slince2), slince2)

	slince2 = append(slince2, 1)
	fmt.Println("len = %d ,cap = %d, slice = %v", len(slince2), cap(slince2), slince2)

	// 按照cap的二倍扩容
	slince2 = append(slince2, 1)
	fmt.Println("len = %d ,cap = %d, slice = %v", len(slince2), cap(slince2), slince2)

	a := []int{12, 3, 4}
	fmt.Println(a)

	b := a[1:2]
	fmt.Println(b)

}
