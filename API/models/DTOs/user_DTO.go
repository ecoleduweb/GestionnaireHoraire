package DTOs

import "llio-api/models/enums"

type UserDTO struct {
	Id        int            `json:"id"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	Role      enums.UserRole `json:"role"`
}
