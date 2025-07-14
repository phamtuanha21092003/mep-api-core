package dto

type (
	RegisterUserDto struct {
		ID           string `db:"id"`
		Email        string `json:"email" form:"email" binding:"required"`
		Username     string `json:"username" form:"username" binding:"required"`
		Password     string `json:"password" form:"password" binding:"required"`
		FirstName    string `json:"first_name" form:"first_name" binding:"required"`
		LastName     string `json:"last_name" form:"last_name" binding:"required"`
		IsSupperuser bool   `json:"is_supperuser" form:"is_supperuser"`
		IsActive     bool   `json:"is_active" form:"is_active"`
	}
)
