package auth

const (
	PermissionCreateUser = "create_user"
	PermissionUpdateUser = "update_user"
	PermissionDeleteUser = "delete_user"

	PermissionCreateLesson = "create_plan"
	PermissionUpdateLesson = "update_plan"
	PermissionDeleteLesson = "delete_plan"
)

var RolePermissions = map[string][]string{
	"admin": {PermissionCreateUser, PermissionUpdateUser, PermissionDeleteUser, PermissionCreateLesson, PermissionUpdateLesson, PermissionDeleteLesson},
	"user":  {PermissionCreateLesson, PermissionUpdateLesson, PermissionDeleteLesson},
}
