package main

import (
	"fmt"
	"reflect"
)

/**
https://www.cxyzjd.com/article/somanlee/108304525
*/

type class struct {
	ClassName string //导出字段
	student
	teacher teacher // 非导出字段
}

type student struct {
	name string
	age  int
}

type teacher struct {
	name string
	age  string
}

func studentF() {
	student2 := student{"xiaoming", 22}
	fmt.Println("学生名字："+student2.name+", 年龄：", student2.age)
}

func teacherF(teacher2 teacher) {
	fmt.Println("老师名字："+teacher2.name+", 年龄：", teacher2.age)

}

func main() {
	//// chan 有缓存区正常，无缓存中存入获取会死锁
	//ints := make(chan int, 1)
	//ints <- 1 //存入 箭头指向chan
	//
	//fmt.Println(<-ints) // 获取，减去chan
	//
	//select {
	//case <-ints:
	//	fmt.Println("1231231")
	//default:
	//	fmt.Println("========")
	//}
	//studentF(student{
	//	"test", 12,
	//})

	//more(student{"stu", 12})
	//more(teacher{"teacher", "22"})
	//	student2 := student{"xiaoming", 22}
	class2 := class{"大学一年级", student{"小红", 22}, teacher{"老李", "33"}}
	reflectTest(&class2)
}

//func more(v interface{}) {
//	switch v.(type) {
//	case teacher:
//		fmt.Println(v.(teacher).name)
//	case student:
//		fmt.Println(v.(student).name)
//	}
//}

func reflectTest(i interface{}) {
	//t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	//fmt.Println("_>...", t, v)
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println("for——>....", t.Field(i))
	//}
	//
	//vv := v.FieldByIndex([]int{1}) // 取层次，第一层
	//fmt.Println("vv_>...", vv)
	//
	//tv := v.FieldByName("teacher")
	//fmt.Println("tv->...", tv)
	e := v.Elem()
	//ee := reflect.ValueOf(i).Elem().FieldByName("ClassName")
	//	fmt.Println(e.FieldByName("className1"))
	//fmt.Println(e.FieldByIndex([]int{0}))
	e.FieldByName("ClassName1").SetString("大学二年级")
	fmt.Println("updateT->...", i)

}
