package model

type CreateBackupTaskPayload struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

type ApiStatus struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
