package service

import (
	"fmt"

	"github.com/Azure/azure-service-broker/pkg/crypto/noop"
)

type ArbitraryType struct {
	Foo string `json:"foo"`
}

// SetResourceGroup is implemented only so that ArbitraryType will fulfill
// the ProvisioningParameters interface. This function isn't used anywhere.
func (a *ArbitraryType) SetResourceGroup(resourceGroup string) {}

const fooValue = "bar"

var (
	testArbitraryObject = &ArbitraryType{
		Foo: fooValue,
	}
	testArbitraryObjectJSON = []byte(fmt.Sprintf(`{"foo":"%s"}`, fooValue))
	noopCodec               = noop.NewCodec()
)
