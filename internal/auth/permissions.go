package auth

const (
	PermissionCreateUser = "create_user"
	PermissionUpdateUser = "update_user"
	PermissionDeleteUser = "delete_user"

	PermissionCreateLesson = "create_lesson"
	PermissionUpdateLesson = "update_lesson"
	PermissionDeleteLesson = "delete_lesson"
)

var RolePermissions = map[string][]string{
	"admin": {PermissionCreateUser, PermissionUpdateUser, PermissionDeleteUser, PermissionCreateLesson, PermissionUpdateLesson, PermissionDeleteLesson},
	"user":  {PermissionCreateLesson, PermissionUpdateLesson, PermissionDeleteLesson},
}
