package models

import "gorm/utils"

func Query() {
	//查询单条数据
	//查询第一条
	t := Teacher{}
	res := DB.First(&t)
	utils.Println(res.RowsAffected, res.Error, t)

	//查询最后一条
	t = Teacher{}
	res = DB.Last(&t)
	utils.Println(res.RowsAffected, res.Error, t)

	//无排序，取最后一条
	t = Teacher{}
	res = DB.Take(&t)
	utils.Println(res.RowsAffected, res.Error, t)

	//将结果填充到集合，不支持特殊类型处理，无法完成类型转换
	result := map[string]interface{}{}
	res = DB.Model(&Teacher{}).Omit("Birthday", "Roles", "JobInfo2").First(&result)
	utils.Println(res.RowsAffected, res.Error, result)

	//基于表名，不支持First和Last
	result = map[string]interface{}{}
	res = DB.Table("teachers").Take(&result)
	utils.Println(res.RowsAffected, res.Error, result)

	var teachers []Teacher
	res = DB.Where("name=?", "Nick").Or("name=?", "King").Order("id desc").Limit(10).Find(&teachers)
	utils.Println(res.RowsAffected, res.Error, teachers)
}
