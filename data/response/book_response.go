package response

type BooksResponse struct {
	Id int `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Status bool   `json:"status"`
}
