package structure

type Group struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type Role struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
