package schema

import "time"

type EmergencyAlarmEventSchema struct {
	Type              string    `json:"type"`
	StampDateTime     time.Time `json:"stampDateTime"`
	SiteID            string    `json:"siteId"`
	DeviceInformation struct {
		ID         string `json:"id"`
		Name       string `json:"Name"`
		IPAddress  string `json:"ipAddress"`
		MACAddress string `json:"macAddress"`
	}
}
