package pkg01

type Struct01 struct {
	Name     string
	Age      string
	_Message string
	_message string
	字段1      string
	F字段1     string
}

func NewStruct01(v1, v2, v3, v4, v5, v6 string) *Struct01 {
	return &Struct01{
		Name:     v1,
		Age:      v2,
		_Message: v3,
		_message: v4,
		字段1:      v5,
		F字段1:     v6,
	}
}
