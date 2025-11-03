package schema

// Define necessary structs and types here for Dahua NVR schema
type DahuaNVRFaceRecognitionEvent struct {
	ID            string `json:"id"`
	Picture       string `json:"picture"`
	FullName      string `json:"fullName"`
	Gender        string `json:"gender"`
	Province      string `json:"province"`
	DateTimestamp string `json:"dateTimestamp"`
}
