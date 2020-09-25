package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lechengbao/datamodels"
	"github.com/lechengbao/datasource"
)

func CreateUser(user *datamodels.User) error {
	if user == nil {
		return errors.New("user params is nil")
	}

	return  datasource.PqDB.Create(user).Error
}

func GetUser(query interface{}, args ...interface{}) ( *datamodels.User, error) {
	var user = &datamodels.User{} 	//这里需要对user初始化，创建实例，否则读取失败
	err := datasource.PqDB.Where(query, args...).Take(user).Error	//在这里犯了一个错误，先写Take再写Where，查询是：select * form le_users Limit 1
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
