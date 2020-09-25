package datamodels

import "github.com/jinzhu/gorm"

func GetModelList() []interface{}{
	//表名加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "le_" + defaultTableName
	}

	return []interface{}{
		&User{},
	}
}
