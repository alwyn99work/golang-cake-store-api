package request

type CakeUpdateRequest struct {
	Id          int     `validate:"required"`
	Title       string  `validate:"required,max=255,min=1" json:"title"`
	Description string  `validate:"min=1" json:"description"`
	Rating      float32 `validate:"numeric" json:"rating"`
	Image       string  `validate:"max=191,min=1" json:"image"`
}
