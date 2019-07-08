package controllers

import (
	"encoding/json"
	"web_api/models"
)

type OrderController struct {
	baseController
}

type OrderStruct struct {
	Id     int         `json:"id"`
	Price  float64     `json:"price"`
	UserId int         `json:"user_id"`
	Remark interface{} `json:"remark"`
}

type Remark struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (this *OrderController) List() {
	var orderModel models.OrderModel
	_, data := orderModel.GetOneInfo()
	num := len(data)
	orderStruct := make([]OrderStruct, num)
	// data.remark字段是json字符串，需要解析
	for key, value := range data {
		orderStruct[key].Id = value.Id
		orderStruct[key].UserId = value.UserId
		orderStruct[key].Price = value.Price
		var remark []Remark
		json.Unmarshal([]byte(value.Remark), &remark)
		orderStruct[key].Remark = remark
	}
	this.result.Data = orderStruct
	this.Success()
}

func (this *OrderController) Test() {
}
