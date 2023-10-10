package books

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	Id        uint           `gorm:"primaryKey" json:"id,omitempty"`
	Judul     string         `gorm:"type:varchar(255)" json:"judul,omitempty"`
	Penulis   string         `gorm:"type:varchar(50)" json:"penulis,omitempty"`
	Penerbit  string         `gorm:"type:varchar(50)" json:"penerbit,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (book Books) GetBooksData(Request BooksRequest) Books {

	var getBook Books

	getBook.Judul = Request.Judul
	getBook.Penerbit = Request.Penerbit
	getBook.Penulis = Request.Penulis

	return getBook

}
