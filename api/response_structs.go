package api

import (
	"time"

	"github.com/google/uuid"
)
// These structs will match the generated database structs of the same name.
// They are used to set the json metadata by casting the generated structs to these.

type User struct {
	ID			uuid.UUID	`json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	Email		string		`json:"email"`	
}

type Chirp struct {
	ID			uuid.UUID	`json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	Body		string		`json:"body"`	
	UserID   	uuid.UUID	`json:"user_id"`
}