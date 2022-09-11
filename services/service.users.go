package services

import (
	"github.com/restuwahyu13/go-trakteer-api/dtos"
	"github.com/restuwahyu13/go-trakteer-api/helpers"
	"github.com/restuwahyu13/go-trakteer-api/repositorys"
)

type UsersService struct {
	repository *repositorys.UsersRepository
}

func NewUsersService(repository *repositorys.UsersRepository) *UsersService {
	return &UsersService{repository: repository}
}

func (ctx *UsersService) LoginService(payload dtos.DTOLogin) helpers.APIResponse {
	return ctx.repository.LoginRepository(payload)
}
