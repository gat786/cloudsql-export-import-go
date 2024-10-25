package models

import (
	"time"
)

type DefaultUserLabels struct {
	Environment    string `json:"environment"`
	TerraformRepo  string `json:"terraform_repo"`
	Team           string `json:"team"`
	PrimaryContact string `json:"primary_contact"`
	Stakeholder    string `json:"stakeholder"`
}

type CloudSQLServerCaCert struct {
	Kind             string    `json:"kind"`
	CertSerialNumber string    `json:"certSerialNumber"`
	Cert             string    `json:"cert"`
	CommonName       string    `json:"commonName"`
	Sha1Fingerprint  string    `json:"sha1Fingerprint"`
	Instance         string    `json:"instance"`
	CreateTime       time.Time `json:"createTime"`
	ExpirationTime   time.Time `json:"expirationTime"`
}

type CloudSQLIPConfiguration struct {
	PrivateNetwork                          string        `json:"privateNetwork"`
	AuthorizedNetworks                      []interface{} `json:"authorizedNetworks"`
	SslMode                                 string        `json:"sslMode"`
	Ipv4Enabled                             bool          `json:"ipv4Enabled"`
	RequireSsl                              bool          `json:"requireSsl"`
	EnablePrivatePathForGoogleCloudServices bool          `json:"enablePrivatePathForGoogleCloudServices"`
	ServerCaMode                            string        `json:"serverCaMode"`
}

type CloudSQLDatabaseSettings struct {
	AuthorizedGaeApplications []interface{}           `json:"authorizedGaeApplications"`
	Tier                      string                  `json:"tier"`
	Kind                      string                  `json:"kind"`
	UserLabels                map[string]string       `json:"userLabels"`
	AvailabilityType          string                  `json:"availabilityType"`
	PricingPlan               string                  `json:"pricingPlan"`
	ReplicationType           string                  `json:"replicationType"`
	ActivationPolicy          string                  `json:"activationPolicy"`
	IPConfiguration           CloudSQLIPConfiguration `json:"ipConfiguration"`
	LocationPreference        struct {
		Zone string `json:"zone"`
		Kind string `json:"kind"`
	} `json:"locationPreference"`
	DatabaseFlags []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"databaseFlags"`
	DataDiskType      string `json:"dataDiskType"`
	MaintenanceWindow struct {
		Kind string `json:"kind"`
		Hour int    `json:"hour"`
		Day  int    `json:"day"`
	} `json:"maintenanceWindow"`
	BackupConfiguration struct {
		StartTime               string `json:"startTime"`
		Kind                    string `json:"kind"`
		BackupRetentionSettings struct {
			RetentionUnit   string `json:"retentionUnit"`
			RetainedBackups int    `json:"retainedBackups"`
		} `json:"backupRetentionSettings"`
		Enabled                        bool   `json:"enabled"`
		ReplicationLogArchivingEnabled bool   `json:"replicationLogArchivingEnabled"`
		PointInTimeRecoveryEnabled     bool   `json:"pointInTimeRecoveryEnabled"`
		TransactionLogRetentionDays    int    `json:"transactionLogRetentionDays"`
		TransactionalLogStorageState   string `json:"transactionalLogStorageState"`
	} `json:"backupConfiguration"`
	ConnectorEnforcement      string `json:"connectorEnforcement"`
	SettingsVersion           string `json:"settingsVersion"`
	StorageAutoResizeLimit    string `json:"storageAutoResizeLimit"`
	StorageAutoResize         bool   `json:"storageAutoResize"`
	DataDiskSizeGb            string `json:"dataDiskSizeGb"`
	DeletionProtectionEnabled bool   `json:"deletionProtectionEnabled"`
}

type CloudSQLDatabase struct {
	Kind            string                   `json:"kind"`
	State           string                   `json:"state"`
	DatabaseVersion string                   `json:"databaseVersion"`
	Settings        CloudSQLDatabaseSettings `json:"settings"`
	Etag            string                   `json:"etag"`
	IPAddresses     []struct {
		Type      string `json:"type"`
		IPAddress string `json:"ipAddress"`
	} `json:"ipAddresses"`
	ServerCaCert                 CloudSQLServerCaCert `json:"serverCaCert"`
	InstanceType                 string               `json:"instanceType"`
	Project                      string               `json:"project"`
	ServiceAccountEmailAddress   string               `json:"serviceAccountEmailAddress"`
	BackendType                  string               `json:"backendType"`
	SelfLink                     string               `json:"selfLink"`
	ConnectionName               string               `json:"connectionName"`
	Name                         string               `json:"name"`
	Region                       string               `json:"region"`
	GceZone                      string               `json:"gceZone"`
	DatabaseInstalledVersion     string               `json:"databaseInstalledVersion"`
	AvailableMaintenanceVersions []string             `json:"availableMaintenanceVersions"`
	MaintenanceVersion           string               `json:"maintenanceVersion"`
	CreateTime                   time.Time            `json:"createTime"`
	SqlNetworkArchitecture       string               `json:"sqlNetworkArchitecture"`
	SatisfiesPzi                 bool                 `json:"satisfiesPzi"`
}

type CloudSQLDatabaseListResponse struct {
	Items []CloudSQLDatabase `json:"items"`
}
