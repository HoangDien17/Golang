package controller

import (
	db "employee/connect"
	service "employee/services"
	"net/http"

	form "employee/app/interface"
	"employee/utils"

	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Tags Roles
// @Summary Create role
// @Description Create role
// @Accept  json
// @Produce  json
// @Param createRole body RoleDTO true "Information of role"
// @Success 200 {object} Role
// @BasePath /api/v1
// @Router /role [post]
// @security ApiKeyAuth
func CreateRole(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		roleEntity := service.NewRoleEntity(resource)
		roleRequest := form.Role{}
		if err := c.Bind(&roleRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		role, code, err := roleEntity.CreateRole(roleRequest)
		response := map[string]interface{}{
			"role":  role,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// FindRoleById godoc
// @Tags Roles
// @Summary Find role by id
// @Description Find role by id
// @Accept  json
// @Produce  json
// @Param id path string true "Id of role"
// @Success 200 {object} Role
// @BasePath /api/v1
// @Router /role/{id} [get]
// @security ApiKeyAuth
func FindRoleById(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		roleEntity := service.NewRoleEntity(resource)
		id := c.Param("id")
		role, code, err := roleEntity.FindRoleById(id)
		response := map[string]interface{}{
			"role":  role,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// FindAllRoles godoc
// @Tags Roles
// @Summary Find all roles
// @Description Find all roles
// @Accept  json
// @Produce  json
// @Success 200 {array} Role
// @BasePath /api/v1
// @Router /role/all [get]
// @security ApiKeyAuth
func FindAllRoles(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		roleEntity := service.NewRoleEntity(resource)
		roles, code, err := roleEntity.FindAllRoles()
		response := map[string]interface{}{
			"roles": roles,
			"error": utils.GetErrorMessage(err),
		}
		c.JSON(code, response)
	}
}

// DeleteRoleById godoc
// @Tags Roles
// @Summary Delete role by id
// @Description Delete role by id
// @Accept  json
// @Produce  json
// @Param id path string true "Id of role"
// @Success 200 {object} Role
// @BasePath /api/v1
// @Router /role/{id} [delete]
// @security ApiKeyAuth
func DeleteRoleById(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		roleEntity := service.NewRoleEntity(resource)
		id := c.Param("id")
		code, err := roleEntity.DeleteRoleById(id)
		response := make(map[string]interface{})
		if err != nil {
			response["error"] = utils.GetErrorMessage(err)
		} else {
			response["message"] = "Delete role successfully"
		}
		c.JSON(code, response)
	}
}

// UpdatePermissionByAdmin godoc
// @Tags Roles
// @Summary Update permission by admin
// @Description Update permission by admin
// @Accept  json
// @Produce  json
// @Param id path string true "Id of role"
// @Param updatePermission body PermissionDTO true "Information of permission"
// @Success 200 {object} Role
// @BasePath /api/v1
// @Router /role/{id}/permission [put]
// @security ApiKeyAuth
func UpdatePermissionByAdmin(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		roleEntity := service.NewRoleEntity(resource)
		id := c.Param("id")
		permissionRequest := form.Permissions{}
		if err := c.Bind(&permissionRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		code, err := roleEntity.UpdatePermissionByAdmin(id, permissionRequest)
		response := make(map[string]interface{})
		if err != nil {
			response["error"] = utils.GetErrorMessage(err)
		} else {
			response["message"] = "Update role successfully"
		}
		c.JSON(code, response)
	}
}
