package models

type DAPs struct {
	Metadata APIMetadata `json:"metadata"`
	DAPs     []DAP       `json:"response"`
}

type DAP struct {
	AvailabilityMachine  string                  `json:"availabilityMachine"`
	DateModified         string                  `json:"dateModified"`
	DateCreated          string                  `json:"dateCreated"`
	EngineType           string                  `json:"engineType"`
	Name                 string                  `json:"name"`
	Owner                string                  `json:"owner"`
	Status               string                  `json:"status"`
	SharedWithUsers      []string                `json:"sharedWithUsers"`
	SharedWithUserGroups []string                `json:"sharedWithUserGroups"`
	TargetCloudLocations DAPTargetCloudLocations `json:"targetCloudLocations"`
	RetentionConfig      DAPRetentionConfig      `json:"retentionConfig"`
}

type DAPCreationUpdationPayload struct {
	Name                 string                  `json:"name"`
	UserGroups           []string                `json:"userGroups"`
	Users                []string                `json:"users"`
	TargetCloudLocations DAPTargetCloudLocations `json:"targetCloudLocations"`
	RetentionConfig      DAPRetentionConfig      `json:"retentionConfig"`
}

type DAPTargetCloudLocations struct {
	AWS   []string `json:"aws"`
	Azure []string `json:"azure"`
}

type DAPRetentionConfig struct {
	PitrRetention    DAPPitrRetention    `json:"pitrRetention"`
	DailyRetention   DAPDailyRetention   `json:"dailyRetention"`
	WeeklyRetention  DAPWeeklyRetention  `json:"weeklyRetention"`
	MonthlyRetention DAPMonthlyRetention `json:"monthlyRetention"`
	YearlyRetention  DAPYearlyRetention  `json:"yearlyRetention"`
}

type DAPPitrRetention struct {
	Days int `json:"days"`
}

type DAPDailyRetention struct {
	Days int `json:"days"`
}

type DAPWeeklyRetention struct {
	Weeks int      `json:"weeks"`
	Days  []string `json:"days"`
}

type DAPMonthlyRetention struct {
	Months                int                               `json:"months"`
	CommonSchedule        DAPMonthlyRetentionCommonSchedule `json:"commonSchedule"`
	MonthSpecificSchedule []DAPMonthSpecificSchedule        `json:"monthSpecificSchedule"`
}

type DAPMonthlyRetentionCommonSchedule struct {
	Dates          []int `json:"dates"`
	LastDayOfMonth bool  `json:"lastDayOfMonth"`
}

type DAPYearlyRetention struct {
	Years                 int                        `json:"years"`
	MonthSpecificSchedule []DAPMonthSpecificSchedule `json:"monthSpecificSchedule"`
}

type DAPMonthSpecificSchedule struct {
	Month string `json:"month"`
	Dates []int  `json:"dates"`
}
