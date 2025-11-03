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

type DahuaNVRLicensePlateRecognitionEvent struct {
	ID               string   `json:"id"`
	Picture          string   `json:"picture"`
	PlatePicture     string   `json:"platePicture"`
	PlateText        string   `json:"plateText"`
	Province         string   `json:"province"`
	ConfidenceHeader []string `json:"confidenceHeader"`
	CameraName       string   `json:"cameraName"`
	DateTimestamp    string   `json:"dateTimestamp"`
}
