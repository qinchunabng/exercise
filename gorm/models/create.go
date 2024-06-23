package models

import (
	"fmt"
	"gorm/utils"
	"log"
	"time"
)

var teacherTemp = Teacher{
	Name:         "nick",
	Age:          40,
	WorkingYears: 10,
	Email:        "nick@voice.com",
	Birthday:     time.Now().Unix(),
	StuNumber: struct {
		String string
		Valid  bool
	}{String: "10", Valid: true},
	Roles: []string{"普通用户", "讲师"},
	JobInfo: Job{
		Title:    "讲师",
		Location: "湖南长沙",
	},
	JobInfo2: Job{
		Title:    "讲师",
		Location: "湖南长沙",
	},
}

func CreateRecord() {
	t := teacherTemp
	res := DB.Create(&t)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	//主键ID会被反向填充
	utils.Println(t)
	t1 := teacherTemp

	//正向选择
	res = DB.Select("Name", "Age").Create(&t1)
	utils.Println(res.RowsAffected, res.Error, t1)

	//反向选择
	t2 := teacherTemp
	res = DB.Omit("Email", "Birthday").Create(&t2)
	utils.Println(res.RowsAffected, res.Error, t1)

	//批量插入
	var teachers = []Teacher{{
		Name: "King",
		Age:  40,
	}, {
		Name: "Darren",
		Age:  35,
	}, {
		Name: "Nick",
		Age:  33,
	}}
	DB.Create(teachers)
	for _, t := range teachers {
		fmt.Println(t.ID)
	}

	//分批操作
	var teachers1 = []Teacher{{
		Name: "King",
		Age:  40,
	}, {
		Name: "Darren",
		Age:  35,
	}, {
		Name: "Nick",
		Age:  33,
	}}
	DB.CreateInBatches(teachers1, 2)
	for _, t := range teachers {
		fmt.Println(t.ID)
	}
}
