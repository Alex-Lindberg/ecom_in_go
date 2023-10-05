package models

import "time"

// CREATE TABLE customers (
//     id SERIAL PRIMARY KEY,
//     "name" VARCHAR(255) NOT NULL,
//     email VARCHAR(255) NOT NULL,
//     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
//     created_at TIMESTAMP NOT NULL DEFAULT NOW()
// );

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
