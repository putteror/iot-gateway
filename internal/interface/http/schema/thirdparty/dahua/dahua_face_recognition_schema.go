package dahuschema

// DahuaCameraFaceRecognitionEventSchema - Struct หลักที่รวมทุกส่วน
type DahuaCameraFaceRecognitionEventSchema struct {
	Ack         bool             `json:"Ack"`
	Channel     int              `json:"Channel"`
	EOF         int              `json:"EOF"`
	Events      []DahuaEvent     `json:"Events"`
	Flags       []string         `json:"Flags"`
	Image       []DahuaImageInfo `json:"Image"`
	Length      int              `json:"Length"`
	PicID       []int            `json:"PicID"`
	PictureType int              `json:"PictureType"`
	Resolution  []int            `json:"Resolution"`
	SOF         int              `json:"SOF"`
	Time        string           `json:"Time"`
	Transfer    string           `json:"Transfer"`
}

// DahuaEvent - ส่วน Events[]
type DahuaEvent struct {
	Action string    `json:"Action"`
	Code   string    `json:"Code"`
	Data   DahuaData `json:"Data"`
	Index  int       `json:"Index"`
}

// DahuaData - ส่วน Data
type DahuaData struct {
	Address string `json:"Address"`
	// Candidates คือจุดที่ต่างกัน: มีข้อมูล (พบ) หรือเป็นอาร์เรย์ว่าง (ไม่พบ)
	Candidates     []DahuaCandidate   `json:"Candidates"`
	CfgRuleID      int                `json:"CfgRuleId"`
	Class          string             `json:"Class"`
	CountInGroup   int                `json:"CountInGroup"`
	DetectRegion   *interface{}       `json:"DetectRegion"` // รับค่า null
	EventID        int                `json:"EventID"`
	EventSeq       int                `json:"EventSeq"`
	Face           DahuaObjectData    `json:"Face"`
	Faces          []DahuaObjectData  `json:"Faces"`
	FeatureVector  DahuaFeatureVector `json:"FeatureVector"`
	FeatureVersion string             `json:"FeatureVersion"`
	Flag           int                `json:"Flag"`
	FrameSequence  int                `json:"FrameSequence"`
	GroupID        int                `json:"GroupID"`
	ImageInfo      DahuaImageInfo     `json:"ImageInfo"`
	IndexInGroup   int                `json:"IndexInGroup"`
	IsGlobalScene  bool               `json:"IsGlobalScene"`
	Mark           int                `json:"Mark"`
	Name           string             `json:"Name"`
	Object         DahuaObjectData    `json:"Object"`
	Objects        []DahuaObjectData  `json:"Objects"`
	PTS            float64            `json:"PTS"`
	Priority       int                `json:"Priority"`
	RealUTC        int                `json:"RealUTC"`
	RuleID         int                `json:"RuleID"`
	RuleId         int                `json:"RuleId"`
	Source         float64            `json:"Source"`
	UTC            int                `json:"UTC"`
	UTCMS          int                `json:"UTCMS"`
}

// DahuaCandidate - ผู้ถูกเสนอชื่อ (มีเฉพาะตอนที่พบ)
type DahuaCandidate struct {
	Person     DahuaPerson `json:"Person"`
	Similarity int         `json:"Similarity"`
}

// DahuaPerson - ข้อมูลบุคคลที่ถูกจดจำ (มีเฉพาะตอนที่พบ)
type DahuaPerson struct {
	Age             int              `json:"Age"`
	Birthday        string           `json:"Birthday"`
	CertificateType string           `json:"CertificateType"`
	City            string           `json:"City"`
	Comment         string           `json:"Comment"`
	Country         string           `json:"Country"`
	FeatureErrCode  int              `json:"FeatureErrCode"`
	FeatureState    int              `json:"FeatureState"`
	GroupAttribute  string           `json:"GroupAttribute"`
	GroupID         string           `json:"GroupID"`
	GroupName       string           `json:"GroupName"`
	HomeAddress     string           `json:"HomeAddress"`
	ID              string           `json:"ID"`
	Image           []DahuaImageInfo `json:"Image"` // ใช้ DahuaImageInfo ซ้ำ
	Imagepath       string           `json:"Imagepath"`
	Important       int              `json:"Important"`
	IsCustomType    int              `json:"IsCustomType"`
	Name            string           `json:"Name"`
	Province        string           `json:"Province"`
	Sex             string           `json:"Sex"`
	TimeSection     string           `json:"TimeSection"`
	Type            string           `json:"Type"`
	UID             string           `json:"UID"`
}

// DahuaObjectData - โครงสร้างของ Face, Faces[], Object, Objects[]
type DahuaObjectData struct {
	Action                string                     `json:"Action"`
	Age                   int                        `json:"Age"`
	Angle                 []int                      `json:"Angle"`
	ArgFS                 int                        `json:"ArgFS"`
	ArgResv2              int                        `json:"ArgResv2"`
	Attractive            int                        `json:"Attractive"`
	Beard                 int                        `json:"Beard"`
	BelongID              int                        `json:"BelongID"`
	BoundingBox           []int                      `json:"BoundingBox"`
	Category              string                     `json:"Category,omitempty"`
	Center                []int                      `json:"Center"`
	Complexion            int                        `json:"Complexion,omitempty"`
	Confidence            int                        `json:"Confidence,omitempty"`
	Emotion               string                     `json:"Emotion,omitempty"`
	Express               int                        `json:"Express,omitempty"`
	Eye                   int                        `json:"Eye,omitempty"`
	FaceAlignScore        int                        `json:"FaceAlignScore,omitempty"`
	FaceQuality           int                        `json:"FaceQuality,omitempty"`
	Feature               []string                   `json:"Feature"`
	FrameSequence         int                        `json:"FrameSequence,omitempty"`
	Gender                int                        `json:"Gender,omitempty"`
	Glass                 int                        `json:"Glass"`
	Image                 DahuaImageInfo             `json:"Image,omitempty"`
	Mask                  int                        `json:"Mask"`
	Mouth                 int                        `json:"Mouth,omitempty"`
	ObjectID              int                        `json:"ObjectID"`
	ObjectType            string                     `json:"ObjectType"`
	Recnum                int                        `json:"Recnum,omitempty"`
	Recresult             []DahuaRecResult           `json:"Recresult,omitempty"`
	RelativeID            int                        `json:"RelativeID"`
	SerialUUID            string                     `json:"SerialUUID,omitempty"`
	Sex                   string                     `json:"Sex"`
	Source                float32                    `json:"Source,omitempty"`
	Speed                 int                        `json:"Speed,omitempty"`
	SpeedTypeInternal     int                        `json:"SpeedTypeInternal,omitempty"`
	Strabismus            int                        `json:"Strabismus,omitempty"`
	faceTripLineDirection int                        `json:"faceTripLineDirection,omitempty"`
	InstallDiagnosticStat DahuaInstallDiagnosticStat `json:"installDiagnosticStat,omitempty"`
}

// DahuaRecResult - ผลลัพธ์การจดจำ
type DahuaRecResult struct {
	DbID       int `json:"DbId"`
	DbType     int `json:"DbType"`
	FeatureID  int `json:"FeatureId"`
	Similarity int `json:"Similarity"`
}

// DahuaInstallDiagnosticStat - ข้อมูลสถานะการติดตั้ง
type DahuaInstallDiagnosticStat struct {
	Hight      int `json:"Hight"`
	PitchAngle int `json:"PitchAngle"`
	Quality    int `json:"Quality"`
	RollAngle  int `json:"RollAngle"`
	Width      int `json:"Width"`
	YawAngle   int `json:"YawAngle"`
}

// DahuaFeatureVector - ข้อมูล Feature Vector
type DahuaFeatureVector struct {
	Length int `json:"Length"`
	Offset int `json:"Offset"`
}

// DahuaImageInfo - ข้อมูลภาพ (ใช้ใน Image[] และ Data.ImageInfo, Person.Image[])
type DahuaImageInfo struct {
	IndexInData int `json:"IndexInData"`
	Length      int `json:"Length"`
	Offset      int `json:"Offset"`
	// เฉพาะใน Event.Image บางตัว
	Height int    `json:"Height,omitempty"`
	Width  int    `json:"Width,omitempty"`
	Type   string `json:"Type,omitempty"`
}
