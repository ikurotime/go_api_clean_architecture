package database

type SqlHandler interface {
    Create(obj interface{})
    FindAll(obj interface{})
    DeleteById(obj interface{}, id string)
}
