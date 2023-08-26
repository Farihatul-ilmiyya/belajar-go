package product

import "time"

type ResponseProduct struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Weight    int       `json:"weight"`
	Brand     string    `json:"brand"`
}
