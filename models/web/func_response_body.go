package web

import "my-gram-1/models/entity"

func RegisterResponse(user *entity.User) UserBodyResponse {
	return UserBodyResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func LoginResponse(token string) LoginBodyResponse {
	return LoginBodyResponse{
		Token: token,
	}
}

func ProfileResponse(user entity.User) UserBodyResponse {
	return UserBodyResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}
}

func UpdateResponse(user *entity.User) UpdateUserBodyResponse {
	return UpdateUserBodyResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: *user.UpdatedAt,
	}
}
