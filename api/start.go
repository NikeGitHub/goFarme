package api

import (
	bsGin "code.shomes.cn/pkg/bootstrap/gin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start() {
	noAuth := func(r *gin.Engine) {
		Nick(r)
	}
	auth := func(r *gin.Engine) {
	}
	if err := bsGin.Start2(noAuth, auth); err != nil {
		fmt.Println(err)
	}
}
