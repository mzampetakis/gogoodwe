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

package utils

var (
	// Version string
	VersionString string = "gogoodwe v2.0.3"
)

// Args - struct using go-arg- https://github.com/alexflint/go-arg
type Args struct {
	Account        string `arg:"required,-a,--account" help:"SEMS Email Account."`
	Password       string `arg:"required,-p,--password" help:"SEMS Account password."`
	PowerStationID string `arg:"required,-i,--powerstationid" help:"SEMS Powerstation ID."`
	DailySummary   bool   `arg:"-s,--summary" help:"Output as a daily summary."`
}

// Description - App description
func (Args) Description() string {
	return "A command line tool to query the GOODWE SEMS Portal APIs and Solar SEMS API."
}

// Version - Version info
func (Args) Version() string {
	return VersionString
}

// set version string
func (Args) SetVersion(versionString string) {
	VersionString = versionString
}
