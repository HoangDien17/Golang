package enum

type PermissionEnum string

const (
	READ   PermissionEnum = "READ"
	WRITE  PermissionEnum = "WRITE"
	UPDATE PermissionEnum = "UPDATE"
	DELETE PermissionEnum = "DELETE"
)
