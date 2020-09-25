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
	var user = &datamodels.User{}
	err := datasource.PqDB.Take(user).Where(query, args...).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
