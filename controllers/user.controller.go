package controller

import (
	db "employee/connect"
	service "employee/services"
	"net/http"

	form "employee/app/interface"
	"employee/utils"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Tags Users
// @Summary Create user
// @Description Create user
// @Accept  json
// @Produce  json
// @Param createUser body UserLoginDTO true "Information of user"
// @Success 200 {object} User
// @BasePath /api/v1
// @Router /signup [post]
func CreateUser(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		userEntity := service.NewUserEntity(resource)
		userRequest := form.UserLogin{}
		if err := c.Bind(&userRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		user, code, err := userEntity.CreateOne(userRequest)
		response := map[string]interface{}{
			"user":  user,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// FindUserById godoc
// @Tags Users
// @Summary Find user by id
// @Description Find user by id
// @Accept  json
// @Produce  json
// @Param id path string true "Id of user"
// @Success 200 {object} User
// @BasePath /api/v1
// @Router /user/{id} [get]
// @security ApiKeyAuth
func FindUserById(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		userEntity := service.NewUserEntity(resource)
		id := c.Param("id")
		user, code, err := userEntity.FindUserById(id)
		response := map[string]interface{}{
			"user":  user,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// FindAllUser godoc
// @Tags Users
// @Summary Find all user
// @Description Find all user
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @BasePath /api/v1
// @Router /user [get]
// @security ApiKeyAuth
func FindAllUser(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		userEntity := service.NewUserEntity(resource)
		users, code, err := userEntity.FindAllUser()
		response := map[string]interface{}{
			"users": users,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// DeleteUserById godoc
// @Tags Users
// @Summary Delete user by id
// @Description Delete user by id
// @Accept  json
// @Produce  json
// @Param id path string true "Id of user"
// @Success 200 {object} DeleteResponse
// @BasePath /api/v1
// @Router /user/{id} [delete]
// @security ApiKeyAuth
func DeleteUserById(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		userEntity := service.NewUserEntity(resource)
		id := c.Param("id")
		code, _ := userEntity.DeleteUserById(id)
		if code != 200 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Delete user failed"})
			return
		}
		c.JSON(code, gin.H{"message": "Delete user successfully"})
	}
}

// Login godoc
// @Tags Users
// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param login body UserDTO true "Information of user"
// @Success 200 {object} LoginResponse
// @BasePath /api/v1
// @Router /login [post]
func Login(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		userEntity := service.NewUserEntity(resource)
		loginRequest := form.User{}
		if err := c.Bind(&loginRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		user, code, err := userEntity.Login(loginRequest)
		response := map[string]interface{}{
			"user":  user,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}
