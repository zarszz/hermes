package controllers

import (
	"fmt"
	"hermes/controllers/entity"
	"hermes/models/user/repo"
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
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "invalid data"},
		)
		return
	}
	err := s.UserService.Create(input)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "error", "error": err.Error()},
		)
		return
	}
	
	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "registration successfully",
			"data": input,
		},
	)
}

func (s *User) Get(c *gin.Context) {
	users, err := s.UserService.GetAll()
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"users": &users.Users},
	)
}

func (s *User) GetByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := s.UserService.GetByPK(id)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if user == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"user": user},
	)
}

func (s *User) Login(c *gin.Context)  {
	var input entity.UserLoginInput
	if (c.ShouldBindWith(&input, binding.JSON)) != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "invalid data"},
		)
		return
	}
	user, err := s.UserService.GetByEmail(input.Email)
	if user == nil || err != nil {
		if err != nil {
			fmt.Printf("an error occured : %s", err.Error())
		}
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "username or password incorrect"},
		)
		return
	}
	if user.Password != input.Password {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "username or password incorrect"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"data": user, "message": "login success"},
	)
}

func (s *User) Update(c *gin.Context) {
	var form entity.UserUpdateInput
	id := c.Param("id")
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data"},
		)
		c.Abort()
		return
	}
	userId, _ := strconv.Atoi(id)
	err := s.UserService.UpdateOne(form, userId)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"message": "successfully updated"},
	)
}

func (s *User) DeleteByPK(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := s.UserService.GetByPK(id)
	if err != nil {		
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if user == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{ "message": "user not found"},
		)
		return
	}

	errDelete := s.UserService.DeleteOne(user.Id)

	if errDelete != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error()},
		)
		return
	}
	if user == nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"user": user},
	)
}

