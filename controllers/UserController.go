package controllers

import (
	"net/http"
	"praktikum/configs"
	"praktikum/models"
	users "praktikum/models/Users"
	"praktikum/resources"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	var GetAllUsers []models.Users
	var CountUsers int64

	UsersException := configs.DB.Model(&models.Users{}).Preload("Blogs").Find(&GetAllUsers)

	if UsersException.Error != nil {
		return c.JSON(http.StatusInternalServerError, UsersException.Error)
	}

	CountUsersException := configs.DB.Model(&models.Users{}).Count(&CountUsers)

	if CountUsersException.Error != nil {
		return c.JSON(http.StatusInternalServerError, CountUsersException.Error)
	}

	if CountUsers == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "Users is empty",
			Data:    nil,
		})
	}

	UsersResource := resources.MappingAllDataUsers(GetAllUsers)

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "success get all users",
		Data:    UsersResource,
	})

}

func GetUserByIdController(c echo.Context) error {

	getId, IdException := strconv.Atoi(c.Param("id"))

	var (
		GetUser      models.Users
		UserResponse resources.UserResponseDetail
	)

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	UserException := configs.DB.Where("id = ?", getId).First(&GetUser)

	UserResponse.GetDetailUserResponse(GetUser)

	if UserException.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Searching User",
			Data:    nil,
		})
	}

	if UserException.RowsAffected == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "User Not Found",
			Data:    UserException.Error,
		})
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "User Detail",
		Data:    UserResponse,
	})

}

func CreateUserController(c echo.Context) error {

	var (
		GetAllRequestData users.UserCreateRequest
		UsersModel        models.Users
	)

	BindingException := c.Bind(&GetAllRequestData)

	if BindingException != nil {
		return c.JSON(http.StatusInternalServerError, BindingException.Error())
	}

	UsersModel = UsersModel.CreateFormData(GetAllRequestData)

	CreateException := configs.DB.Create(&UsersModel)

	if CreateException.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Create User",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusCreated,
		Message: "User Created Success!",
		Data:    UsersModel,
	})

}

func DeleteUserController(c echo.Context) error {

	getId, IdException := strconv.Atoi(c.Param("id"))

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	Deleted := configs.DB.Delete(&models.Users{}, getId)

	if Deleted.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error when delete user",
			Data:    nil,
		})
	}

	if Deleted.RowsAffected == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "User not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusNoContent,
		Message: "User Deleted!",
		Data:    nil,
	})

}

func UpdateUserController(c echo.Context) error {

	var (
		GetAllRequestData users.UserCreateRequest
		UsersModel        models.Users
	)

	getId, IdException := strconv.Atoi(c.Param("id"))
	RequestBindingException := c.Bind(&GetAllRequestData)

	if RequestBindingException != nil {
		return c.JSON(http.StatusInternalServerError, RequestBindingException.Error())
	}

	if IdException != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id Is Not Number",
			Data:    nil,
		})
	}

	UsersModel = UsersModel.CreateFormData(GetAllRequestData)

	updatedException := configs.DB.Where("id = ?", getId).Updates(&UsersModel)

	if updatedException.Error != nil {
		return c.JSON(http.StatusInternalServerError, updatedException.Error)
	}

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusCreated,
		Message: "User update success",
		Data:    UsersModel,
	})

}
