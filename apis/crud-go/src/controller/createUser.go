package controller

import (
	"fmt"
	"log"

	"github.com/Kovokar/estudos-go/tree/main/apis/crud-go/src/configuration/rest_err/validation"
	"github.com/Kovokar/estudos-go/tree/main/apis/crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	log.Println("Init CreateUser controler")
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(int(restErr.Code), restErr)
		return
	}

	fmt.Println(UserRequest)
}
