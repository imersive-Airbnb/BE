package handler

import (
	"airnbn/features/user"
	"airnbn/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(handler user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: handler,
	}
}
func (handler *UserHandler) RegisterUser(c echo.Context) error {
	// Mendekode data JSON dari body permintaan
	var newUser user.UserCore
	err := c.Bind(&newUser)
	if err != nil {
		response := helper.FailedResponse("Failed To Decode Request Body")
		return c.JSON(http.StatusBadRequest, response)
	}

	// Panggil fungsi Create dari UserService untuk membuat pengguna baru
	err = handler.userService.Create(newUser)
	if err != nil {
		response := helper.FailedResponse(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.SuccessResponse("User Created Successfully")
	return c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) Login(c echo.Context) error {
	// Memeriksa apakah email dan password inputan dapat di bind
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// Memeriksa apakah email & password telah diinputkan di database
	userData, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
	}

	response := map[string]interface{}{
		"user_id": userData.UserID,
		"email":   userData.Email,
		"token":   token,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login success", response))
}
