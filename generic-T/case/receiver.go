package _case

import "fmt"

//泛型结构体
//1.不支持断言
//2.不支持泛型方法
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

//泛型receiver
func (myStruct MyStruct[T]) GetData() T {
	return myStruct.Data
}

func ReceiverCase() {
	data := 8
	myStruct := MyStruct[*int]{
		Name: "nick",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println(*data1)

	str := "abcdefg"
	myStruct1 := MyStruct[*string]{
		Name: "nick",
		Data: &str,
	}
	str1 := myStruct1.GetData()
	fmt.Println(*str1)
}
