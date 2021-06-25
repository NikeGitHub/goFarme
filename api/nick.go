package api

import (
	"code.shomes.cn/tfblue/smartgy/api/api"
	"github.com/gin-gonic/gin"
)

func Nick(r *gin.Engine) {
	// 下拉列表
	r.POST("/smart/select-option/project-use", api.SelectOption.ProjectUse)
	r.POST("/smart/select-option/project-type", api.SelectOption.ProjectType)
	r.POST("/smart/select-option/schedule-type", api.SelectOption.ScheduleType)
	r.POST("/smart/select-option/structure-type", api.SelectOption.StructureType)
}
