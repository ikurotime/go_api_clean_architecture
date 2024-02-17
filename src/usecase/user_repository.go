package usecase

import "ikurotime/go_api_clean_architecture/src/domain"

type UserRepository interface {
    Store(domain.User)
    Select() []domain.User
    Delete(id string)
}
