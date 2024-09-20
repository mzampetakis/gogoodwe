/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/utils"
	"github.com/AaronSaikovski/gogoodwe/pkg/monitordetail"
	"github.com/alexflint/go-arg"
	"strconv"
	"strings"
)

type ResponseData struct {
	InverterTemperature float64 `json:"inverter_temperature"`
	PvOutput            float64 `json:"pv_output"`
	HouseLoad           float64 `json:"house_load"`
	BatteryPercentage   int     `json:"battery_percentage"`
	BatteryInput        float64 `json:"battery_input"`
	BatteryOutput       float64 `json:"battery_output"`
	BatteryState        string  `json:"battery_state"`
	GridInput           float64 `json:"grid_input"`
	GridOutput          float64 `json:"grid_output"`
	GridDirection       string  `json:"grid_direction"`
}

// Run is the main program runner.
//
// It takes a version string as a parameter and returns an error.
// The version string is used to set the build information.
// The function parses the command line arguments using the utils.Args struct.
// It checks if the email address and powerstation ID are in the correct format.
// If not, it fails with an error message.
// Finally, it calls the fetchData function to get data from the API and returns any errors.
//
// Parameters:
// - ctx: the context.Context object for cancellation and timeouts.
// - versionString: the version string used to set the build information.
//
// Returns:
// - error: an error if there was a problem with the input or fetching the data from the API.
func Run(ctx context.Context, versionString string) error {

	// Set version build info
	var args utils.Args
	args.SetVersion(versionString)

	// Parse args
	p := arg.MustParse(&args)

	// Check for valid email address input
	if !utils.CheckValidEmail(args.Account) {
		p.Fail("invalid email address format: should be 'user@somedomain.com'")
		return ctx.Err()

	}

	// Check for valid powerstation ID
	if !utils.CheckValidPowerstationID(args.PowerStationID) {
		p.Fail("invalid Powerstation ID format: should be 'XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX'")
		return ctx.Err()
	}

	data, err := fetchData(ctx, args.Account, args.Password, args.PowerStationID, args.ReportType)
	if err != nil {
		return ctx.Err()
	} else {
		ctx.Done()
		monitorData := monitordetail.MonitorData{}
		err = json.Unmarshal([]byte(data), &monitorData)
		if err != nil {
			return err
		}
		responseData := ResponseData{}

		if len(monitorData.Data.Inverter) > 0 {
			responseData.InverterTemperature = monitorData.Data.Inverter[0].Temperature
		}
		responseData.PvOutput = monitorData.Data.Kpi.Pac
		if len(monitorData.Data.Powerflow.Load) > 3 {
			batteryLoadText, _ := strings.CutSuffix(monitorData.Data.Powerflow.Load, "(W)")
			batteryLoadFloat, err := strconv.ParseFloat(batteryLoadText, 64)
			if nil == err {
				responseData.HouseLoad = batteryLoadFloat
			}
		}
		responseData.BatteryPercentage = monitorData.Data.Powerflow.Soc

		batteryText, _ := strings.CutSuffix(monitorData.Data.Powerflow.Battery, "(W)")
		batteryFloat, err := strconv.ParseFloat(batteryText, 64)

		if monitorData.Data.Powerflow.BatteryStatus == 1 {
			responseData.BatteryState = "charging"
			if nil == err {
				responseData.BatteryInput = batteryFloat
			}
		} else {
			responseData.BatteryState = "discharging"
			if nil == err {
				responseData.BatteryOutput = batteryFloat
			}
		}
		if monitorData.Data.Powerflow.Soc == 100 {
			responseData.BatteryState = "charged"
		}

		gridLoadText, _ := strings.CutSuffix(monitorData.Data.Powerflow.Grid, "(W)")
		gridLoadFloat, err := strconv.ParseFloat(gridLoadText, 64)
		if monitorData.Data.Powerflow.LoadStatus == 1 {
			responseData.GridDirection = "incoming"
			if nil == err {
				responseData.GridInput = gridLoadFloat
			}
		} else {
			responseData.GridDirection = "outgoing"
			if nil == err {
				responseData.GridOutput = gridLoadFloat
			}
		}

		result, _ := json.Marshal(&responseData)
		fmt.Print(string(result))
		return nil
	}

}
