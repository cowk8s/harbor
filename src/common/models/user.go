package models

type User struct {
	UserID int `json:"user_id"`
}

type Users []*User

// MapByUserID returns map which key is UserID of the user and value is the user itself
func (users Users) MapByUserID() map[int]*User {
	m := map[int]*User{}
	for _, user := range users {
		m[user.UserID] = user
	}
	return m
}
