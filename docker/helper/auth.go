package helper

import (
	"docker/configs"
	"docker/model"
)

func AuthDB(Name string, Password string)(bool, error){
	var user model.User
	result := configs.DB.First(&user,"name = ? AND password = ?", Name, Password)
	if result != nil {
		return false, result.Error
	}
	return true, nil
}