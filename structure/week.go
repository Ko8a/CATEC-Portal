package structure

type Week struct {
	Id        uint64 `json:"-"`
	GroupId   uint64 `json:"group_id"`
	StartDate uint64 `json:"start_date"`
	EndDate   uint64 `json:"end_date"`
	TypeId    uint64 `json:"type_id"`
}
