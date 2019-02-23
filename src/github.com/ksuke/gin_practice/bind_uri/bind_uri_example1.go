package bind_uri

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

// call BindUriServer func in main.go
// terminal
// $ go run main.go

// test it with:
//$ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//$ curl -v localhost:8088/thinkerou/not-uuid
func BindUriServer() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"name": person.Name, "uuid": person.ID})
		}
	})
	route.Run(":8080")
}
