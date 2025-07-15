package dto

type (
	RegisterUserDto struct {
		ID          string `json:"id" db:"id"`
		Email       string `json:"email" form:"email" binding:"required" db:"email"`
		Username    string `json:"username" form:"username" binding:"required" db:"username"`
		Password    string `json:"password" form:"password" binding:"required" db:"password"`
		FirstName   string `json:"first_name" form:"first_name" binding:"required" db:"first_name"`
		LastName    string `json:"last_name" form:"last_name" binding:"required" db:"last_name"`
		IsSuperuser bool   `json:"is_superuser" form:"is_superuser" db:"is_superuser"`
		IsActive    bool   `json:"is_active" form:"is_active" db:"is_active"`
	}
)
