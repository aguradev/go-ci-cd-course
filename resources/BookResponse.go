package resources

import books "praktikum/models/Books"

type BooksResponse struct {
	Id       uint   `json:"id,omitempty"`
	Judul    string `json:"judul,omitempty"`
	Penulis  string `json:"penulis,omitempty"`
	Penerbit string `json:"penerbit,omitempty"`
}

func MappingBooksResponse(book []books.Books) []BooksResponse {

	var SetBooks []BooksResponse

	for _, val := range book {

		BooksResponse := BooksResponse{
			Id:       val.Id,
			Judul:    val.Judul,
			Penulis:  val.Penulis,
			Penerbit: val.Penerbit,
		}

		SetBooks = append(SetBooks, BooksResponse)

	}

	return SetBooks

}
