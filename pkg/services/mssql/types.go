package mssql

import "github.com/Azure/azure-service-broker/pkg/service"

// ProvisioningParameters encapsulates MSSQL-specific provisioning options
type ProvisioningParameters struct {
	Location      string            `json:"location"`
	ServerName    string            `json:"server"`
	ResourceGroup string            `json:"resourceGroup"`
	EnableTDE     bool              `json:"enableTDE"`
	Tags          map[string]string `json:"tags"`
}

type mssqlProvisioningContext struct {
	ResourceGroupName          string            `json:"resourceGroup"`
	ARMDeploymentName          string            `json:"armDeployment"`
	ServerName                 string            `json:"server"`
	IsNewServer                bool              `json:"isNewServer"`
	Location                   string            `json:"location"`
	AdministratorLogin         string            `json:"administratorLogin"`
	AdministratorLoginPassword string            `json:"administratorLoginPassword"` // nolint: lll
	DatabaseName               string            `json:"database"`
	FullyQualifiedDomainName   string            `json:"fullyQualifiedDomainName"`
	Tags                       map[string]string `json:"tags"`
}

// UpdatingParameters encapsulates MSSQL-specific updating options
type UpdatingParameters struct {
	AdministratorLogin         string `json:"administratorLogin"`
	AdministratorLoginPassword string `json:"administratorLoginPassword"`
}

// BindingParameters encapsulates MSSQL-specific binding options
type BindingParameters struct {
}

type mssqlBindingContext struct {
	LoginName string `json:"loginName"`
}

// Credentials encapsulates MSSQL-specific coonection details and credentials.
type Credentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ServerConfig represents all configuration details needed for connecting to
// an Azure SQL Server.
type ServerConfig struct {
	ServerName                 string `json:"serverName"`
	ResourceGroupName          string `json:"resourceGroup"`
	Location                   string `json:"location"`
	AdministratorLogin         string `json:"administratorLogin"`
	AdministratorLoginPassword string `json:"administratorLoginPassword"`
}

// Config contains only a map of ServerConfig
type Config struct {
	Servers map[string]ServerConfig
}

func (
	m *module,
) GetEmptyProvisioningParameters() service.ProvisioningParameters {
	return &ProvisioningParameters{}
}

// SetResourceGroup sets the name of the resource group into which service
// instances will be deployed
func (p *ProvisioningParameters) SetResourceGroup(resourceGroup string) {
	p.ResourceGroup = resourceGroup
}

func (
	m *module,
) GetEmptyUpdatingParameters() service.UpdatingParameters {
	return &UpdatingParameters{}
}

func (m *module) GetEmptyProvisioningContext() service.ProvisioningContext {
	return &mssqlProvisioningContext{}
}

func (m *module) GetEmptyBindingParameters() service.BindingParameters {
	return &BindingParameters{}
}

func (m *module) GetEmptyBindingContext() service.BindingContext {
	return &mssqlBindingContext{}
}

func (m *module) GetEmptyCredentials() service.Credentials {
	return &Credentials{}
}
