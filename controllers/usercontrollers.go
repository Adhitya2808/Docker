package controllers

import (
	"Praktikum/configs"
	m "Praktikum/helper"
	"Praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsersController(c echo.Context) error {
    var users []model.User
    if err := configs.DB.Find(&users).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users":   users,
    })
}

func GetUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
    var user []model.User
    if err := configs.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get user",
        "user":    user,
    })
}

func CreateUserController(c echo.Context) error {
    user := model.User{}
    c.Bind(&user)
    if err := configs.DB.Create(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create new user",
        "user":    user,
    })
}

func DeleteUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
	var user model.User
    if err := configs.DB.Delete(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
    }
    return c.JSON(http.StatusOK, map[string]string{
        "message": "User deleted",
    })
}

func UpdateUserController(c echo.Context) error {
    var update_user model.User
    if err := c.Bind(&update_user); err != nil{
        return c.JSON(http.StatusBadRequest, model.BaseResponse{
            Status: false,
            Message: "User Not Found",
            Data : nil,
        })
    }
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil{
        return c.JSON(http.StatusBadRequest, model.BaseResponse{
            Status: false,
            Message: "Invalid User ID",
            Data : nil,
        })
    }
    exsitinguser := model.User{}
    err = configs.DB.First(&exsitinguser, userID).Error
    if err != nil {
            return c.JSON(http.StatusInternalServerError, model.BaseResponse{
            Status: false,
            Message: "Failed to Fetch user",
            Data : nil,
        })
    }
    exsitinguser.Name = update_user.Name
    exsitinguser.Email = update_user.Email
    exsitinguser.Password = update_user.Password

    err = configs.DB.Save(&exsitinguser).Error
    if err != nil{
        return c.JSON(http.StatusInternalServerError, model.BaseResponse{
            Status: false,
            Message: "Failed to update user",
            Data : nil,
        })
    }
    return c.JSON(http.StatusOK, model.BaseResponse{
            Status: true,
            Message: "Update Successfully",
            Data : nil,
    })
}

func LoginUserController(c echo.Context) error {
    var user model.User
    c.Bind(&user)
    _, err := m.AuthDB(user.Name, user.Password)
    if err == gorm.ErrRecordNotFound{
        c.JSON(http.StatusBadRequest, model.BaseResponse{
            Status: false,
            Message: "Email/Password Not Found",
            Data : nil,
        })
    } else if err != nil{
        c.JSON(http.StatusInternalServerError, model.BaseResponse{
            Status : false,
            Message : "Failed to Authenticate",
            Data : nil,
        })
    }
	token := m.GenerateJWT(user.ID, user.Name)
    
    var response model.UserResponse
    response.ID = user.ID
    response.Email = user.Email
    response.Name = user.Name
    response.Token = token

    return c.JSON(http.StatusOK, model.BaseResponse{
        Status: true,
        Message: "Success Login",
        Data: response,
    })
}