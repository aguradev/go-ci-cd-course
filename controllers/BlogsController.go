package controllers

import (
	"net/http"
	"praktikum/configs"
	"praktikum/models"
	"praktikum/resources"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllBlogController(c echo.Context) error {

	var GetBlog []models.Blogs

	if BlogException := configs.DB.Preload("Users").Find(&GetBlog); BlogException.Error != nil {

		if BlogException.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNoContent, resources.BaseResponse{
				Status:  http.StatusNoContent,
				Message: "Data Not Found",
				Data:    nil,
			})
		}

	}

	MappingResponse := resources.MappingBlogResponse(GetBlog)

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "Data Not Found",
		Data:    MappingResponse,
	})
}

func CreateBlogController(c echo.Context) error {

	var GetBlog models.BlogRequest
	var GetUsers models.Users
	var DataBlog models.Blogs

	if getData := c.Bind(&GetBlog); getData != nil {
		return c.JSON(http.StatusInternalServerError, getData.Error)
	}

	CheckUsersExists := configs.DB.Where("id = ?", GetBlog.UserID).Find(&GetUsers)

	if CheckUsersExists.Error != nil {
		if CheckUsersExists.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, resources.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Data Users Not Found",
				Data:    nil,
			})
		}
	} else if CheckUsersExists.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Data Users Not Found",
			Data:    nil,
		})
	}

	DataBlog = DataBlog.RequestBlog(GetBlog)

	if CreateBlog := configs.DB.Create(&DataBlog); CreateBlog.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Create Book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusAccepted,
		Message: "Success Create Blog",
		Data:    DataBlog,
	})

}

func GetBlogByIdController(c echo.Context) error {

	var (
		GetBlog      models.Blogs
		BlogResponse resources.BlogResponse
	)

	GetId, ErrId := strconv.Atoi(c.Param("id"))

	if ErrId != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id is not number",
			Data:    nil,
		})
	}

	if BlogExp := configs.DB.Preload("Users").Where("id = ?", GetId).First(&GetBlog); BlogExp.Error != nil {

		if BlogExp.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, resources.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Blog Not Found",
				Data:    nil,
			})

		} else {
			return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error When Searching Blog",
				Data:    nil,
			})
		}

	}

	BlogResponse.DetailBlogResponse(GetBlog)

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success show detail blog",
		Data:    BlogResponse,
	})

}

func DeleteBlogController(c echo.Context) error {

	GetId, ErrId := strconv.Atoi(c.Param("id"))

	if ErrId != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id is not number",
			Data:    nil,
		})
	}

	Deleted := configs.DB.Delete(&models.Blogs{}, GetId)

	if Deleted.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error when delete blog",
			Data:    nil,
		})
	}

	if Deleted.RowsAffected == 0 {
		return c.JSON(http.StatusNoContent, resources.BaseResponse{
			Status:  http.StatusNoContent,
			Message: "Error when delete book",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusNoContent,
		Message: "Blog Deleted!",
		Data:    nil,
	})
}

func UpdateBlogController(c echo.Context) error {

	var GetBlog models.BlogRequest
	var DataBlog models.Blogs

	GetId, ErrId := strconv.Atoi(c.Param("id"))

	if ErrId != nil {
		return c.JSON(http.StatusBadRequest, resources.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "Id is not number",
			Data:    nil,
		})
	}

	if getData := c.Bind(&GetBlog); getData != nil {
		return c.JSON(http.StatusInternalServerError, getData.Error)
	}

	DataBlog = DataBlog.RequestBlog(GetBlog)

	if UpdateBlog := configs.DB.Where("id = ?", GetId).Updates(&DataBlog); UpdateBlog.Error != nil {
		return c.JSON(http.StatusInternalServerError, resources.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error When Update Book",
			Data:    nil,
		})
	}

	configs.DB.Where("id = ?", GetId).Find(&DataBlog)

	return c.JSON(http.StatusCreated, resources.BaseResponse{
		Status:  http.StatusAccepted,
		Message: "Blog Updated",
		Data:    DataBlog,
	})

}
