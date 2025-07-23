package main

import "time"

// V1Devices is the exported struct for https://api.ui.com/v1/devices
type V1Devices struct {
	Data []struct {
		HostID   string `json:"hostId"`
		HostName string `json:"hostName"`
		Devices  []struct {
			ID              string    `json:"id"`
			Mac             string    `json:"mac"`
			Name            string    `json:"name"`
			Model           string    `json:"model"`
			Shortname       string    `json:"shortname"`
			IP              string    `json:"ip"`
			ProductLine     string    `json:"productLine"`
			Status          string    `json:"status"`
			Version         string    `json:"version"`
			FirmwareStatus  string    `json:"firmwareStatus"`
			UpdateAvailable any       `json:"updateAvailable"`
			IsConsole       bool      `json:"isConsole"`
			IsManaged       bool      `json:"isManaged"`
			StartupTime     time.Time `json:"startupTime"`
			AdoptionTime    time.Time `json:"adoptionTime"`
			Note            any       `json:"note"`
			Uidb            struct {
				GUID   string `json:"guid"`
				IconID string `json:"iconId"`
				ID     string `json:"id"`
				Images struct {
					Default   string `json:"default"`
					Nopadding string `json:"nopadding"`
					Topology  string `json:"topology"`
				} `json:"images"`
			} `json:"uidb"`
		} `json:"devices"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"data"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	TraceID        string `json:"traceId"`
}
