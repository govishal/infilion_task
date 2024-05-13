package logic

import (
	"project/model"
	"project/mysql"
)

func GetPerson(Id string) (model.Person, error) {
	response, err := mysql.GetPersonInfo(Id)
	if err != nil {
		return response, err
	}
	return response, nil
}


func AddPerson(input model.Person) (bool, error) {
	isInserted, err := mysql.CreatePerson(input)
	if err != nil {
		return false, err
	}
	return isInserted, nil
}