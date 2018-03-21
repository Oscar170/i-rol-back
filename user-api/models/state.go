package models

type Status struct {
	Active bool `json:"active"`
}

type User struct {
	_id      string
	name     string
	mail     string
	password string
	friends  []string
	rooms    []string
}
