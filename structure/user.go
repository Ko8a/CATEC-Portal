package structure

type User struct {
	Id         int     `json:"-" db:"id"`
	Name       string  `json:"name" binding: "required"`
	Surname    string  `json:"surname" binding: "required"`
	Age        int     `json:"age"`
	Email      string  `json:"email" binding: "required"`
	Password   string  `json:"password" binding: "required"`
	Phone      int     `json:"phone"`
	GroupId    *uint64 `json:"group_id" db:"group_id"`
	TimeUpdate string  `json:"time_update"`
	RoleId     int     `json:"role_id" db:"role_id"`
}

type UserInfo struct {
	Id      int    `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Group   string `json:"group" db:"group"`
}

type UserFullInfo struct {
	Id         *int    `json:"id" db:"id"`
	Name       *string `json:"name" db:"name"`
	Surname    *string `json:"surname" db:"surname"`
	Age        *int    `json:"age" db:"age"`
	Email      *string `json:"email" db:"email"`
	Phone      *int    `json:"phone" db:"phone"`
	GroupId    *uint64 `json:"group_id" db:"group_id"`
	TimeUpdate *string `json:"time_update" db:"time_update"`
	RoleId     *int    `json:"role_id" db:"role_id"`
}
