package routes

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.RouterGroup) {
	Student(r)

	// r.GET("/hehe", func(c *gin.Context) {
	// 	time.Sleep(15 * time.Second)

	// 	fmt.Println("HEHEHEHEHEHEHEHEHEEHE")
	// 	c.JSON(200, gin.H{
	// 		"result": "HEHEHEHEHE",
	// 	})
	// })

	// r.GET("/haha", func(c *gin.Context) {

	// 	fmt.Println("HAHAHAHAHAHAHAHAHAAHA")
	// 	c.JSON(200, gin.H{
	// 		"result": "HAHAHAHAHA",
	// 	})
	// })

}
