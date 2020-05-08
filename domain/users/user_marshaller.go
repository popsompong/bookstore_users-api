package users

import "encoding/json"

type PublicUser struct {
	Id         int64  `json:"id"`
	DataCreate string `json:"data_create"`
	Status     string `json:"status"`
}

type PrivateUser struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DataCreate string `json:"data_create"`
	Status     string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for key, user := range users {
		result[key] = user.Marshall(isPublic)
	}
	return result
}
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:         user.Id,
			DataCreate: user.DataCreate,
			Status:     user.Status,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
