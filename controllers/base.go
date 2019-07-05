package controllers

import (
	"github.com/astaxie/beego"
)

type IsLogin interface {
	IsLogin()
}

type baseController struct {
	beego.Controller
	result  Result
	isLogin bool
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 空的对象 EmptyObject{} 空的数组 [0]EmptyObject{}
type EmptyObject struct{}

func (this *baseController) Prepare() {
	//做判断是否登录
	if app, ok := this.AppController.(IsLogin); ok {
		app.IsLogin()
	}
}

func (this *baseController) CheckError(err error) {
	if err != nil {
		this.result.Msg = err.Error()
		this.Fail()
	}
}

func (this *baseController) Fail() {
	this.result.Data = EmptyObject{}
	this.Data["json"] = this.result
	this.ServeJSON()
}

func (this *baseController) Success() {
	if this.result.Data == nil {
		this.result.Data = EmptyObject{}
	}
	this.Data["json"] = this.result
	this.ServeJSON()
}
