package main

//Book type with Name, Author, and ISBN

type Book struct{
	Title   string `json:"title"`
	Author   string `json:"author"`
	ISBN   string `json:"isbn"`
	Description   string `json:"description,onitempty"`

}

var books = map[string]Book{
	"01234567" : Book{Title: "The sample Book", Author:"Abh",ISBN: "1234567"},
	"00000000" : Book{Title: "The sample Book", Author:"Abh",ISBN: "0000000"},
}
//AllBooks rteurns a slice of all books

func AllBooks() []Book{
values := make([]Book, len(books))
idx := 0
for _,book := range books{
	values[idx] = book
	idx++
}
return values
}

func GetBook(isbn string)(Book,bool){
	book,found := books[isbn]
	return book,found
}

func CreateBook(book Book)(string, bool){
	_,exists := books[book.ISBN]
	if exists {
		return "",false
	}
	books[book.ISBN] = book
	return book.ISBN, true

}

func UpdateBook(isbn string,book Book) bool{
	_,exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

func DeleteBook(isbn string) {
	delete(books,isbn)
}
