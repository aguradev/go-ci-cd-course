package resources

import (
	"praktikum/models"
	"time"
)

type UserBlogResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdateAt  time.Time `json:"updated_at,omitempty"`
}

type BlogResponse struct {
	Judul  string            `json:"judul"`
	Konten string            `json:"konten"`
	Users  *UserBlogResponse `json:"created_user"`
}

func MappingBlogResponse(GetBlog []models.Blogs) []BlogResponse {

	var AllBlogResponse []BlogResponse

	for _, item := range GetBlog {

		BlogResponse := BlogResponse{
			Judul:  item.Judul,
			Konten: item.Konten,
			Users: &UserBlogResponse{
				Name:      item.Users.Name,
				Email:     item.Users.Email,
				Password:  item.Users.Password,
				CreatedAt: item.Users.CreatedAt,
				UpdateAt:  item.Users.UpdateAt,
			},
		}

		AllBlogResponse = append(AllBlogResponse, BlogResponse)

	}

	return AllBlogResponse

}

func (blog *BlogResponse) DetailBlogResponse(GetBlog models.Blogs) {

	blog.Judul = GetBlog.Judul
	blog.Konten = GetBlog.Konten
	blog.Users = &UserBlogResponse{
		Name:  GetBlog.Users.Name,
		Email: GetBlog.Users.Email,
	}

}
