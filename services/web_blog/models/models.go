package models

import "time"

type Authors struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
}
