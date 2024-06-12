package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不使用泛型，数字比较：", getMaxNumInt(a, b))
	fmt.Println("不使用泛型，数字比较：", getMaxNumFloat(c, d))

	//由编译器推断输入的类型
	fmt.Println("使用泛型，数字比较：", getMaxNum(a, b))
	//显式指定传入的类型
	fmt.Println("使用泛型，数字比较：", getMaxNum[float64](c, d))
}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CustNumT interface {
	//支持uint8、int32、float64与int64及其衍生类型
	//~表示支持类型的衍生类型
	//|表示取并集
	//多行之间取交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

//MyInt64的衍生类型，是具有基础类型int64的新类型，与int64是不同的类型
type MyInt64 int64

//MyInt32为int32的别名，与int32是同一类型
type MyInt32 = int32

func CustNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a, b))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a1, b1))

	var c, d float64 = 5, 6
	//由编译器推断输入的类型
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(c, d))

	var e, f float64 = 7, 8
	// var g, h MyInt64 = e, f
	fmt.Println("衍生类型，数字比较：", getMaxCusNum(e, f))
}

func getMaxCusNum[T CustNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

//内置类型
func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("内置comparable泛型类型约束", getBuiltInComparable(a, b))
	var c, d = 100, 100
	fmt.Println("内置comparable泛型类型约束", getBuiltInComparable(c, d))

	var f = 100.123
	getBuiltInAny(f)
}

func getBuiltInComparable[T comparable](a, b T) bool {
	//comparable类型只支持== !=操作
	return a == b
}

func getBuiltInAny[T any](a T) {
	fmt.Println("内置any泛型类型约束", a)
}
