package _case

import "fmt"

//基本接口，可用于变量的定义
type ToString interface {
	String() string
}

func (u user) String() string {
	return fmt.Sprintf("ID: %d,Name: %s,Age:%d", u.ID, u.Name, u.Age)
}

func (addr address) String() string {
	return fmt.Sprintf("ID: %d,Province: %s,Age:%s", addr.ID, addr.Province, addr.City)
}

// var s ToString

//泛型接口，不可以作为变量声明
type GetKey[T comparable] interface {
	any
	Get() T
}

func (u user) Get() int64 {
	return u.ID
}

func (addr address) Get() int {
	return addr.ID
}

// var s GetKey[string]

//列表转集合
func listToMap[k comparable, T GetKey[k]](list []T) map[k]T {
	mp := make(MapT[k, T], len(list))
	for _, data := range list {
		mp[data.Get()] = data
	}
	return mp
}

func InterfaceCase() {
	userList := []GetKey[int64]{
		user{ID: 1, Name: "nick", Age: 18},
		user{ID: 2, Name: "king", Age: 18},
	}
	userMp := listToMap[int64, GetKey[int64]](userList)
	fmt.Println(userMp)

	addrList := []GetKey[int]{
		address{ID: 1, Province: "湖北", City: "武汉"},
		address{ID: 2, Province: "湖南", City: "长沙"},
	}
	addrMp := listToMap[int, GetKey[int]](addrList)
	fmt.Println(addrMp)

}
