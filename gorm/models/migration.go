package models

func init() {
	DB.AutoMigrate(Teacher{})
}
