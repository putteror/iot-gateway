package hivision

// {
//    "Action" : "Stop",
//    "Code" : "AlarmLocal",
//    "Data" : {
//       "EventID" : 228,
//       "Name" : "Nonamed",
//       "PTS" : 42949483440.0,
//       "RealUTC" : 1769506604,
//       "SenseMethod" : "",
//       "UTC" : 1769531804.0
//    },
//    "Index" : 0
// }

// DahuaCameraFaceRecognitionEventSchema - Struct หลักที่รวมทุกส่วน
type HikvisionCameraEmergencyAlarmEventSchema struct {
	Action string `json:"Action"`
	Code   string `json:"Code"`
	Data   struct {
		EventID     int     `json:"EventID"`
		Name        string  `json:"Name"`
		PTS         float64 `json:"PTS"`
		RealUTC     int     `json:"RealUTC"`
		SenseMethod string  `json:"SenseMethod"`
		UTC         float64 `json:"UTC"`
	} `json:"Data"`
	Index int `json:"Index"`
}
