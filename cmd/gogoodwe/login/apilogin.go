/*
# Name: apiLogin - Logs in to the SEMS API
# Author: Aaron Saikovski - asaikovski@outlook.com
*/
package login

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/authentication"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
)

// apiLogin -  Login to the API
func ApiLogin(UserLoginCreds *types.LoginCredentials, UserLoginResponse *types.LoginResponse) error {

	// Do the login - update the pointer to the struct SemsResponseData
	err := authentication.DoLogin(UserLoginResponse, UserLoginCreds)
	if err != nil {
		utils.HandleError(err)
		return err
	} else {
		return nil
	}
}
