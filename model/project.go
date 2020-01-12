package model

type ProjectsCollection []Project

type Project struct {
	ID string `json:"id,omitempty"`

	AuthorFirstName string `json:"authorFirstName,omitempty"`

	AuthorLastName string `json:"authorLastName,omitempty"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`
}
