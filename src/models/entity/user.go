package entity

import "time"

type User struct {
	id        int
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}
