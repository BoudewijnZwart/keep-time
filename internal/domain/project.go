package domain

type Project struct {
	Id int
	Name string
	Alias string
	Color Color
	Code string
	Description string
	Client *Client
}