package form

type Role struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
} //@name RoleDTO

type Permissions struct {
	Permissions []string `json:"permissions" binding:"required"`
} // @name PermissionDTO
