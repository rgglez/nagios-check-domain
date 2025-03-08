/*
Copyright 2025 Rodolfo González González

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kr/pretty"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/spf13/pflag"
	"github.com/xorpaul/go-nagios"
	"github.com/ztrue/tracerr"
)

//-----------------------------------------------------------------------------

// CLI parameters
var (
	domain  string
	warn    int
	crit    int
	servers string
)

//-----------------------------------------------------------------------------

// Init!
func init() {
	// Parse command line parameters
	pflag.StringVarP(&domain, "domain", "D", "", "Domain to check")
	pflag.IntVarP(&warn, "warn", "w", 30, "Warning threshold in days")
	pflag.IntVarP(&crit, "crit", "c", 15, "Critical threshold days")
	pflag.StringVarP(&servers, "servers", "s", "", "Path to the whois servers list file")
	pflag.Parse()
}

//-----------------------------------------------------------------------------

// Function to extract the date part from the given timestamp string
func extractDate(timestamp string) string {
	// Split the string at "T" and take the first part
	parts := strings.Split(timestamp, "T")
	if len(parts) > 0 {
		return parts[0]
	}
	return "" // Return an empty string if the format is invalid
}

//-----------------------------------------------------------------------------

// Function to calculate the difference in days between a given date and the current date
func daysDifference(targetDate string) (int, error) {
	// Parse the target date in the format "YYYY-MM-DD"
	target, err := time.Parse("2006-01-02", targetDate)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %v", err)
	}

	// Get the current date (without time)
	current := time.Now().UTC().Truncate(24 * time.Hour)

	// Calculate the difference in days
	diff := target.Sub(current).Hours() / 24

	// Return the difference as an integer
	return int(diff), nil
}

//-----------------------------------------------------------------------------

// Run, Lola, Run!
func main() {
	// No domain given
	if domain == "" {
		log.Fatal("Domain argument is required")
	}

	// Get the whois server for the given domain
	ws := NewWhoisServers(servers)
	var server string
	var exists bool
	if server, exists = ws.GetWhoisServer(domain); !exists {
		server = "whois.iana.org"
	}

	// Query the whois servers
	raw, err := whois.Whois(domain, server)

	pretty.Println(raw)

	if err != nil {
		tracerr.PrintSource(err)
		nr := nagios.NagiosResult{
			ExitCode: 3,
			Text:     fmt.Sprintf("Whois query failed: %s", err.Error()),
			Perfdata: "",
		}
		nagios.NagiosExit(nr)
	}

	// Parse the whois raw response data
	result, err := whoisparser.Parse(raw)
	if err != nil {
		tracerr.PrintSource(err)
		nr := nagios.NagiosResult{
			ExitCode: 3,
			Text:     fmt.Sprintf("Whois output could not be parsed: %s", err.Error()),
			Perfdata: "",
		}
		nagios.NagiosExit(nr)
	}

	// Get the expiration date
	date := extractDate(result.Domain.ExpirationDate)

	// Calculate the days left until the domain expiration
	daysLeft, err := daysDifference(date)
	if err != nil {
		tracerr.PrintSource(err)
		nr := nagios.NagiosResult{
			ExitCode: 3,
			Text:     fmt.Sprintf("Could not calculate date difference: %s", err.Error()),
			Perfdata: "",
		}
		nagios.NagiosExit(nr)
	}

	// Prepare performance data
	perfdata := fmt.Sprintf("expires=%s", result.Domain.ExpirationDate)

	// Return variables for Nagios
	var exitCode int
	var statusText string

	// Set the status for Nagios
	if daysLeft <= crit {
		exitCode = 2 // Critical
	} else if daysLeft <= warn {
		exitCode = 1 // Warning
	} else {
		exitCode = 0 // OK
	}

	// Set the message for Nagios, considering if the domain has already expired.
	if daysLeft > 0 {
		statusText = fmt.Sprintf("Domain will expire in %d days", daysLeft)
	} else {
		statusText = fmt.Sprintf("Domain has expired %d days ago on %s", daysLeft*-1, date)
	}

	// Return Nagios result
	nr := nagios.NagiosResult{
		ExitCode: exitCode,
		Text:     statusText,
		Perfdata: perfdata,
	}
	nagios.NagiosExit(nr)
}
