package infrastructure

import (
	controllers "ikurotime/go_api_clean_architecture/src/interfaces/api"
	"net/http"

	"github.com/labstack/echo/v4"
)


func Init(e *echo.Echo,conn string){

    userController := controllers.NewUserController(NewSqlHandler(conn))
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK,"Hello World")
    })

    e.GET("/users", func(c echo.Context) error {
        users := userController.GetUser()
        c.Bind(&users)
        return c.JSON(http.StatusOK, users)
    })

    e.POST("/users", func(c echo.Context) error {
        userController.Create(c)
        return c.String(http.StatusOK, "created")
    })

    e.DELETE("/users/:id", func(c echo.Context) error {
        id := c.Param("id")
        userController.Delete(id)
        return c.String(http.StatusOK, "deleted")
    })
} 
