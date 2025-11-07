package schema

//   "type": "alert",
//   "severity": "critical",
//   "titleKey": "notis.fallDetected",
//   "title": "ตรวจพบคนล้มบริเวณ Lobby",
//   "img": "https://example.com/lobby/fall-frame.jpg",
//   "occurredAt": "2025-02-11T14:32:10+07:00",
//   "meta": {
//     "cameraName": "Lobby Camera",
//     "zone": "Lobby North",
//     "confidence": 0.91,
//     "clip": "https://example.com/lobby/fall-highlight.mp4"
//   }
// }

type HikvisionEmergencyAlarmEventMeta struct {
	CameraName string  `json:"cameraName"`
	Zone       string  `json:"zone"`
	Confidence float64 `json:"confidence"`
	Clip       string  `json:"clip"`
}

type HikvisionEmergencyAlarmEvent struct {
	Type       string                 `json:"type"`
	Severity   string                 `json:"severity"`
	TitleKey   string                 `json:"titleKey"`
	Title      string                 `json:"title"`
	Img        string                 `json:"img"`
	OccurredAt string                 `json:"occurredAt"`
	Meta       map[string]interface{} `json:"meta"`
}
