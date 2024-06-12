package _case

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

// 集合转列表
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(MapT[int64, user], 0)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "king", Age: 18}
	var userList List[user]
	userList = mapToList[int64, user](userMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(MapT[int64, address], 0)
	addrMp[1] = address{ID: 1, Province: "湖北", City: "武汉"}
	addrMp[2] = address{ID: 2, Province: "湖南", City: "长沙"}
	var addrList List[address]
	addrList = mapToList[int64, address](addrMp)

	ch1 := make(Chan[address])
	go myPrintln(ch1)
	for _, a := range addrList {
		ch1 <- a
	}
}

//泛型切片的定义
type List[T any] []T

//泛型集合定义
type MapT[k comparable, v any] map[k]v

//泛型通道的定义
type Chan[T any] chan T
