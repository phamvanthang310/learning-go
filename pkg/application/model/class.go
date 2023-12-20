package model

import "time"

type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
	Students  []Student `json:"students"`
	ManagedBy *Teacher  `json:"managedBy"`
}
