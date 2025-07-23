package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationName string = "unifi-golang-test"
const applicationDescription string = "A simple Unifi API client in Go"
const applicationVersion string = "v0.0.1"

// a struct that hosts name, url, and method for the API call
type APIRequest struct {
	Name        string
	URL         string
	Method      string
	Description string
}

var unifiKey string
var debug bool

// Define time intervals for API requests
var timeIntervals = []string{"5m", "1h"}

// create an array of APIRequest structs
var apiRequests = []APIRequest{
	{Name: "GetDevices", URL: "https://api.ui.com/v1/devices", Method: "GET", Description: "Fetches devices from the Unifi API"},
	{Name: "GetSites", URL: "https://api.ui.com/v1/sites", Method: "GET", Description: "Fetches sites from the Unifi API"},
	/*{Name: "GetUsers", URL: "https://api.ui.com/v1/users", Method: "GET", Description: "Fetches users from the Unifi API"},
	{Name: "GetEvents", URL: "https://api.ui.com/v1/events", Method: "GET", Description: "Fetches events from the Unifi API"},
	{Name: "GetInsights", URL: "https://api.ui.com/v1/insights", Method: "GET", Description: "Fetches insights from the Unifi API"},
	{Name: "GetConfig", URL: "https://api.ui.com/v1/config", Method: "GET", Description: "Fetches configuration from the Unifi API"},
	{Name: "GetStatistics", URL: "https://api.ui.com/v1/statistics", Method: "GET", Description: "Fetches statistics from the Unifi API"},
	{Name: "GetAlerts", URL: "https://api.ui.com/v1/alerts", Method: "GET", Description: "Fetches alerts from the Unifi API"},
	{Name: "GetFirmware", URL: "https://api.ui.com/v1/firmware", Method: "GET", Description: "Fetches firmware information from the Unifi API"},
	{Name: "GetNetwork", URL: "https://api.ui.com/v1/network", Method: "GET", Description: "Fetches network information from the Unifi API"},
	{Name: "GetSettings", URL: "https://api.ui.com/v1/settings", Method: "GET", Description: "Fetches settings from the Unifi API"},
	{Name: "GetLogs", URL: "https://api.ui.com/v1/logs", Method: "GET", Description: "Fetches logs from the Unifi API"},
	{Name: "GetNotifications", URL: "https://api.ui.com/v1/notifications", Method: "GET", Description: "Fetches notifications from the Unifi API"},
	{Name: "GetSystem", URL: "https://api.ui.com/v1/system", Method: "GET", Description: "Fetches system information from the Unifi API"},
	{Name: "GetHealth", URL: "https://api.ui.com/v1/health", Method: "GET", Description: "Fetches health status from the Unifi API"},
	{Name: "GetPerformance", URL: "https://api.ui.com/v1/performance", Method: "GET", Description: "Fetches performance metrics from the Unifi API"},
	{Name: "GetSupport", URL: "https://api.ui.com/v1/support", Method: "GET", Description: "Fetches support information from the Unifi API"},
	{Name: "GetDiagnostics", URL: "https://api.ui.com/v1/diagnostics", Method: "GET", Description: "Fetches diagnostics from the Unifi API"},
	{Name: "GetBackup", URL: "https://api.ui.com/v1/backup", Method: "GET", Description: "Fetches backup information from the Unifi API"},
	{Name: "GetRestore", URL: "https://api.ui.com/v1/restore", Method: "GET", Description: "Fetches restore information from the Unifi API"},
	{Name: "GetUpdates", URL: "https://api.ui.com/v1/updates", Method: "GET", Description: "Fetches updates from the Unifi API"},
	{Name: "GetFeatures", URL: "https://api.ui.com/v1/features", Method: "GET", Description: "Fetches features from the Unifi API"},
	{Name: "GetLicenses", URL: "https://api.ui.com/v1/licenses", Method: "GET", Description: "Fetches licenses from the Unifi API"},
	{Name: "GetSubscriptions", URL: "https://api.ui.com/v1/subscriptions", Method: "GET", Description: "Fetches subscriptions from the Unifi API"},
	{Name: "GetIntegrations", URL: "https://api.ui.com/v1/integrations", Method: "GET", Description: "Fetches integrations from the Unifi API"},
	{Name: "GetAnalytics", URL: "https://api.ui.com/v1/analytics", Method: "GET", Description: "Fetches analytics from the Unifi API"},
	{Name: "GetReports", URL: "https://api.ui.com/v1/reports", Method: "GET", Description: "Fetches reports from the Unifi API"},
	{Name: "GetAlertsHistory", URL: "https://api.ui.com/v1/alerts/history", Method: "GET", Description: "Fetches alert history from the Unifi API"},
	{Name: "GetDeviceGroups", URL: "https://api.ui.com/v1/device-groups", Method: "GET", Description: "Fetches device groups from the Unifi API"},
	{Name: "GetDeviceConfigurations", URL: "https://api.ui.com/v1/device-configurations", Method: "GET", Description: "Fetches device configurations from the Unifi API"},
	{Name: "GetDeviceFirmware", URL: "https://api.ui.com/v1/device-firmware", Method: "GET", Description: "Fetches device firmware from the Unifi API"},
	{Name: "GetDeviceLogs", URL: "https://api.ui.com/v1/device-logs", Method: "GET", Description: "Fetches device logs from the Unifi API"},
	{Name: "GetDeviceStatistics", URL: "https://api.ui.com/v1/device-statistics", Method: "GET", Description: "Fetches device statistics from the Unifi API"},
	{Name: "GetDeviceSettings", URL: "https://api.ui.com/v1/device-settings", Method: "GET", Description: "Fetches device settings from the Unifi API"},
	{Name: "GetDeviceHealth", URL: "https://api.ui.com/v1/device-health", Method: "GET", Description: "Fetches device health from the Unifi API"},
	{Name: "GetDevicePerformance", URL: "https://api.ui.com/v1/device-performance", Method: "GET", Description: "Fetches device performance from the Unifi API"},
	{Name: "GetDeviceSupport", URL: "https://api.ui.com/v1/device-support", Method: "GET", Description: "Fetches device support information from the Unifi API"},
	{Name: "GetDeviceDiagnostics", URL: "https://api.ui.com/v1/device-diagnostics", Method: "GET", Description: "Fetches device diagnostics from the Unifi API"},
	{Name: "GetDeviceBackup", URL: "https://api.ui.com/v1/device-backup", Method: "GET", Description: "Fetches device backup information from the Unifi API"},
	{Name: "GetDeviceRestore", URL: "https://api.ui.com/v1/device-restore", Method: "GET", Description: "Fetches device restore information from the Unifi API"},
	{Name: "GetDeviceUpdates", URL: "https://api.ui.com/v1/device-updates", Method: "GET", Description: "Fetches device updates from the Unifi API"},
	{Name: "GetDeviceFeatures", URL: "https://api.ui.com/v1/device-features", Method: "GET", Description: "Fetches device features from the Unifi API"},
	{Name: "GetDeviceLicenses", URL: "https://api.ui.com/v1/device-licenses", Method: "GET", Description: "Fetches device licenses from the Unifi API"},
	{Name: "GetDeviceSubscriptions", URL: "https://api.ui.com/v1/device-subscriptions", Method: "GET", Description: "Fetches device subscriptions from the Unifi API"},
	{Name: "GetDeviceIntegrations", URL: "https://api.ui.com/v1/device-integrations", Method: "GET", Description: "Fetches device integrations from the Unifi API"},
	{Name: "GetDeviceAnalytics", URL: "https://api.ui.com/v1/device-analytics", Method: "GET", Description: "Fetches device analytics from the Unifi API"},
	{Name: "GetDeviceReports", URL: "https://api.ui.com/v1/device-reports", Method: "GET", Description: "Fetches device reports from the Unifi API"},
	{Name: "GetDeviceAlerts History", URL: "https://api.ui.com/v1/device-alerts/history", Method: "GET", Description: "Fetches device alert history from the Unifi API"},
	*/
}

func init() {
	flag.Bool("debug", false, "Enable debug mode")
	flag.Bool("help", false, "Display help")
	flag.String("action", "", "Specify action to perform")
	flag.String("interval", "5m", "Specify time interval for API requests (e.g., 5m, 1h)")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetBool("help") {
		displayHelp()
		os.Exit(0)
	}

	debug = viper.GetBool("debug")

	if !viper.IsSet("action") {
		fmt.Println("No action specified. Use --action to specify an action.")
		os.Exit(1)
	}

	unifiKey = os.Getenv("UNIFI_KEY")
	if unifiKey == "" {
		fmt.Println("UNIFI_KEY environment variable not set")
		os.Exit(1)
	}

	// if debug is enabled, be verbose
	if debug {
		fmt.Println("Using Unifi API Key:", unifiKey)
		fmt.Println("Action:", viper.GetString("action"))

		// Sort apiRequests by Name alphabetically
		sort.Slice(apiRequests, func(i, j int) bool {
			return apiRequests[i].Name < apiRequests[j].Name
		})

		// Create a new tab writer
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(writer, "Action\tURL\tMethod\tDescription")

		// Display the sorted API requests
		for _, req := range apiRequests {
			fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", req.Name, req.URL, req.Method, req.Description)
		}

		// write out the table
		writer.Flush()

	}

	// check if the interval is valid
	if !checkInterval(viper.GetString("interval")) {
		fmt.Println("Invalid interval specified. Use --interval to specify a valid interval (default = 5m), either 5m or 1h.")
		os.Exit(1)
	}

}

func main() {

	// check if the action is valid
	if !checkAction(viper.GetString("action")) {
		fmt.Println("Invalid action specified. Use --action to specify a valid action.")
		os.Exit(1)
	} else {
		if debug {
			fmt.Println("Valid action specified:", viper.GetString("action"))
		}
	}

	findApiIndex := -1
	for i, req := range apiRequests {
		if strings.ToLower(req.Name) == strings.ToLower(viper.GetString("action")) {
			findApiIndex = i
			break
		}
	}
	if findApiIndex == -1 {
		fmt.Println("Action " + viper.GetString("action") + " not found in apiRequests")
		os.Exit(1)
	}

	if debug {
		fmt.Printf(viper.GetString("action")+" is at position: %d\n", findApiIndex)
	}

	returnValue, returnString := callAPI(apiRequests[0].URL, apiRequests[0].Method)

	if !returnValue {
		fmt.Println("Failed to call API:", apiRequests[0].Name)
		os.Exit(1)
	}

	switch strings.ToLower(viper.GetString("action")) {
	case "getdevices":
		if debug {
			fmt.Println("Printing results for GetDevices..." + returnString + "__")
		}

		var devices V1Devices
		err := json.Unmarshal([]byte(returnString), &devices)
		if err != nil {
			fmt.Println("Error unmarshalling response:", err)
			os.Exit(1)
		}

		// Count total devices across all hosts
		totalDevices := 0
		for _, host := range devices.Data {
			totalDevices += len(host.Devices)
		}
		fmt.Println("Devices count:", totalDevices)

		if debug {
			fmt.Println("\n\n")
			for _, host := range devices.Data {
				//fmt.Printf("Host: %s (%s)\n", host.HostName, host.HostID)
				fmt.Println(prettyPrint(host.Devices))
			}
			fmt.Println("")
		}

		// Create a new tab writer
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(writer, "#\tMAC\tName\tModel\tIP\tStatus\tVersion\tFirmware\tMngd\tStartup Time")

		deviceIndex := 1
		for _, host := range devices.Data {
			for _, device := range host.Devices {
				//fmt.Printf("Device %d: MAC: %s, Name: %s\n", deviceIndex, device.Mac, device.Name)
				fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%t\t%s\n", deviceIndex, device.Mac, device.Name, device.Model, device.IP, device.Status, device.Version, device.FirmwareStatus, device.IsManaged, device.StartupTime.Format(time.RFC3339))
				deviceIndex++
			}
		}

		// write out the table
		writer.Flush()

	case "getsites":
		if debug {
			fmt.Println("Printing results for GetSites... " + returnString + "__")
		}
	// add more cases as needed for other actions
	default:
		fmt.Println("Unknown action:", viper.GetString("action"))
		os.Exit(1)
	}
}

// display the help message
func displayHelp() {

	fmt.Println(applicationName + " " + applicationVersion + "\n" + applicationDescription + "\n")
	flags := []struct {
		Name        string
		Description string
	}{
		{"--debug", "Enable debug mode"},
		{"--help", "Display help"},
		{"--action", "Specify action to perform"},
		{"--interval", "Specify time interval for API requests (e.g., 5m, 1h)"},
	}

	fmt.Println("Flags:")
	for _, flag := range flags {
		fmt.Printf("  %-15s %s\n", flag.Name, flag.Description)
	}

}

// calls an API with the provided URL and method
// returns true if the call was successful, false otherwise
// prints out the result
func callAPI(url string, method string) (bool, string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, ""
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", unifiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false, ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return false, ""
	}
	//fmt.Println(string(body))
	return res.StatusCode >= 200 && res.StatusCode < 300, string(body) // Return the body for further processing if needed
}

// checks if the provided action is valid
func checkAction(action string) bool {
	for _, req := range apiRequests {
		if strings.ToLower(req.Name) == strings.ToLower(action) {
			return true
		}
	}
	return false
}

// checks if the provided interval is valid
func checkInterval(interval string) bool {
	for _, t := range timeIntervals {
		if t == interval {
			return true
		}
	}
	return false
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
