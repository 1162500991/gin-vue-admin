package api

import (
	"QMPlusCommon/controller/servers"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/model/dbModel"
)

// @Tags WebExample
// @Summary 创建基础api
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExampleModule true "多服务模式示例"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exp/createExp [post]
func CreateExp(c *gin.Context) {
	var exm dbModel.ExampleModule
	_ = c.BindJSON(&exm)
	err := exm.CreateExampleModule()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("创建失败：%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "创建成功", gin.H{})
	}
}
