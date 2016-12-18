package models

type User struct {
	login, password string
}

func (u User) GetLogin() string {
	return u.login
}

func (u *User) SetLogin(newLogin string) {
	u.login = newLogin
}
