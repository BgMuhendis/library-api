package request

// CreateBookRequest
// @Description Request about creating Book
type CreateBookRequest struct {
	Name   string `validate:"required,min=2,max=30" json:"name"`
	Page   int    `validate:"required,min=20,max=500" json:"page"`
	Author string `validate:"required,min=10,max=30" json:"author"`
	Status bool   `json:"status"`
}

// UpdateBookRequest
// @Description Request about updating Book
type UpdateBookRequest struct {
	Id     int  `validate:"reguired" json:"id"`
	Status bool `validate:"required" json:"status"`
}

