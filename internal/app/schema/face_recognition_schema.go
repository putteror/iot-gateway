package schema

import "time"

type FaceRecognitionEventSchema struct {
	Success           bool      `json:"success"`
	Type              string    `json:"type"`
	ImageBase64       string    `json:"imageBase64"`
	StampDateTime     time.Time `json:"stampDateTime"`
	Confidence        int       `json:"confidence"`
	PersonInformation struct {
		ID         string `json:"id"`
		FirstName  string `json:"firstName"`
		MiddleName string `json:"middleName"`
		LastName   string `json:"lastName"`
		Gender     string `json:"gender"`
		Age        int    `json:"age"`
	}
	DeviceInformation struct {
		ID         string `json:"id"`
		Name       string `json:"Name"`
		IPAddress  string `json:"ipAddress"`
		MACAddress string `json:"macAddress"`
	}
}
