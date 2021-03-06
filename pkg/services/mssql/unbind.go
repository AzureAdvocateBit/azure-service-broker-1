package mssql

import (
	"fmt"

	"github.com/Azure/azure-service-broker/pkg/service"
)

func (m *module) Unbind(
	provisioningContext service.ProvisioningContext, // nolint: unparam
	bindingContext service.BindingContext,
) error {
	pc, ok := provisioningContext.(*mssqlProvisioningContext)
	if !ok {
		return fmt.Errorf(
			"error casting provisioningContext as *mssqlProvisioningContext",
		)
	}
	bc, ok := bindingContext.(*mssqlBindingContext)
	if !ok {
		return fmt.Errorf(
			"error casting bindingContext as *mssqlBindingContext",
		)
	}

	// connect to new database to drop user for the login
	db, err := getDBConnection(pc, pc.DatabaseName)
	if err != nil {
		return err
	}
	defer db.Close() // nolint: errcheck

	if _, err = db.Exec(
		fmt.Sprintf("DROP USER \"%s\"", bc.LoginName),
	); err != nil {
		return fmt.Errorf(
			`error dropping user "%s": %s`,
			bc.LoginName,
			err,
		)
	}

	// connect to master database to drop login
	masterDb, err := getDBConnection(pc, "master")
	if err != nil {
		return err
	}
	defer masterDb.Close() // nolint: errcheck

	if _, err = masterDb.Exec(
		fmt.Sprintf("DROP LOGIN \"%s\"", bc.LoginName),
	); err != nil {
		return fmt.Errorf(
			`error dropping login "%s": %s`,
			bc.LoginName,
			err,
		)
	}

	return nil
}
