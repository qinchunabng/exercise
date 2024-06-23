package models

import (
	"errors"

	"gorm.io/gorm"
)

func Transaction() {
	t := teacherTemp
	t1 := teacherTemp

	//返回err不为nil，则回滚
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		if err := tx.Create(&t1).Error; err != nil {
			return err
		}
		return nil
	})
}

func NestTrasaction() {
	t := teacherTemp
	t1 := teacherTemp
	t2 := teacherTemp
	t3 := teacherTemp

	//返回err不为nil，则回滚
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}

		//回滚子事务不影响外部事务
		tx.Transaction(func(tx1 *gorm.DB) error {
			tx1.Create(&t1)
			return errors.New("rollback t1")
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(&t2).Error; err != nil {
				return err
			}
			return nil
		})

		if err := tx.Create(&t3).Error; err != nil {
			return err
		}
		return nil
	})
}
