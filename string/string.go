package main

import "fmt"

func main() {
	s := "hello，沙河"
	//byte和rune类型
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	//字符串修改
	s2 := "白萝卜"
	//把字符串强制转换为rune切片
	s3 := []rune(s2)
	s3[0] = '红'
	//把rune字符串强制转换为字符串并打印
	fmt.Println(string(s3))
}
