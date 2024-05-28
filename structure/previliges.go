package structure

const (
	admin   = "administrator"
	manager = "manager"
	teacher = "teacher"
	parent  = "parent"
	student = "student"
	guest   = "guest"
)

func isAdmin(roleName string) bool {
	return roleName == admin
}

var adminRoles = []string{
	admin,
}

var manageRoles = []string{
	admin,
	manager,
}

var teacherRoles = []string{
	admin,
	manager,
	teacher,
}

var parentRoles = []string{
	admin,
	manager,
	teacher,
	parent,
}

var studentRoles = []string{
	admin,
	manager,
	teacher,
	student,
}
