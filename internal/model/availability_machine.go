package model

type TessellDmmServiceConsumerDTO struct {
	Id                *string                            `json:"id,omitempty"`
	TessellServiceId  *string                            `json:"tessellServiceId,omitempty"`
	ServiceName       *string                            `json:"serviceName,omitempty"`
	Tenant            *string                            `json:"tenant,omitempty"`       // Dmm&#39;s tenancy details
	Subscription      *string                            `json:"subscription,omitempty"` // Dmm&#39;s subscription name
	EngineType        *string                            `json:"engineType,omitempty"`
	Status            *string                            `json:"status,omitempty"`
	UserId            *string                            `json:"userId,omitempty"`           // Data Management Machine&#39;s user
	Owner             *string                            `json:"owner,omitempty"`            // Availability Machine&#39;s owner
	LoggedInUserRole  *string                            `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the Availability Machine
	SharedWith        *EntityAclSharingInfo              `json:"sharedWith,omitempty"`
	CloudAvailability *[]CloudRegionInfo1                `json:"cloudAvailability,omitempty"`
	RpoSla            *TessellDmmAvailabilityServiceView `json:"rpoSla,omitempty"`
	Daps              *[]TessellDapServiceDTO            `json:"daps,omitempty"`
	Clones            *[]TessellCloneSummaryInfo         `json:"clones,omitempty"` // Clone databases that are created from this Availability Machine
	DateCreated       *string                            `json:"dateCreated,omitempty"`
	DateModified      *string                            `json:"dateModified,omitempty"`
}

type TessellDmmAvailabilityServiceView struct {
	AvailabilityMachineId *string             `json:"availabilityMachineId,omitempty"`
	AvailabilityMachine   *string             `json:"availabilityMachine,omitempty"` // Associated Availability Machine Name
	CloudAvailability     *[]CloudRegionInfo1 `json:"cloudAvailability,omitempty"`   // The availability location details: cloudAccount to region
	RpoSlaStatus          *string             `json:"rpoSlaStatus,omitempty"`        // The availability status
	Sla                   *string             `json:"sla,omitempty"`                 // Associated SLA
	Schedule              *ScheduleInfo       `json:"schedule,omitempty"`
}

type ScheduleInfo struct {
	BackupStartTime *ScheduleTimeFormat `json:"backupStartTime,omitempty"`
	DailySchedule   *DailySchedule      `json:"dailySchedule,omitempty"`
}

type ScheduleTimeFormat struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

type DailySchedule struct {
	BackupsPerDay    *int                  `json:"backupsPerDay,omitempty"`    // The number of backups to be captured per day, this is exclusive with &#39;backupStartTimes&#39;
	BackupStartTimes *[]ScheduleTimeFormat `json:"backupStartTimes,omitempty"` // List of times when backup(s) has to be captured at. If this is specified, the &#39;backupStartTime&#39; (at top level) value would be overridern by the first entry in this list
}

type TessellDapServiceDTO struct {
	Id                    *string                   `json:"id,omitempty"`
	Name                  *string                   `json:"name,omitempty"`
	AvailabilityMachineId *string                   `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                   `json:"tessellServiceId,omitempty"`
	ServiceName           *string                   `json:"serviceName,omitempty"`
	EngineType            *string                   `json:"engineType,omitempty"`
	ContentType           *string                   `json:"contentType,omitempty"`
	Status                *string                   `json:"status,omitempty"`
	ContentInfo           *DapContentInfo           `json:"contentInfo,omitempty"`
	CloudAvailability     *[]CloudRegionInfo1       `json:"cloudAvailability,omitempty"`
	DataAccessConfig      *RetentionAndScheduleInfo `json:"dataAccessConfig,omitempty"`
	Owner                 *string                   `json:"owner,omitempty"`
	LoggedInUserRole      *string                   `json:"loggedInUserRole,omitempty"` // The role of the logged in user for accessing the Availability Machine
	SharedWith            *EntityAclSharingInfo     `json:"sharedWith,omitempty"`
	DateCreated           *string                   `json:"dateCreated,omitempty"`
	DateModified          *string                   `json:"dateModified,omitempty"`
}

type DapContentInfo struct {
	SanitizedContent *SanitizationDapContent `json:"sanitizedContent,omitempty"`
}

type SanitizationDapContent struct {
	Automated *SanitizationDapContentAutomated `json:"automated,omitempty"`
}

type SanitizationDapContentAutomated struct {
	SanitizationScheduleId *string `json:"sanitizationScheduleId"` // Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.
}

type RetentionAndScheduleInfo struct {
	DailyBackups *int `json:"dailyBackups,omitempty"` // Number of daily backups to replicate
}

type TessellCloneSummaryInfo struct {
	Id                *string             `json:"id,omitempty"`
	Name              *string             `json:"name"`                   // Name of the clone database
	Subscription      *string             `json:"subscription,omitempty"` // Clone&#39;s subsription name
	Status            *string             `json:"status,omitempty"`       // Status of the clone database
	CloudAvailability *[]CloudRegionInfo1 `json:"cloudAvailability,omitempty"`
	CloneInfo         *map[string]string  `json:"cloneInfo,omitempty"`   // Miscellaneous information
	Owner             *string             `json:"owner,omitempty"`       // The user who created database clone
	DateCreated       *string             `json:"dateCreated,omitempty"` // Timestamp when the entity was created
}

type GetDmmsServiceView struct {
	Metadata *ApiMetadata1                   `json:"metadata,omitempty"`
	Response *[]TessellDmmServiceConsumerDTO `json:"response,omitempty"`
}
