package api

import (
	"code.shomes.cn/pkg/bootstrap/tool"
	"code.shomes.cn/tfblue/smartgy/protobuf"
	"errors"
	"github.com/gin-gonic/gin"
)

type selectOption struct{}

var SelectOption = &selectOption{}

func (this *selectOption) ProjectUse(ginC *gin.Context) {
	c := tool.NewGRpcGinClient(ginC)
	defer c.Close()
	request := protobuf.SelectOptionProjectUseIn{}
	if err := ginC.ShouldBindJSON(&request); err != nil {
		c.Error(errors.New("请求参数错误"))
		return
	}
	client := protobuf.NewSelectOptionClient(c.Conn)
	result, err := client.ProjectUse(c.Ctx, &request)
	if err != nil {
		c.Error(err)
		return
	}
	c.Ok(result)
}

func (this *selectOption) ProjectType(ginC *gin.Context) {
	c := tool.NewGRpcGinClient(ginC)
	defer c.Close()
	request := protobuf.SelectOptionProjectTypeIn{}
	if err := ginC.ShouldBindJSON(&request); err != nil {
		c.Error(errors.New("请求参数错误"))
		return
	}
	client := protobuf.NewSelectOptionClient(c.Conn)
	result, err := client.ProjectType(c.Ctx, &request)
	if err != nil {
		c.Error(err)
		return
	}
	c.Ok(result)
}

func (this *selectOption) ScheduleType(ginC *gin.Context) {
	c := tool.NewGRpcGinClient(ginC)
	defer c.Close()
	request := protobuf.SelectOptionScheduleTypeIn{}
	if err := ginC.ShouldBindJSON(&request); err != nil {
		c.Error(errors.New("请求参数错误"))
		return
	}
	client := protobuf.NewSelectOptionClient(c.Conn)
	result, err := client.ScheduleType(c.Ctx, &request)
	if err != nil {
		c.Error(err)
		return
	}
	c.Ok(result)
}

func (this *selectOption) StructureType(ginC *gin.Context) {
	c := tool.NewGRpcGinClient(ginC)
	defer c.Close()
	request := protobuf.SelectOptionStructureTypeIn{}
	if err := ginC.ShouldBindJSON(&request); err != nil {
		c.Error(errors.New("请求参数错误"))
		return
	}
	client := protobuf.NewSelectOptionClient(c.Conn)
	result, err := client.StructureType(c.Ctx, &request)
	if err != nil {
		c.Error(err)
		return
	}
	c.Ok(result)
}
