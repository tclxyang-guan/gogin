package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func (ctrl *TestController) Router(router *gin.Engine) {
	router.POST("Page/create", ctrl.create)
	router.GET("Page/update", ctrl.update)
}

func (ctrl *TestController) create(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		return
	}
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)
	fmt.Println(m)
}
func (ctrl *TestController) update(ctx *gin.Context) {
	a := ctx.Request.FormValue("a")
	fmt.Println(a)
}
