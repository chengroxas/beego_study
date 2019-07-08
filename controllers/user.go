package controllers

import (
	"web_api/models"
)

type UserController struct {
	baseController
	// LoginController
}

type User struct {
	Name string `form:"name" json:"name"`
	Id   int    `form:"id" json:"id"`
}

type ResultList struct {
	List interface{} `json:"list"`
	Num  int64       `json:"num"`
}

// user list
func (this *UserController) List() {
	var userModel models.UserModel
	num, data := userModel.GetOneInfo()
	users := make([]User, len(data)) //使用slice可以解决不定长数组的情况
	for key, value := range data {
		users[key].Name = value.Name
		users[key].Id = value.Id
	}
	var list ResultList
	list.List = users
	list.Num = num
	this.result.Data = list
	this.Success()
}
