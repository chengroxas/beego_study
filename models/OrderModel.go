package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type OrderModel struct {
	BaseModel
	Id     int
	Price  float64
	UserId int
	Remark string
}

func init() {
	orm.RegisterModel(new(OrderModel))
}

func (this *OrderModel) TableName() string {
	return "test_order"
}

func (this *OrderModel) GetOneInfo() (int64, []OrderModel) {
	var orders []OrderModel
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrderModel))
	num, _ := qs.All(&orders)
	return num, orders
}
