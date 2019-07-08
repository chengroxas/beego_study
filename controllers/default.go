package controllers

import (
	"encoding/json"
)

type MainController struct {
	baseController
	// LoginController
}

type People struct {
	Name string `form:"name" json:"name"`
	Id   int    `form:"id" json:"id"`
}

//解析 formdata的数据
func (this *MainController) MyParseForm() {
	this.Abort("Db")
	var people People
	this.ParseForm(&people)
	this.result.Data = people
	this.Success()
}

//解析 request body json的数据
func (this *MainController) ParseJson() {
	var people People
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &people)
	this.CheckError(err)
	this.result.Data = people
	this.Success()
}
