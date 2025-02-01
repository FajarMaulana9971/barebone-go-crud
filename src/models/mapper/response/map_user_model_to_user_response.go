package response

import (
	"barebone-go-crud/src/models/dto/response"
	"barebone-go-crud/src/models/entity"
)

func MapUserModelToUserResponse(user *entity.User) *response.UserResponse {
	return &response.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
