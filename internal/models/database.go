package models

type Databases struct {
	Metadata  APIMetadata `json:"metadata"`
	Databases []Database  `json:"response"`
}

type Database struct {
	Name                  string                 `json:"name"`
	Subscription          string                 `json:"subscription"`
	Description           string                 `json:"description"`
	AvailabilityMachine   string                 `json:"availabilityMachine"`
	EngineType            string                 `json:"engineType"`
	Status                string                 `json:"status"`
	Clone                 bool                   `json:"clone"`
	ConfigType            string                 `json:"configType"`
	Topology              string                 `json:"topology"`
	DatabaseConfiguration DatabaseConfiguration  `json:"databaseConfiguration"`
	Infrastructure        DatabaseInfrastructure `json:"infrastructure"`
	Info                  interface{}            `json:"info"`
	DateCreated           string                 `json:"dateCreated"`
	DateModified          string                 `json:"dateModified"`
	User                  string                 `json:"user"`
}

type DatabaseConfiguration struct {
	AutoMinorVersionUpdate     bool                                   `json:"autoMinorVersionUpdate"`
	CharacterSet               string                                 `json:"characterSet"`
	CreateReplica              bool                                   `json:"createReplica"`
	DatabaseComputeId          string                                 `json:"databaseComputeId"`
	DatabaseNetworkProfileId   string                                 `json:"databaseNetworkProfileId"`
	DatabaseOptionsProfileId   string                                 `json:"databaseOptionsProfileId"`
	DatabaseParameterProfileId string                                 `json:"databaseParameterProfileId"`
	DatabaseSoftwareImageId    string                                 `json:"databaseSoftwareImageId"`
	EnableDeletionProtection   bool                                   `json:"enableDeletionProtection"`
	LicenseType                string                                 `json:"licenseType"`
	MaintenanceWindow          DatabaseConfigurationMaintenanceWindow `json:"maintenanceWindow"`
	Multitenant                bool                                   `json:"multitenant"`
	NationalCharacterSet       string                                 `json:"nationalCharacterSet"`
	OptionsProfile             string                                 `json:"optionsProfile"`
	ParameterProfile           string                                 `json:"parameterProfile"`
	SoftwareImageVersion       string                                 `json:"softwareImageVersion"`
	CloneMetadataInfo          struct {
		CreatedFromBackupId   string `json:"createdFromBackupId"`
		CreatedFromBackupName string `json:"createdFromBackupName"`
		CreatedFromPitr       string `json:"createdFromPitr"`
		RpoBased              bool   `json:"rpoBased"`
		ManualBackup          bool   `json:"manualBackup"`
		CreatedFromDapId      string `json:"createdFromDapId"`
		CreatedFromDapName    string `json:"createdFromDapName"`
	} `json:"CloneMetadataInfo"`
}

type DatabaseConfigurationMaintenanceWindow struct {
	Day      string `json:"day"`
	Duration int32  `json:"duration"`
	Time     string `json:"time"`
}

type DatabaseInfrastructure struct {
	CloudAvailability struct {
		AWS   map[string][]string `json:"aws"`
		Azure map[string][]string `json:"azure"`
	} `json:"cloudAvailability"`
	ConnectionInfo struct {
		Data struct {
			ConnectDesc        string   `json:"connectDesc"`
			Endpoint           string   `json:"endpoint"`
			DatabasePort       string   `json:"databasePort"`
			AllowedIpAddresses []string `json:"allowedIpAddresses"`
			PublicAccess       bool     `json:"publicAccess"`
			DeploymentId       string   `json:"deploymentId"`
			LicenseToken       string   `json:"licenseToken"`
		} `json:"data"`
	} `json:"connectionInfo"`
	Compute struct {
		ReadIops  int     `json:"readIops"`
		WriteIops int     `json:"writeIops"`
		MemoryGB  float64 `json:"memoryGB"`
		Name      string  `json:"name"`
		NoOfDisks int     `json:"noOfDisks"`
		StorageGB float64 `json:"storageGB"`
		Vcpus     int     `json:"vcpus"`
	} `json:"compute"`
}

type DatabaseCreationPayload struct {
	BackupConfiguration   DatabaseCreationBackupConfiguration   `json:"backupConfiguration"`
	ConfigType            string                                `json:"configType"`
	DatabaseConfiguration DatabaseCreationDatabaseConfiguration `json:"databaseConfiguration"`
	DatabaseName          string                                `json:"databaseName"`
	Description           string                                `json:"description"`
	DmmName               string                                `json:"dmmName"`
	EngineType            string                                `json:"engineType"`
	Infrastructure        DatabaseCreationInfrastructure        `json:"infrastructure"`
	MasterPassword        string                                `json:"masterPassword"`
	MasterUser            string                                `json:"masterUser"`
	PostScript            string                                `json:"postScript"`
	PreScript             string                                `json:"preScript"`
	ServiceName           string                                `json:"serviceName"`
	Subscription          string                                `json:"subscription"`
	Tags                  []interface{}                         `json:"tags"`
}

type DatabaseCreationBackupConfiguration struct {
	AutoBackup   bool                         `json:"autoBackup"`
	BackupSLA    string                       `json:"backupSla"`
	BackupWindow DatabaseCreationBackupWindow `json:"backupWindow"`
}

type DatabaseCreationBackupWindow struct {
	Duration int    `json:"duration"`
	Time     string `json:"time"`
}

type DatabaseCreationInfrastructure struct {
	AllowedIpAddresses []string `json:"allowedIpAddresses"`
	AvailabilityZone   string   `json:"availabilityZone"`
	Cloud              string   `json:"cloud"`
	ComputeType        string   `json:"computeType"`
	DatabasePort       int      `json:"databasePort"`
	EnablePublicAccess bool     `json:"enablePublicAccess"`
	Region             string   `json:"region"`
	VPCName            string   `json:"vpcName"`
}

type DatabaseCreationDatabaseConfiguration struct {
	AutoMinorVersionUpdate   bool                              `json:"autoMinorVersionUpdate"`
	CharacterSet             string                            `json:"characterSet"`
	CreateReplica            bool                              `json:"createReplica"`
	EnableDeletionProtection bool                              `json:"enableDeletionProtection"`
	LicenseType              string                            `json:"licenseType"`
	MaintenanceWindow        DatabaseCreationMaintenanceWindow `json:"maintenanceWindow"`
	Multitenant              bool                              `json:"multitenant"`
	NationalCharacterSet     string                            `json:"nationalCharacterSet"`
	OptionsProfile           string                            `json:"optionsProfile"`
	ParameterProfile         string                            `json:"parameterProfile"`
	SoftwareImageVersion     string                            `json:"softwareImageVersion"`
}

type DatabaseCreationMaintenanceWindow struct {
	Day      string `json:"day"`
	Duration int    `json:"duration"`
	Time     string `json:"time"`
}

type DatabaseCreationResponse struct {
	TaskId   string `json:"taskId"`
	TaskType string `json:"taskType"`
	Details  struct {
		DatabaseName string `json:"databaseName"`
		Action       string `json:"action"`
	} `json:"details"`
}

type DatabaseCloningPayload struct {
	Backup                string                               `json:"backup"`
	Pitr                  string                               `json:"pitr"`
	CloneDatabaseName     string                               `json:"cloneDatabaseName"`
	Description           string                               `json:"description"`
	ConfigType            string                               `json:"configType"`
	DatabaseConfiguration DatabaseCloningDatabaseConfiguration `json:"databaseConfiguration"`
	DmmName               string                               `json:"dmmName"`
	EngineType            string                               `json:"engineType"`
	Infrastructure        DatabaseCreationInfrastructure       `json:"infrastructure"`
	MasterPassword        string                               `json:"masterPassword"`
	MasterUser            string                               `json:"masterUser"`
	PostScript            string                               `json:"postScript"`
	PreScript             string                               `json:"preScript"`
	ServiceName           string                               `json:"serviceName"`
	Subscription          string                               `json:"subscription"`
	Tags                  []interface{}                        `json:"tags"`
}

type DatabaseCloningDatabaseConfiguration struct {
	AutoMinorVersionUpdate   bool   `json:"autoMinorVersionUpdate"`
	EnableDeletionProtection bool   `json:"enableDeletionProtection"`
	LicenseType              string `json:"licenseType"`
	OptionsProfile           string `json:"optionsProfile"`
	ParameterProfile         string `json:"parameterProfile"`
	SoftwareImageVersion     string `json:"softwareImageVersion"`
}

type DatabaseUpdationPayload struct {
	AllowedIpAddresses []string      `json:"allowedIpAddresses"`
	EnablePublicAccess bool          `json:"enablePublicAccess"`
	Description        string        `json:"description"`
	Tags               []interface{} `json:"tags"`
}

type DatabaseDeletionPayload struct {
	RetainAvailabilityMachine bool     `json:"retainAvailabilityMachine"`
	TakeFinalBackup           bool     `json:"takeFinalBackup"`
	DapsToRetain              []string `json:"dapsToRetain"`
}
