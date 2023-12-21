package model

type Student struct {
	User
	Classes []Class `json:"classes,omitempty"`
}
