package database

import "ikurotime/go_api_clean_architecture/src/domain"

type UserRepository struct {
    SqlHandler
}

func (db *UserRepository) Store(user domain.User){
    db.Create(&user)
}

func (db *UserRepository) Select() []domain.User{
    user := []domain.User{}
    db.FindAll(&user)
    return user
}

func (db *UserRepository) Delete(id string){
    user := domain.User{}
    db.DeleteById(&user, id)
}
