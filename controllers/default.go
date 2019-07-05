package controllers

// import "encoding/json"

type MainController struct {
	baseController
	// LoginController
}

type People struct {
	Name string `form:"name" json:"name"`
	Id   int    `form:"id" json:"id"`
}

func (this *MainController) Get() {

}

func (this *MainController) Post() {
	// body json
	// var people People
	// err := json.Unmarshal(this.Ctx.Input.RequestBody, &people)
	// if err != nil {
	// 	this.Ctx.WriteString(err.Error())
	// 	this.Finish()
	// }
	var people People
	err := this.ParseForm(&people)
	this.CheckError(err)
	this.result.Data = people
	this.Success()
}
