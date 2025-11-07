package schema

type SendDahuaCameraFaceRecognitionPayload struct {
	Type       string `json:"type"`
	Severity   string `json:"severity"`
	TitleKey   string `json:"titleKey"`
	OccurredAt string `json:"occurredAt"`
	Meta       struct {
		Kind   string `json:"kind"`
		Person struct {
			FullName string `json:"fullName"`
			Gender   string `json:"gender"`
		} `json:"person"`
		RawID string `json:"rawId"`
	} `json:"meta"`
}

// type ReceiveDahuaCameraFaceRecognitionPayload struct {
// 	SiteID      string `json:"siteId"`
// 	CameraName  string `json:"cameraName"`
// 	PersonName  string `json:"personName"`
// 	G
