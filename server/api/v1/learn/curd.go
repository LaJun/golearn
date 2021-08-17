package learn

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"gin-vue-admin/model/learn"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LearnApi struct {
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createApi [post]
func (s * LearnApi) CreateApi(c *gin.Context) {
	var data learn.Curd
	_ = c.ShouldBindJSON(&data)
	if err := utils.Verify(data, utils.CurdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := learnService.Add(data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}