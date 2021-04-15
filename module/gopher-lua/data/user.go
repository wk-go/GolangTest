package data

type User struct {
	Name  string
	token string
}

func (u *User) SetToken(t string) {
	u.token = t
}

func (u *User) Token() string {
	return u.token
}
