package model

import "time"

type Class struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	StartDate time.Time  `json:"startDate"`
	EndDate   time.Time  `json:"endDate"`
	CreatedAt time.Time  `json:"createdAt"`
	Subjects  []*Subject `json:"subjects"`
	ManagedBy *Teacher   `json:"managedBy"`
}

type NewClass struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	ManagedBy int64     `json:"managedBy"`
	Subjects  []int64   `json:"subjects"`
}
