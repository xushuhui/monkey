package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func setup() {

	gin.SetMode(gin.TestMode)
	//fmt.Println(global.JWTSetting.Secret)
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}
