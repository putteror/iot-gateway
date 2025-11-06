package schema

// {
//     "siteId" : "7e56ca2b-b741-4eba-9fe7-3b30222ecc28",
//     "type" : "emergency",
//     "description" : "alarm trigger",
//     "status" : "active",
//     "img" : null,
//     "alarmDateTime" : "2025-11-05T02:20:02Z"
// }

type HikvisionEmergencyAlarmEvent struct {
	SiteID        string `json:"siteId"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	Img           string `json:"img"`
	AlarmDateTime string `json:"alarmDateTime"`
}
