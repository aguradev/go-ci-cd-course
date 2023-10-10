package resources

import (
	"praktikum/models"
	"time"
)

type UserResponseDetail struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type BlogUserResponse struct {
	Judul  string `json:"judul"`
	Konten string `json:"konten"`
}

type UserResponse struct {
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Blogs []BlogUserResponse `json:"list_blog"`
}

func (GetUser *UserResponseDetail) GetDetailUserResponse(user models.Users) {
	GetUser.Id = user.Id
	GetUser.Name = user.Name
	GetUser.Email = user.Email
	GetUser.Password = user.Password
	GetUser.CreatedAt = user.CreatedAt
	GetUser.UpdateAt = user.UpdateAt
}

func MappingAllDataUsers(Users []models.Users) []UserResponse {

	var GetUsers []UserResponse

	for _, val := range Users {
		UserResponse := UserResponse{
			Name:  val.Name,
			Email: val.Email,
		}

		var GetBlog []BlogUserResponse

		for _, val := range val.Blogs {
			blogResponse := BlogUserResponse{
				Judul:  val.Judul,
				Konten: val.Konten,
			}

			GetBlog = append(GetBlog, blogResponse)
		}

		UserResponse.Blogs = GetBlog

		GetUsers = append(GetUsers, UserResponse)
	}

	return GetUsers

}
