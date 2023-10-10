package books

type BooksRequest struct {
	Judul    string `form:"judul"`
	Penulis  string `form:"penulis"`
	Penerbit string `form:"penerbit"`
}
