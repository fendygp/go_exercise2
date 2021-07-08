package usecase

import (
	"bitbucket.org/bridce/go_exercise2/repo"
)

//UsersUsecaseStruct class for define business logic
type UsersUsecaseStruct struct {
	users repo.UsersInterface
}

type UsersInterface interface {
	CreateUsers() error
}

func CreateGetTokenUsecaseImp(users repo.UsersInterface) UsersInterface {
	return &UsersUsecaseStruct{users}
}

func (g *UsersUsecaseStruct) CreateUsers() error {

	//your business logic define here

	return nil
}
