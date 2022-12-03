package handlers

import (
	"my-gram-1/accepted_responses"
	"my-gram-1/exceptions"
	"my-gram-1/helpers"
	"my-gram-1/models/entity"
	"my-gram-1/models/web"
	"my-gram-1/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserHandlerImpl struct {
	UserService services.UserService
}

type GormErr struct {
	Number  int    `json:"Number"`
	Message string `json:"Message"`
}

var (
	appJSON = "application/json"
)

func NewUserHandler(userService services.UserService) UserHandler {
	return &UserHandlerImpl{UserService: userService}
}

func (h *UserHandlerImpl) Register(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		exceptions.Errors(ctx, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	newUser, err := h.UserService.Register(&user)

	if err != nil {
		exceptions.Errors(ctx, http.StatusBadRequest, "Failed Register Input", err.Error())
		return
	}
	userResp := web.RegisterResponse(newUser)
	accepted_responses.SuccessResponse(ctx, 201, "Success Register User", userResp)
}

func (h *UserHandlerImpl) Login(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		exceptions.Errors(ctx, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	loginUser, err := h.UserService.Login(user)

	if err != nil {
		exceptions.Errors(ctx, http.StatusNotFound, "Failede User Not Found", err.Error())
		return
	}

	validPass := helpers.ComparePass([]byte(loginUser.Password), []byte(user.Password))
	if !validPass {
		exceptions.Errors(ctx, http.StatusUnauthorized, "Password Failed", "Unauthenthicated")
		return
	}

	genToken := helpers.GenerateToken(loginUser.ID, loginUser.Email)
	accepted_responses.SuccessResponse(ctx, http.StatusOK, "Login Success", genToken)
}

func (h *UserHandlerImpl) Profile(ctx *gin.Context) {
	var user entity.User
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userEmail := string(userData["email"].(string))
	user = entity.User{
		Email: userEmail,
	}
	getUser, err := h.UserService.Profile(user)

	if err != nil {
		exceptions.Errors(ctx, http.StatusNotFound, "Profile not found", err.Error())
	}

	userResp := web.ProfileResponse(getUser)
	accepted_responses.SuccessResponse(ctx, http.StatusOK, "Success Profile", userResp)
}

func (h *UserHandlerImpl) Update(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBind(&user); err != nil {
		exceptions.Errors(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	user.ID = userID

	updateUser, err := h.UserService.Update(&user)

	if err != nil {
		exceptions.Errors(ctx, http.StatusBadRequest, "Failed Update User", err.Error())
	}

	userResp := web.UpdateResponse(updateUser)
	accepted_responses.SuccessResponse(ctx, http.StatusOK, "Update User Success", userResp)
}

func (h *UserHandlerImpl) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	err := h.UserService.Delete(userID)

	if err != nil {
		exceptions.Errors(ctx, http.StatusNotFound, "Failed Delete User", "Failed Delete User")
	}

	accepted_responses.SuccessResponse(ctx, http.StatusAccepted, "Delete User Success", "Delete User Success")
}
