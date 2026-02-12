package models

type MigrationStatus struct {
	ID         int64  `json:"id"`
	MigratorName string `json:"migrator_name"`
	Created    string `json:"created"`
}

type MigrationAuth struct {
	URL string `json:"url"`
}
