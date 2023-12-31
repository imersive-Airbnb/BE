package handler

import (
	"airnbn/app/middlewares"
	"airnbn/features/user"
	"airnbn/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
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
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind Data"))
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
		"name":    userData.Name,
		"email":   userData.Email,
		"token":   token,
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Success", response))
}

func (handler *UserHandler) CheckProfileByID(c echo.Context) error {
	// Extract user ID from the path parameter
	userID := c.Param("id")

	// Parse userID to uuid.UUID
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid User ID"))
	}

	// Retrieve user profile from the userService
	userData, err := handler.userService.CheckProfile(uuidUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to Retrieve User Profile"))
	}

	// Create the UserResponse
	response := UserResponse{
		UserID:    userData.UserID,
		Name:      userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	// Return the response
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("User Profile Retrieved Successfully", response))
}

func (handler *UserHandler) UpdateUserByID(c echo.Context) error {
	// Get the user ID from the path parameter
	userID := c.Param("id")

	// Get the updated user data from the request
	var updatedUser user.UserCore
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to Bind Data"))
	}

	if updatedUser.Password != "" {
		hashedPassword, err := helper.HashPassword(updatedUser.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
		}
		updatedUser.Password = hashedPassword
	}

	if err := handler.userService.UpdateUser(userID, updatedUser); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Updated Successfully"))

}

func (handler *UserHandler) UpgradeStatus(c echo.Context) error {
	// Mendapatkan ID pengguna dari token JWT yang valid
	userID, err := middlewares.ExtractTokenUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Unauthorized"))
	}

	newStatus := c.FormValue("status")

	// Validasi status baru
	if newStatus != "default" && newStatus != "hosting" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid status"))
	}

	// Upgrade status pengguna
	err = handler.userService.UpgradeStatus(strconv.Itoa(userID), newStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Status upgraded successfully"))
}
