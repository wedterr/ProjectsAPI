package model

// An author with Id, FirstName and LastName fields
type Author struct {
	// The id of the author
	ID string `json:"id,omitempty"`
	// The first name of the author
	FirstName string `json:"firstName,omitempty"`
	// The last name of the author
	LastName string `json:"lastName,omitempty"`
	// Author's projects
	ProjectsCollection
}
