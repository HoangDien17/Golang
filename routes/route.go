package route

import (
	constant "employee/app/constants"
	db "employee/connect"
	controller "employee/controllers"
	middleware "employee/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup, resource *db.Resource) {
	// public route
	r.POST("/signup", controller.CreateUser(resource))
	r.POST("/login", controller.Login(resource))

	// private route
	user := r.Group("/user")
	{
		user.GET("/:id", middleware.Authorization(resource, []string{constant.ADMIN}), controller.FindUserById(resource))
		user.GET("/", middleware.Authorization(resource, []string{constant.ADMIN}), controller.FindAllUser(resource))
		user.DELETE("/:id", middleware.Authorization(resource, []string{constant.ADMIN}), controller.DeleteUserById(resource))
	}
}

func RoleRoute(r *gin.RouterGroup, resource *db.Resource) {
	// private route
	role := r.Group("/role")
	role.POST("/", middleware.Authorization(resource, []string{constant.ADMIN}), controller.CreateRole(resource))
	role.GET("/:id", middleware.Authorization(resource, []string{constant.ADMIN}), controller.FindRoleById(resource))
	role.GET("/all", middleware.Authorization(resource, []string{constant.ADMIN}), controller.FindAllRoles(resource))
	role.DELETE("/:id", middleware.Authorization(resource, []string{constant.ADMIN}), controller.DeleteRoleById(resource))
	role.PUT("/:id/permission", middleware.Authorization(resource, []string{constant.ADMIN}), controller.UpdatePermissionByAdmin(resource))
}
