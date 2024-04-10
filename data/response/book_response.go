package response

type BooksResponse struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Status bool   `json:"status"`
}
