package controller

import (
	"fmt"

	"github.com/Kovokar/estudos-go/tree/main/apis/crud-go/src/configuration/rest_err"
	"github.com/Kovokar/estudos-go/tree/main/apis/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Campo passado de forma incorreta, error=%s\n", err.Error()))
		c.JSON(int(restErr.Code), restErr)
		return
	}

	fmt.Println(UserRequest)
}
