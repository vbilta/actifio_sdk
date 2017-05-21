package main
import (
	"actifio/sdk"
	"fmt"
)


func main () {
	sdk.StartUp()
	connection := sdk.ActConnection{
		Name: sdk.ActConfig.Name,
		Password: sdk.ActConfig.Password,
		VendorKey: sdk.ActConfig.VendorKey,
		Appliance: sdk.ActConfig.Appliance,
	}
	key, err := connection.Login()
	if err != nil {
		fmt.Println("Couldn't login to actifio: %s", err)
	}

	pools, err := connection.GetDiskPools(key.SessionId)
	if err != nil {
		fmt.Println("Error: %s", err)
	}
	for _, pool := range pools {
		fmt.Println(pool.Name)
	}
}
