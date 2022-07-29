package model

type TessellDmmDataflixServiceView struct {
	AvailabilityMachineId *string                        `json:"availabilityMachineId,omitempty"`
	TessellServiceId      *string                        `json:"tessellServiceId,omitempty"`
	ServiceName           *string                        `json:"serviceName,omitempty"`
	EngineType            *string                        `json:"engineType,omitempty"`
	TimeZone              *string                        `json:"timeZone,omitempty"`    // Output timezone
	Owner                 *string                        `json:"owner,omitempty"`       // Owner of the Availability Machine
	PitrCatalog           *[]TessellDataflixPitrInfo     `json:"pitrCatalog,omitempty"` // PITR availability catalog
	SnapshotCatalog       *[]TessellDmmDataflixBackupDTO `json:"snapshotCatalog,omitempty"`
}

type TessellDataflixPitrInfo struct {
	Cloud   *string                             `json:"cloud,omitempty"`
	Regions *[]TessellDataflixPitrInfoForRegion `json:"regions,omitempty"`
}

type TessellDataflixPitrInfoForRegion struct {
	Region     *string                        `json:"region,omitempty"` // Region name
	TimeRanges *[]TessellDataflixFromTimeInfo `json:"timeRanges,omitempty"`
}

type TessellDataflixFromTimeInfo struct {
	FromTime   *string                      `json:"fromTime,omitempty"` // PITR recovery from-time
	ToTime     *string                      `json:"toTime,omitempty"`   // PITR recovery to-time
	SharedWith *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}

type TessellDmmDataflixBackupDTO struct {
	Id                *string                      `json:"id,omitempty"`           // DB Service snapshot Id
	Name              *string                      `json:"name,omitempty"`         // DB Service snapshot name
	Description       *string                      `json:"description,omitempty"`  // Description for the snapshot
	SnapshotTime      *string                      `json:"snapshotTime,omitempty"` // DB Service snapshot capture time
	Status            *string                      `json:"status,omitempty"`
	Size              *int                         `json:"size,omitempty"`   // Database Backup size in bytes
	Manual            *bool                        `json:"manual,omitempty"` // Specifies whether the backup is captured manually
	CloudAvailability *[]CloudRegionInfo1          `json:"cloudAvailability,omitempty"`
	Databases         *[]BackupDatabaseInfo        `json:"databases,omitempty"` // The databases that are captured as part of the snapshot
	SharedWith        *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}
