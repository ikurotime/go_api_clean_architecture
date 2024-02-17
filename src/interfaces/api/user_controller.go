package controllers

import (
	"ikurotime/go_api_clean_architecture/src/domain"
	"ikurotime/go_api_clean_architecture/src/interfaces/database"
	"ikurotime/go_api_clean_architecture/src/usecase"

	"github.com/labstack/echo/v4"
)



type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
} 

func (controller *UserController) Create(c echo.Context){
    user := domain.User{}
    c.Bind(&user)
    controller.Interactor.Add(user)
    createdUsers := controller.Interactor.GetInfo()
    c.JSON(201, createdUsers)
}
func (controller *UserController) GetUser() []domain.User{
    return controller.Interactor.GetInfo()
}
func (controller *UserController) Delete(id string){
    controller.Interactor.Delete(id)
}
