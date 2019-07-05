package controllers

//继承了LoginController都要登录
type LoginController struct {
	baseController
}

func (this *LoginController) IsLogin() {
	if this.isLogin == false {
		this.result.Code = 10021
		this.result.Msg = "请先登录"
		this.Fail()
	}
}
