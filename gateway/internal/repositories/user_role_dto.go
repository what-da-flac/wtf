package repositories

type UserRoleDto struct {
	RoleId string
	UserId string
}

func (x *UserRoleDto) TableName() string { return "user_role" }
