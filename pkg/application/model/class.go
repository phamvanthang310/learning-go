package model

import "time"

type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
	Students  []Student `json:"students,omitempty"`
	ManagedBy *Teacher  `json:"managedBy,omitempty"`
}

type AssignStudent struct {
	StudentIds []string `json:"studentIds" validate:"required"`
}
