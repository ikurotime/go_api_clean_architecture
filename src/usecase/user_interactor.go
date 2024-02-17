package usecase

import "ikurotime/go_api_hexagonal_architecture/src/domain"

type UserInteractor struct { 
    UserRepository UserRepository
}
func (interactor *UserInteractor) Add(user domain.User){
    interactor.UserRepository.Store(user)
}
func (interactor *UserInteractor) GetInfo() []domain.User{
    return interactor.UserRepository.Select()
}
func (interactor *UserInteractor) Delete(id string){
    interactor.UserRepository.Delete(id)
}
