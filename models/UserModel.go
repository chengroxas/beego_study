package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	BaseModel
	Name    string
	Id      int
	Age     int
	Comment string
}

func init() {
	orm.RegisterModel(new(UserModel))
}

func (this *UserModel) TableName() string {
	return "test_user"
}

func (this *UserModel) GetOneInfo() (int64, []UserModel) {
	var users []UserModel
	o := orm.NewOrm()
	qs := o.QueryTable(this.TableName())
	num, _ := qs.All(&users)
	return num, users
}
