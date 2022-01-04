package models

type DatabaseBackups struct {
	DatabaseBackups []DatabaseBackup
}

type DatabaseBackup struct {
	BackupTime    string `json:"backupTime"`
	BackupType    string `json:"backupType"`
	CloudLocation string `json:"cloudLocation"`
	Manual        bool   `json:"manual"`
	Name          string `json:"name"`
	RetentionType string `json:"retentionType"`
	Status        string `json:"status"`
	Size          int    `json:"size"`
}

type DatabaseBackupCreationPayload struct {
	BackupName  string `json:"backupName"`
	Description string `json:"description"`
}

type DatabaseBackupCreationResponse struct {
	TaskId   string `json:"taskId"`
	TaskType string `json:"taskType"`
	Details  struct {
		DatabaseBackup        string `json:"databaseBackup"`
		Database              string `json:"database"`
		DataManagementMachine string `json:"dataManagementMachine"`
	} `json:"details"`
}

type DatabaseBackupDeletionResponse struct {
	TaskId   string `json:"taskId"`
	TaskType string `json:"taskType"`
	Details  struct {
		DatabaseBackup        string `json:"databaseBackup"`
		DataManagementMachine string `json:"dataManagementMachine"`
	} `json:"details"`
}
