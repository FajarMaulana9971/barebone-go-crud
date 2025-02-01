package request

import (
	"barebone-go-crud/src/models/dto/request"
	"barebone-go-crud/src/models/entity"
)

func MapUserRequestToUserModel(req *request.UserRequest) *entity.User {
	return &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
}
