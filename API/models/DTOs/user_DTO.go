package DTOs

import "llio-api/models/enums"

type UserDTO struct {
	Id        int            `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Role      enums.UserRole `json:"role"`
}
