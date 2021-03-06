package service

// Stability is a type that represents the relative stability of a service
// module
type Stability int

const (
	// StabilityAlpha represents relative stability of the most immature and
	// experimental service modules
	StabilityAlpha Stability = iota
	// StabilityBeta represents relative stability of the moderately immature and
	// semi-experimental service modules
	StabilityBeta
	// StabilityStable represents relative stability of the mature, production-
	// ready service modules
	StabilityStable
)

// ProvisioningParameters is an interface to be implemented by module-specific
// types that represent provisioning parameters. This interface doesn't require
// any functions to be implemented. It exists to improve the clarity of function
// signatures and documentation.
type ProvisioningParameters interface {
	// SetResourceGroup sets the name of the resource group that a service
	// instance should be provisioned into.
	SetResourceGroup(string)
}

// ProvisioningContext is an interface to be implemented by module-specific
// types that represent provisioning context.
type ProvisioningContext interface{}

// UpdatingParameters is an interface to be implemented by module-specific
// types that represent updating parameters. This interface doesn't require
// any functions to be implemented. It exists to improve the clarity of function
// signatures and documentation.
type UpdatingParameters interface{}

// BindingParameters is an interface to be implemented by module-specific types
// that represent binding parameters. This interface doesn't require any
// functions to be implemented. It exists to improve the clarity of function
// signatures and documentation.
type BindingParameters interface{}

// BindingContext is an interface to be implemented by module-specific types
// that represent binding context. This interface doesn't require any functions
// to be implemented. It exists to improve the clarity of function signatures
// and documentation.
type BindingContext interface{}

// Credentials is an interface to be implemented by module-specific types
// that represent service credentials. This interface doesn't require any
// functions to be implemented. It exists to improve the clarity of function
// signatures and documentation.
type Credentials interface{}
