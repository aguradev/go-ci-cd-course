package models

type Blogs struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `json:"user_id"`
	Judul  string `gorm:"type:varchar(100)" json:"judul"`
	Konten string `gorem:"type:longtext" json:"konten"`
	Users  Users  `gorm:"foreignKey:UserID" json:"created_user"`
}

type BlogRequest struct {
	UserID uint   `json:"user_id"`
	Judul  string `json:"judul"`
	Konten string `json:"konten"`
}

func (blog Blogs) RequestBlog(request BlogRequest) Blogs {

	var getBlog Blogs

	getBlog.Judul = request.Judul
	getBlog.Konten = request.Konten
	getBlog.UserID = request.UserID

	return getBlog

}
