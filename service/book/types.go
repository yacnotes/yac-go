package book

import "yac-go/model/book"

type Response struct {
	Id   int        `json:"id"`
	Book *book.Book `json:"book"`
}
