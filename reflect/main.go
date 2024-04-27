package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	//判断参数是不是结构体
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	//通过类型变量里面的Field获取结构体变量
	field0 := t.Field(0)
	fmt.Printf("%v\n", field0)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))

	//通过字段名称获取Field
	field1, ok := t.FieldByName("Name")
	if ok {
		fmt.Printf("%v\n", field1)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}

	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//获取结构体属性对应的值
	v := reflect.ValueOf(s)
	fmt.Printf("%v\n", v)
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Score"))
}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	//判断参数是不是结构体
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	//通过类型变量获取获取结构体的方法
	//方法的顺序与方法名的ASC码有关
	method0 := t.Method(0)
	fmt.Println(method0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	//通过方法获取方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	//调用无参方法
	v := reflect.ValueOf(s)
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)

	//调用有参方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params)

	info = v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)
}

// 反射修改结构体属性
func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的不是结构体指针类型")
		return
	}
	//修改结构体的属性
	name := v.Elem().FieldByName("Name")
	name.SetString("小李")

	age := v.Elem().FieldByName("Age")
	age.SetInt(22)
}

func main() {
	stu1 := Student{
		Name:  "小米",
		Age:   18,
		Score: 100,
	}
	// PrintStructField(&stu1)
	// PrintStructFn(stu1)
	// fmt.Printf("%#v\n", stu1)
	// var a = 11
	// reflectChangeStruct(&a)
	reflectChangeStruct(&stu1)

	fmt.Printf("%#v\n", stu1)
}
