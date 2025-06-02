package entity

import "manage-se/internal/provider/user"

type UserContext struct {
	ID     string    `json:"id" db:"id"`
	Name   string    `json:"name" db:"name"`
	Phone  string    `json:"phone" db:"phone"`
	RoleID string    `json:"role_id" db:"role_id"`
	Role   user.Role `json:"role" db:"role"`
}
