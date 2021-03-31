package web

import (
	"fmt"
	"github.com/weiqiang333/infra-skywalking-webhook/configs"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weiqiang333/infra-skywalking-webhook/internal/dingtalk"
)


// Web 路由入口
func Web() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.POST("/dingtalk", routerDingtalk)
	address := configs.V.GetString("address")
	err := r.Run(address) // default : listen and serve on 0.0.0.0:8000
	if err != nil {
		fmt.Println(err.Error())
	}
}


func routerDingtalk(c *gin.Context)  {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error Body": err.Error()})
		return
	}
	err = dingtalk.Dingtalk(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error dingtalk": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"INFO": fmt.Sprint(c.Request, c.Request.Header)})
	return
}