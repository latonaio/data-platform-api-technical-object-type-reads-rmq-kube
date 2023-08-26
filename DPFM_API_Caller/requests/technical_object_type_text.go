package requests

type TechnicalObjectTypeText struct {
	TechnicalObjectType     string `json:"TechnicalObjectType"`
	Language          		string `json:"Language"`
	TechnicalObjectTypeName	string `json:"TechnicalObjectTypeName"`
	CreationDate			string `json:"CreationDate"`
	LastChangeDate			string `json:"LastChangeDate"`
	IsMarkedForDeletion		*bool  `json:"IsMarkedForDeletion"`
}
