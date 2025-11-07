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

// Demo receive paramer
// {"Action":"Stop","Code":"FaceRecognition","Data":{"Address":"","Class":"FaceRecognition","Face":{"Age":25,"Beard":1,"BoundingBox":[4416,5784,4896,6840],"Center":[4656,6312],"Feature":["Neutral","NoGlasses"],"Frequency":null,"Glass":1,"Mask":1,"ObjectID":294,"ObjectType":"HumanFace","RelativeID":0,"Sex":"Man"},"IsGlobalScene":false,"Name":"FaceAnalysis","Object":{"Action":"Appear","Age":25,"Angle":[20,0,20],"ArgFS":1,"ArgResv2":2,"Attractive":48,"Beard":1,"BelongID":0,"BoundingBox":[4416,5784,4896,6840],"Center":[4656,6312],"Complexion":1,"Confidence":255,"Emotion":"Neutral","Express":6,"Eye":2,"FaceAlignScore":0,"FaceQuality":76,"Feature":["Neutral","NoGlasses"],"FeatureVersion":"1003002001002","FrameSequence":50457,"Gender":2,"Glass":1,"IsNewPassby":0,"Mask":1,"Mouth":1,"ObjectID":294,"ObjectType":"HumanFace","RelativeID":0,"SerialUUID":"","Sex":"Man","Source":0,"Speed":0,"SpeedTypeInternal":0,"Strabismus":0,"faceTripLineDirection":0,"installDiagnosticStat":{"Hight":1056,"PitchAngle":20,"Quality":76,"RollAngle":20,"Width":480,"YawAngle":0}},"RealUTC":1762497130,"RuleID":2},"Index":0}

type ReceiveDahuaCameraFaceRecognitionPayload struct {
	Action string `json:"Action"`
	Code   string `json:"Code"`
	Data   struct {
		Address       string `json:"Address"`
		Class         string `json:"Class"`
		Face          any    `json:"Face"`
		IsGlobalScene bool   `json:"IsGlobalScene"`
		Name          string `json:"Name"`
		Object        any    `json:"Object"`
		RealUTC       int64  `json:"RealUTC"`
		RuleID        int    `json:"RuleID"`
	} `json:"Data"`
	Index int `json:"Index"`
}
