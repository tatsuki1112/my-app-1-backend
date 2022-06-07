package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tatsuki1112/my-app-1-backend/db"
	"github.com/tatsuki1112/my-app-1-backend/entity"
	"net/http"
)

type CreateTrashUserForm struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateTrashUserForm struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db.Init()
	r := gin.Default()
	r.GET("/app/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "こんにちはｗｗｗｗ",
		})
	})
	r.GET("app/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "my name is Tatsuki",
		})
	})

	r.POST("app/trash-user", func(c *gin.Context) {
		var form CreateTrashUserForm
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "BadRequest: ",
			})
			return
		}
		trashUser := entity.TrashUser{
			Name: form.Name,
		}

		db.GetDB().Create(&trashUser)
		c.JSON(http.StatusOK, trashUser)

	})
	r.GET("app/trash-users", func(c *gin.Context) {
		trashUsers := []entity.TrashUser{}
		db.GetDB().Find(&trashUsers)

		c.JSON(http.StatusOK, trashUsers)
	})
	r.GET("app/trash-user/:id", func(c *gin.Context) {
		trashUser := entity.TrashUser{}
		id := c.Param("id")
		db.GetDB().First(&trashUser, id)
		c.JSON(http.StatusOK, trashUser)
	})

	r.PUT("app/trash-user/:id", func(c *gin.Context) {
		var form UpdateTrashUserForm
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "BadRequest: ",
			})
			return
		}
		trashUser := entity.TrashUser{}
		id := c.Param("id")
		data := entity.TrashUser{
			Name: form.Name,
		}
		db.GetDB().First(&trashUser, id).Updates(&data)
		c.JSON(http.StatusOK, trashUser)
	})

	r.DELETE("app/trash-user/:id", func(c *gin.Context) {
		trashUser := entity.TrashUser{}
		id := c.Param("id")
		db.GetDB().Delete(&trashUser, id)
		c.JSON(http.StatusOK, trashUser)
	})

	r.Run()
}
