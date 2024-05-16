package main

import (
	"fmt"
	"gosplitwise/app/api/routes"
	"gosplitwise/app/pkg/mongodbmdl"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("gosplitwise")
	mongodbmdl.Init()
	//router
	g := gin.Default()
	routes.Init(g)

	g.Run(":9000")
}
