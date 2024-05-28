package structure

type Mark struct {
	Id       uint64 `json:"id"`
	UserId   uint64 `json:"user_id"`
	LessonId uint64 `json:"lesson_id"`
	Mark     *byte  `json:"mark"`
	IsCame   *bool  `json:"is_came"`
}

type MarkUserInfo struct {
	Id     uint64 `json:"id"`
	UserId string `json:"user"`
	Mark   byte   `json:"mark"`
	IsCame bool   `json:"is_came"`
}

type MarkLessonInfo struct {
	Id       uint64 `json:"id"`
	LessonId string `json:"lesson"`
	Mark     byte   `json:"mark"`
	IsCame   bool   `json:"is_came"`
}

type MarkInfo struct {
	Mark   *byte `json:"mark" db:"mark"`
	IsCame bool  `json:"is_came" db:"is_came"`
}
