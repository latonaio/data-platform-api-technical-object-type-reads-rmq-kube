package requests

type TechnicalObjectType struct {
	TechnicalObjectType string `json:"TechnicalObjectType"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
