package main

import (
	"arm-test/proto"
	"arm-test/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func predict(c *gin.Context) {
	var predictReq proto.PredictionRequest
	if err := c.ShouldBindJSON(&predictReq); err != nil {
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}

	srv := &service.ModelService{}
	predictRsp, err := srv.PredictV2(c.Request.Context(), &predictReq)
	if err != nil {
		zap.S().Error("srv.PredictV2", "err", err)
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, predictRsp)
}

func main() {
	r := gin.Default()
	r.POST("/v1/predict", predict)
	r.Run() //监听并在 0.0.0.0:8080 上启动服务
}
