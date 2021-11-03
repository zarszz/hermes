package controllers

import (
	"fmt"
	"hermes/controllers/entity"
	"hermes/models/user/repo"
	"hermes/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmoiron/sqlx"
)

type User struct {
	UserService *repo.UserService
}

func NewUserController(conn *sqlx.DB) *User {
	return &User{UserService: repo.NewUserService(conn)}
}

func (s *User) Register(c *gin.Context) {
	var input entity.UserRegistrationInput
	if c.ShouldBindWith(&input, binding.JSON) != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "invalid data")
		return
	}
	err := s.UserService.Create(input)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "registration successfully", http.StatusOK, c.Request.Method, "")
}

func (s *User) Get(c *gin.Context) {
	users, err := s.UserService.GetAll()
	if err != nil {		
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "success retrieve data", http.StatusOK, c.Request.Method, users.Users)
}

func (s *User) GetByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := s.UserService.GetByPK(id)
	if err != nil {		
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	if user == nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "user not found")
		return
	}
	utils.HandleResponse(c, "successfully retrieve data", http.StatusOK, c.Request.Method, user)
}

func (s *User) Login(c *gin.Context)  {
	var input entity.UserLoginInput
	if (c.ShouldBindWith(&input, binding.JSON)) != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "invalid data")
		return
	}
	user, err := s.UserService.GetByEmail(input.Email)
	if user == nil || err != nil {
		if err != nil {
			fmt.Printf("an error occured : %s", err.Error())
		}
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "username or password incorrect")
		return
	}
	if user.Password != input.Password {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "username or password incorrect")
		return
	}
	utils.HandleResponse(c, "login success", http.StatusOK, c.Request.Method, user)
}

func (s *User) Update(c *gin.Context) {
	var form entity.UserUpdateInput
	id := c.Param("id")
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "invalid data")
		return
	}
	userId, _ := strconv.Atoi(id)
	err := s.UserService.UpdateOne(form, userId)
	if err != nil {		
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "update success", http.StatusOK, c.Request.Method, "")
}

func (s *User) DeleteByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := s.UserService.GetByPK(id)
	if err != nil {		
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	if user == nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, "user not found")
		return
	}

	errDelete := s.UserService.DeleteOne(user.Id)

	if errDelete != nil {
		utils.HandleErrorResponse(c, http.StatusBadRequest, c.Request.Method, err)
		return
	}
	utils.HandleResponse(c, "delete success", http.StatusOK, c.Request.Method, "")
}

