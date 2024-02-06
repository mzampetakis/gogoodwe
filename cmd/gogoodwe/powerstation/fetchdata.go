/*
# Name: data - fetches data from the goodwe API - and processes it to pass back to caller
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package powerstation

import (
	"errors"
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/semsapi"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/logrusorgru/aurora"
)

//Old ver
// FetchData - Main API fetch function
// func FetchData(Account string, Password string, PowerStationID string) error {

// 	// Data types
// 	var SemsResponseData types.LoginResponse
// 	var PowerstationData types.InverterData

// 	// Create a new SemsLoginCreds object via a struct literal
// 	var SemsUserLogin = types.LoginCredentials{
// 		Account:        Account,
// 		Password:       Password,
// 		PowerStationID: PowerStationID,
// 	}

// 	// Do the login..check for errors
// 	err := semsapi.ApiLogin(&SemsUserLogin, &SemsResponseData)
// 	if err == nil {

// 		// Fetch the data
// 		dataerr := fetchInverterData(&SemsResponseData, &SemsUserLogin, &PowerstationData)
// 		if dataerr != nil {
// 			utils.HandleError(errors.New("error: fetching powerstation data, check powerstationid is correct"))
// 			return dataerr
// 		} else {
// 			// Get output
// 			dataOutput, jsonerr := utils.GetDataJSON(&PowerstationData)
// 			if jsonerr != nil {
// 				utils.HandleError(errors.New("error: converting powerstation data"))
// 				return jsonerr

// 			} else {
// 				//Display output
// 				fmt.Println(aurora.BrightYellow(string(dataOutput)))
// 			}
// 		}

// 	} else {
// 		utils.HandleError(err)
// 		return err
// 	}

// 	return nil
// }

func FetchData(Account string, Password string, PowerStationID string) error {

	// Data types
	var PowerstationData types.InverterData

	// User account struct
	creds := &types.LoginCredentials{
		Account:        Account,
		Password:       Password,
		PowerStationID: PowerStationID,
	}

	// Login API Response object
	resp := &types.LoginResponse{}

	// Create a new LoginDataFlow object reference
	LoginDataFlow := &types.LoginDataFlow{
		LoginCreds: creds,
		LoginResp:  resp,
	}

	// Do the login..check for errors
	err := semsapi.ApiLoginV2(LoginDataFlow)
	if err == nil {

		// Fetch the data
		dataerr := fetchInverterData(LoginDataFlow, &PowerstationData)
		if dataerr != nil {
			utils.HandleError(errors.New("error: fetching powerstation data, check powerstationid is correct"))
			return dataerr
		} else {
			// Get output
			dataOutput, jsonerr := utils.GetDataJSON(&PowerstationData)
			if jsonerr != nil {
				utils.HandleError(errors.New("error: converting powerstation data"))
				return jsonerr

			} else {
				//Display output
				fmt.Println(aurora.BrightYellow(string(dataOutput)))
			}
		}

	} else {
		utils.HandleError(err)
		return err
	}

	return nil
}
