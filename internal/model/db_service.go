package model

type CloneTessellServicePayload struct {
	SnapshotId               *string                                    `json:"snapshotId,omitempty"`  // Tessell service snapshot Id, using which the clone is to be created
	Pitr                     *string                                    `json:"pitr,omitempty"`        // PITR Timestamp, using which the clone is to be created
	Name                     *string                                    `json:"name"`                  // Tessell Service name
	Description              *string                                    `json:"description,omitempty"` // Tessell Service&#39;s description
	Subscription             *string                                    `json:"subscription"`          // Tessell Subscription in which the Tessell Service is to be created
	EngineType               *string                                    `json:"engineType"`
	Topology                 *string                                    `json:"topology"`
	NumOfInstances           *int                                       `json:"numOfInstances,omitempty"`           // Number of instance (nodes) to be created for the Tessell Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.
	SoftwareImage            *string                                    `json:"softwareImage"`                      // Software Image to be used to create the Tessell Service
	SoftwareImageVersion     *string                                    `json:"softwareImageVersion"`               // Software Image Version to be used to create the Tessell Service
	AutoMinorVersionUpdate   *bool                                      `json:"autoMinorVersionUpdate,omitempty"`   // Specify whether to automatically update minor version for Tessell Service
	EnableDeletionProtection *bool                                      `json:"enableDeletionProtection,omitempty"` // Specify whether to enable deletion protection for the Tessell Service
	Infrastructure           *TessellServiceInfrastructurePayload       `json:"infrastructure"`
	ServiceConnectivity      *TessellServiceConnectivityInfoPayload     `json:"serviceConnectivity"`
	Creds                    *TessellServiceCredsPayload                `json:"creds"`
	MaintenanceWindow        *TessellServiceMaintenanceWindow           `json:"maintenanceWindow,omitempty"`
	SnapshotConfiguration    *TessellServiceBackupConfigurationPayload  `json:"snapshotConfiguration,omitempty"`
	EngineConfiguration      *TessellServiceEngineConfigurationPayload1 `json:"engineConfiguration"`
	Databases                *[]CreateDatabasePayload1                  `json:"databases,omitempty"` // Specify the databases to be created in the Tessell Service
	IntegrationsConfig       *TessellServiceIntegrationsPayload         `json:"integrationsConfig,omitempty"`
	Tags                     *[]TessellTag                              `json:"tags,omitempty"` // The tags to be associated with the Tessell Service
}

type TessellServiceInfrastructurePayload struct {
	Cloud             *string `json:"cloud"`                       // The cloud-type in which the Tessell Service is to be provisioned (ex. aws, azure)
	Region            *string `json:"region"`                      // The region in which the Tessell Service is to be provisioned
	AvailabilityZone  *string `json:"availabilityZone,omitempty"`  // The availability-zone in which the Tessell Service is to be provisioned
	Vpc               *string `json:"vpc"`                         // The VPC to be used for provisioning the Tessell Service
	ComputeType       *string `json:"computeType"`                 // The compute-type to be used for provisioning the Tessell Service
	AdditionalStorage *int    `json:"additionalStorage,omitempty"` // The additional storage (in GBs) to be provisioned for the Tessell Service. This is in addition to what is specified in the compute type.
}

type TessellServiceConnectivityInfoPayload struct {
	DnsPrefix          *string   `json:"dnsPrefix,omitempty"`          // If not specified, Tessell will use a randomly generated prefix
	ServicePort        *int      `json:"servicePort"`                  // The connection port for the Tessell Service
	EnablePublicAccess *bool     `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the Tessell Service, default false
	AllowedIpAddresses *[]string `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the Tessell Service
}

type TessellServiceCredsPayload struct {
	MasterUser     *string `json:"masterUser"`     // Tessell Service&#39;s master username
	MasterPassword *string `json:"masterPassword"` // Tessell Service&#39;s master password
}

type TessellServiceMaintenanceWindow struct {
	Day      *string `json:"day"`
	Time     *string `json:"time"` // Time value in (hh:mm) format. ex. \&quot;02:00\&quot;
	Duration *int    `json:"duration"`
}

type TessellServiceBackupConfigurationPayload struct {
	AutoSnapshot   *bool                                                   `json:"autoSnapshot,omitempty"` // Specify whether to capture automated snapshots for the Tessell Service, default true.
	Sla            *string                                                 `json:"sla,omitempty"`          // The snapshot SLA for the Tessell Service. If not specified, a default SLA would be associated with the Tessell Service
	SnapshotWindow *TessellServiceBackupConfigurationPayloadSnapshotWindow `json:"snapshotWindow,omitempty"`
}

type TessellServiceBackupConfigurationPayloadSnapshotWindow struct {
	Time     *string `json:"time,omitempty"`     // Time value in (hh:mm) format. ex. \&quot;02:00\&quot;
	Duration *int    `json:"duration,omitempty"` // The allowed duration for capturing the Tessell Service backup
}

type TessellServiceEngineConfigurationPayload1 struct {
	PreScriptInfo     *ScriptInfo                     `json:"preScriptInfo,omitempty"`
	PostScriptInfo    *ScriptInfo                     `json:"postScriptInfo,omitempty"`
	OracleConfig      *OracleEngineConfigPayload      `json:"oracleConfig,omitempty"`
	PostgresqlConfig  *PostgresqlEngineConfigPayload  `json:"postgresqlConfig,omitempty"`
	MysqlConfig       *MySqlEngineConfigPayload       `json:"mysqlConfig,omitempty"`
	SqlServerConfig   *SqlServerEngineConfigPayload   `json:"sqlServerConfig,omitempty"`
	ApacheKafkaConfig *ApacheKafkaEngineConfigPayload `json:"apacheKafkaConfig,omitempty"`
}

type ScriptInfo struct {
	ScriptId      *string `json:"scriptId,omitempty"`      // The Tessell Script Id
	ScriptVersion *string `json:"scriptVersion,omitempty"` // The Tessell Script version
}

type OracleEngineConfigPayload struct {
	MultiTenant          *bool   `json:"multiTenant,omitempty"` // Specify whether the Tessell Service is multi-tenant.
	ParameterProfile     *string `json:"parameterProfile"`      // The parameter profile for the database
	OptionsProfile       *string `json:"optionsProfile"`        // The options profile for the database
	CharacterSet         *string `json:"characterSet"`          // The character-set for the database
	NationalCharacterSet *string `json:"nationalCharacterSet"`  // The national-character-set for the database
}

type PostgresqlEngineConfigPayload struct {
	ParameterProfile *string `json:"parameterProfile"` // The parameter profile for the database
}

type MySqlEngineConfigPayload struct {
	ParameterProfile *string `json:"parameterProfile"` // The parameter profile for the database
}

type SqlServerEngineConfigPayload struct {
	ParameterProfile *string `json:"parameterProfile"` // The parameter profile for the database
}

type ApacheKafkaEngineConfigPayload struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type CreateDatabasePayload1 struct {
	DatabaseName          *string                                     `json:"databaseName"`
	SourceDatabaseId      *string                                     `json:"sourceDatabaseId,omitempty"` // Required while creating a clone. It specifies the Id of the source database from which the clone is being created.
	DatabaseConfiguration *CreateDatabasePayloadDatabaseConfiguration `json:"databaseConfiguration,omitempty"`
}

type CreateDatabasePayloadDatabaseConfiguration struct {
	OracleConfig     *OracleDatabaseConfig     `json:"oracleConfig,omitempty"`
	PostgresqlConfig *PostgresqlDatabaseConfig `json:"postgresqlConfig,omitempty"`
	MysqlConfig      *MySqlDatabaseConfig      `json:"mysqlConfig,omitempty"`
	SqlServerConfig  *SqlServerDatabaseConfig  `json:"sqlServerConfig,omitempty"`
}

type OracleDatabaseConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
	OptionsProfile   *string `json:"optionsProfile,omitempty"`   // The options profile for the database
}

type PostgresqlDatabaseConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type MySqlDatabaseConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type SqlServerDatabaseConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type TessellServiceIntegrationsPayload struct {
	Integrations *[]string `json:"integrations,omitempty"`
}

type TessellTag struct {
	Name  *string `json:"name,omitempty"`  // Case sensitive, tag name
	Value *string `json:"value,omitempty"` // Case sensitive, tag value
}

type DeleteTessellServicePayload struct {
	DeletionConfig *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
}

type TessellServiceDeletionConfig struct {
	RetainAvailabilityMachine *bool `json:"retainAvailabilityMachine,omitempty"` // If &#39;retainAvailabilityMachine&#39; is true then set value of field takeFinalBackup and dapsToRetain. By default retainAvailabilityMachine is false, that means delete all details like Availability Machine, Backups, DAPs etc.
}

type TessellServiceClonedFromInfo struct {
	TessellServiceId      *string `json:"tessellServiceId,omitempty"`      // The Tessell Service Id using which this Tessell Service clone is created
	AvailabilityMachineId *string `json:"availabilityMachineId,omitempty"` // The Availability Machine Id using which this Tessell Service clone is created
	TessellService        *string `json:"tessellService,omitempty"`        // The Tessell Service name using which this Tessell Service clone is created
	AvailabilityMachine   *string `json:"availabilityMachine,omitempty"`   // The Availaility Machine name using which this Tessell Service clone is created
	SnapshotName          *string `json:"snapshotName,omitempty"`          // The snapshot using which this Tessell Service clone is created
	SnapshotId            *string `json:"snapshotId,omitempty"`            // The snapshot Id using which this Tessell Service clone is created
	PitrTime              *string `json:"pitrTime,omitempty"`              // If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC
}

type TessellServiceConnectivityInfo struct {
	DnsPrefix            *string                                         `json:"dnsPrefix,omitempty"`
	ServicePort          *int                                            `json:"servicePort,omitempty"`        // The connection port for the Tessell Service
	EnablePublicAccess   *bool                                           `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the Tessell Service, default false
	AllowedIpAddresses   *[]string                                       `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the Tessell Service
	ConnectStrings       *[]TessellServiceConnectString                  `json:"connectStrings,omitempty"`     // The list of connect strings for the Tessell Service
	UpdateInProgressInfo *TessellServiceConnectivityUpdateInProgressInfo `json:"updateInProgressInfo,omitempty"`
}

type TessellServiceConnectString struct {
	Type              *string `json:"type,omitempty"`
	ConnectDescriptor *string `json:"connectDescriptor,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	MasterUser        *string `json:"masterUser,omitempty"`
	ServicePort       *int    `json:"servicePort,omitempty"` // The connection port for the Tessell Service
}

type TessellServiceConnectivityUpdateInProgressInfo struct {
	DnsPrefix          *string   `json:"dnsPrefix,omitempty"`
	EnablePublicAccess *bool     `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the Tessell Service, default false
	AllowedIpAddresses *[]string `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the Tessell Service
}

type TessellServiceInfrastructureInfo struct {
	Cloud             *string             `json:"cloud,omitempty"`            // The cloud-type in which the Tessell Service is provisioned (ex. aws, azure)
	Region            *string             `json:"region,omitempty"`           // The region in which the Tessell Service provisioned
	AvailabilityZone  *string             `json:"availabilityZone,omitempty"` // The availability-zone in which the Tessell Service is provisioned
	CloudAvailability *[]CloudRegionInfo1 `json:"cloudAvailability,omitempty"`
	Vpc               *string             `json:"vpc,omitempty"`               // The VPC to be used for provisioning the Tessell Service
	ComputeType       *string             `json:"computeType,omitempty"`       // The compute-type to be used for provisioning the Tessell Service
	AdditionalStorage *int                `json:"additionalStorage,omitempty"` // The additional storage (in GBs) to be provisioned for the Tessell Service. This is in addition to what is specified in the compute type.
}

type TessellServiceEngineInfo struct {
	OracleConfig      *TessellServiceOracleEngineConfig      `json:"oracleConfig,omitempty"`
	PostgresqlConfig  *TessellServicePostgresqlEngineConfig  `json:"postgresqlConfig,omitempty"`
	MysqlConfig       *TessellServiceMySqlEngineConfig       `json:"mysqlConfig,omitempty"`
	SqlServerConfig   *TessellServiceSqlServerEngineConfig   `json:"sqlServerConfig,omitempty"`
	ApacheKafkaConfig *TessellServiceApacheKafkaEngineConfig `json:"apacheKafkaConfig,omitempty"`
	PreScriptInfo     *ScriptInfo                            `json:"preScriptInfo,omitempty"`
	PostScriptInfo    *ScriptInfo                            `json:"postScriptInfo,omitempty"`
}

type TessellServiceOracleEngineConfig struct {
	MultiTenant          *bool   `json:"multiTenant,omitempty"`          // Specify whether the Tessell Service is multi-tenant.
	ParameterProfile     *string `json:"parameterProfile,omitempty"`     // The parameter profile for the database
	OptionsProfile       *string `json:"optionsProfile,omitempty"`       // The options profile for the database
	CharacterSet         *string `json:"characterSet,omitempty"`         // The character-set for the database
	NationalCharacterSet *string `json:"nationalCharacterSet,omitempty"` // The national-character-set for the database
}

type TessellServicePostgresqlEngineConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type TessellServiceMySqlEngineConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type TessellServiceSqlServerEngineConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type TessellServiceApacheKafkaEngineConfig struct {
	ParameterProfile *string `json:"parameterProfile,omitempty"` // The parameter profile for the database
}

type TessellServiceIntegrationsInfo struct {
	Integrations *[]string `json:"integrations,omitempty"`
}

type TessellServiceInstanceDTO struct {
	Id                *string                              `json:"id,omitempty"`   // Tessell generated UUID for the Tessell Service Instance
	Name              *string                              `json:"name,omitempty"` // Name of the Tessell Service Instance
	Role              *string                              `json:"role,omitempty"` // Tessell Service Topology
	Status            *string                              `json:"status,omitempty"`
	TessellServiceId  *string                              `json:"tessellServiceId,omitempty"` // Tessell Service Instance&#39;s associated Tessell Service id
	ComputeType       *string                              `json:"computeType,omitempty"`      // The compute used for creation of the Tessell Service Instance
	Cloud             *string                              `json:"cloud,omitempty"`            // Tessell Service Instance&#39;s cloud type
	Region            *string                              `json:"region,omitempty"`           // Tessell Service Instance&#39;s cloud region
	AvailabilityZone  *string                              `json:"availabilityZone,omitempty"` // Tessell Service Instance&#39;s cloud availability zone
	DateCreated       *string                              `json:"dateCreated,omitempty"`      // Timestamp when the entity was created
	ConnectString     *TessellServiceInstanceConnectString `json:"connectString,omitempty"`
	UpdatesInProgress *[]TessellResourceUpdateInfo         `json:"updatesInProgress,omitempty"` // The updates that are in progress for this resource
}

type TessellServiceInstanceConnectString struct {
	ConnectDescriptor *string `json:"connectDescriptor,omitempty"`
	MasterUser        *string `json:"masterUser,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	ServicePort       *string `json:"servicePort,omitempty"`
}

type TessellResourceUpdateInfo struct {
	UpdateType  *string                            `json:"updateType,omitempty"`  // Type of the update
	ReferenceId *string                            `json:"referenceId,omitempty"` // The reference-id of the update request
	SubmittedAt *string                            `json:"submittedAt,omitempty"` // Timestamp when the resource update was requested
	UpdateInfo  *map[string]map[string]interface{} `json:"updateInfo,omitempty"`  // The specific details for a Tessell resource that are being updated
}

type TessellDatabaseDTO struct {
	Id                    *string                        `json:"id,omitempty"`
	DatabaseName          *string                        `json:"databaseName,omitempty"`     // Database name
	Description           *string                        `json:"description,omitempty"`      // Database description
	TessellServiceId      *string                        `json:"tessellServiceId,omitempty"` // Associated Tessell Service Id
	EngineType            *string                        `json:"engineType,omitempty"`       // Database engine type
	Status                *string                        `json:"status,omitempty"`           // Database status
	DateCreated           *string                        `json:"dateCreated,omitempty"`      // Timestamp when the entity was created
	ClonedFromInfo        *TessellDatabaseClonedFromInfo `json:"clonedFromInfo,omitempty"`
	DatabaseConfiguration *DatabaseConfiguration         `json:"databaseConfiguration,omitempty"`
}

type TessellDatabaseClonedFromInfo struct {
	DatabaseId *string `json:"databaseId,omitempty"` // The original database Id using which this database clone is created
}

type DatabaseConfiguration struct {
	OracleConfig     *OracleDatabaseConfig     `json:"oracleConfig,omitempty"`
	PostgresqlConfig *PostgresqlDatabaseConfig `json:"postgresqlConfig,omitempty"`
	MySqlConfig      *MySqlDatabaseConfig      `json:"mySqlConfig,omitempty"`
	SqlServerConfig  *SqlServerDatabaseConfig  `json:"sqlServerConfig,omitempty"`
}

type TessellServicesResponse struct {
	Metadata *ApiMetadata1        `json:"metadata,omitempty"`
	Response *[]TessellServiceDTO `json:"response,omitempty"`
}

type TessellServiceDTO struct {
	Id                       *string                           `json:"id,omitempty"`                    // Tessell generated UUID for the Tessell Service
	AvailabilityMachineId    *string                           `json:"availabilityMachineId,omitempty"` // Associated Availability Machine Id
	Name                     *string                           `json:"name,omitempty"`                  // Name of the Tessell Service
	Description              *string                           `json:"description,omitempty"`           // Tessell Service description
	EngineType               *string                           `json:"engineType,omitempty"`
	Topology                 *string                           `json:"topology,omitempty"`
	NumOfInstances           *int                              `json:"numOfInstances,omitempty"`
	Status                   *string                           `json:"status,omitempty"`
	LicenseType              *string                           `json:"licenseType,omitempty"`
	AutoMinorVersionUpdate   *bool                             `json:"autoMinorVersionUpdate,omitempty"`   // Specifies whether to automatically update minor version for Tessell Service
	EnableDeletionProtection *bool                             `json:"enableDeletionProtection,omitempty"` // Specifies whether to enable deletion protection for the Tessell Service
	SoftwareImage            *string                           `json:"softwareImage,omitempty"`            // The Software Image that is used to create the Tessell Service
	SoftwareImageVersion     *string                           `json:"softwareImageVersion,omitempty"`     // The Software Image version that is used to create the Tessell Service
	TenantId                 *string                           `json:"tenantId,omitempty"`                 // The tenant-id for the Tessell Service
	Subscription             *string                           `json:"subscription,omitempty"`             // The subscription-name in which this Tessell Service is created
	Owner                    *string                           `json:"owner,omitempty"`                    // Tessell Service owner email address
	LoggedInUserRole         *string                           `json:"loggedInUserRole,omitempty"`         // Access role for the currently logged in user
	DateCreated              *string                           `json:"dateCreated,omitempty"`              // Timestamp when the Tessell Service was created at
	StartedAt                *string                           `json:"startedAt,omitempty"`                // Timestamp when the Tessell Service was last started at
	StoppedAt                *string                           `json:"stoppedAt,omitempty"`                // Timestamp when the Tessell Service was last stopped at
	ClonedFromInfo           *TessellServiceClonedFromInfo     `json:"clonedFromInfo,omitempty"`
	ServiceConnectivity      *TessellServiceConnectivityInfo   `json:"serviceConnectivity,omitempty"`
	Infrastructure           *TessellServiceInfrastructureInfo `json:"infrastructure,omitempty"`
	MaintenanceWindow        *TessellServiceMaintenanceWindow  `json:"maintenanceWindow,omitempty"`
	EngineConfiguration      *TessellServiceEngineInfo         `json:"engineConfiguration,omitempty"`
	IntegrationsConfig       *TessellServiceIntegrationsInfo   `json:"integrationsConfig,omitempty"`
	DeletionConfig           *TessellServiceDeletionConfig     `json:"deletionConfig,omitempty"`
	Tags                     *[]TessellTag                     `json:"tags,omitempty"`      // The tags associated with the Tessell Service
	Instances                *[]TessellServiceInstanceDTO      `json:"instances,omitempty"` // Instances associated with this Tessell Service
	Databases                *[]TessellDatabaseDTO             `json:"databases,omitempty"` // Databases that are part of this Tessell Service
	SharedWith               *EntityAclSharingInfo             `json:"sharedWith,omitempty"`
}

type ProvisionTessellServicePayload struct {
	Name                     *string                                   `json:"name"`                  // Tessell Service name
	Description              *string                                   `json:"description,omitempty"` // Tessell Service&#39;s description
	Subscription             *string                                   `json:"subscription"`          // Tessell Subscription in which the Tessell Service is to be created
	EngineType               *string                                   `json:"engineType"`
	Topology                 *string                                   `json:"topology"`
	NumOfInstances           *int                                      `json:"numOfInstances,omitempty"`           // Number of instance (nodes) to be created for the Tessell Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.
	SoftwareImage            *string                                   `json:"softwareImage"`                      // Software Image to be used to create the Tessell Service
	SoftwareImageVersion     *string                                   `json:"softwareImageVersion"`               // Software Image Version to be used to create the Tessell Service
	AutoMinorVersionUpdate   *bool                                     `json:"autoMinorVersionUpdate,omitempty"`   // Specify whether to automatically update minor version for Tessell Service
	EnableDeletionProtection *bool                                     `json:"enableDeletionProtection,omitempty"` // Specify whether to enable deletion protection for the Tessell Service
	Infrastructure           *TessellServiceInfrastructurePayload      `json:"infrastructure"`
	ServiceConnectivity      *TessellServiceConnectivityInfoPayload    `json:"serviceConnectivity"`
	Creds                    *TessellServiceCredsPayload               `json:"creds"`
	MaintenanceWindow        *TessellServiceMaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	SnapshotConfiguration    *TessellServiceBackupConfigurationPayload `json:"snapshotConfiguration,omitempty"`
	EngineConfiguration      *TessellServiceEngineConfigurationPayload `json:"engineConfiguration"`
	Databases                *[]CreateDatabasePayload                  `json:"databases,omitempty"` // Specify the databases to be created in the Tessell Service
	IntegrationsConfig       *TessellServiceIntegrationsPayload        `json:"integrationsConfig,omitempty"`
	Tags                     *[]TessellTag                             `json:"tags,omitempty"` // The tags to be associated with the Tessell Service
}

type TessellServiceEngineConfigurationPayload struct {
	PreScriptInfo     *ScriptInfo                     `json:"preScriptInfo,omitempty"`
	PostScriptInfo    *ScriptInfo                     `json:"postScriptInfo,omitempty"`
	OracleConfig      *OracleEngineConfigPayload      `json:"oracleConfig,omitempty"`
	PostgresqlConfig  *PostgresqlEngineConfigPayload  `json:"postgresqlConfig,omitempty"`
	MysqlConfig       *MySqlEngineConfigPayload       `json:"mysqlConfig,omitempty"`
	SqlServerConfig   *SqlServerEngineConfigPayload   `json:"sqlServerConfig,omitempty"`
	ApacheKafkaConfig *ApacheKafkaEngineConfigPayload `json:"apacheKafkaConfig,omitempty"`
}

type CreateDatabasePayload struct {
	DatabaseName          *string                                     `json:"databaseName"`
	SourceDatabaseId      *string                                     `json:"sourceDatabaseId,omitempty"` // Required while creating a clone. It specifies the Id of the source database from which the clone is being created.
	DatabaseConfiguration *CreateDatabasePayloadDatabaseConfiguration `json:"databaseConfiguration,omitempty"`
}
