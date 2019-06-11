package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gogin/controller"
	"gogin/models"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	new(controller.TestController).Router(router)
	/*router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)*/

	return router
}

func HandleTest(g *gin.Context) {
	a := g.Request.FormValue("a")
	fmt.Println(a)
	/*m:=make(map[string]interface{})
	g.PureJSON(101,m)
	fmt.Println(m)*/
	data, err := g.GetRawData()
	if err != nil {
		return
	}
	//m:=make(map[string]interface{})
	t := models.Test{}
	json.Unmarshal(data, &t)
	fmt.Println(t)
}
