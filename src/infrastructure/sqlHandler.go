package infrastructure

import (
	"ikurotime/go_api_clean_architecture/src/interfaces/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
    db *gorm.DB
}

func NewSqlHandler(conn string) database.SqlHandler{
    db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }

    sqlHandler := new(SqlHandler)
    sqlHandler.db = db

    return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}){
    handler.db.Create(obj)
}
func (handler *SqlHandler) FindAll(obj interface{}){
    handler.db.Find(obj)
}
func (handler *SqlHandler) DeleteById(obj interface{}, id string){
    handler.db.Delete(obj,id)
}
