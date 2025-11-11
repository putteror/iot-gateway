package adapter

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type GetWaterSensor interface {
	GetWaterData() WaterMeterData
}

type GetWaterSensorImpl struct {
}

func NewGetWaterSensorImpl() GetWaterSensor {
	return &GetWaterSensorImpl{}
}

type DomesticData struct {
	PH                float64 `json:"ph"`
	FlowRateLpm       float64 `json:"flowRateLpm"`
	ConsumptionLiters int     `json:"consumptionLiters"`
}

// DrinkingData สำหรับข้อมูลน้ำดื่ม
type DrinkingData struct {
	PH                float64 `json:"ph"`
	TdsPpm            int     `json:"tdsPpm"`
	ConsumptionLiters int     `json:"consumptionLiters"`
}

// TotalUsage สำหรับค่ารวม Today, Month, Year
type TotalUsage struct {
	Today int `json:"today"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// Totals สำหรับข้อมูลรวมการใช้งานน้ำใช้และน้ำดื่ม
type Totals struct {
	Domestic TotalUsage `json:"domestic"`
	Drinking TotalUsage `json:"drinking"`
}

// SeriesData สำหรับข้อมูลชุดอนุกรมภายใน StackedSeries
type SeriesData struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

// StackedSeries สำหรับข้อมูลกราฟรายเดือน
type StackedSeries struct {
	Categories []string     `json:"categories"`
	Series     []SeriesData `json:"series"`
}

// Radial สำหรับข้อมูลกราฟวงกลม/มาตรวัด
type Radial struct {
	TotalLiters int      `json:"totalLiters"`
	Labels      []string `json:"labels"`
	Values      []int    `json:"values"`
}

// UsageTimeline สำหรับข้อมูล Timeline รายชั่วโมง/รายวัน
type UsageTimeline struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

// --- Struct หลัก ---

// WaterMeterData เป็น Struct หลักที่รวมข้อมูลทั้งหมดของ Water Meter
type WaterMeterData struct {
	DeviceKey     string          `json:"deviceKey"`
	Timestamp     string          `json:"timestamp"`
	Domestic      DomesticData    `json:"domestic"`
	Drinking      DrinkingData    `json:"drinking"`
	Totals        Totals          `json:"totals"`
	StackedSeries StackedSeries   `json:"stackedSeries"`
	Radial        Radial          `json:"radial"`
	UsageTimeline []UsageTimeline `json:"usageTimeline"`
}

func generateRandomFloat(min, max float64) float64 {
	diff := max - min
	randomValue := rand.Float64()*diff + min
	var roundedValue float64
	// Use Sscanf to ensure precise rounding/formatting similar to your PM25 example
	fmt.Sscanf(fmt.Sprintf("%.2f", randomValue), "%f", &roundedValue)
	return roundedValue
}

// GenerateMockWaterData สร้างข้อมูล Mock ของ Water Meter ทั้งหมด
func (s *GetWaterSensorImpl) GetWaterData() WaterMeterData {
	// 1. สุ่มค่า Domestic (น้ำใช้ทั่วไป)
	domesticPH := generateRandomFloat(7.0, 7.5)    // pH 7.0 - 7.5 (เป็นกลางถึงด่างเล็กน้อย)
	flowRateLpm := generateRandomFloat(30.0, 45.0) // 30 - 45 ลิตร/นาที
	// ปริมาณการใช้น้ำใช้ในวันนี้ (Domestic Today)
	domesticToday := int(generateRandomFloat(800, 1100))

	// 2. สุ่มค่า Drinking (น้ำดื่ม)
	drinkingPH := generateRandomFloat(6.5, 7.0) // pH 6.5 - 7.0 (กรดอ่อนๆ ถึงกลาง)
	// TDS 80 - 150 ppm
	tdsPpm := int(generateRandomFloat(80, 150))
	// ปริมาณการใช้น้ำดื่มในวันนี้ (Drinking Today)
	drinkingToday := int(generateRandomFloat(350, 500))

	// 3. กำหนดค่าคงที่สำหรับข้อมูลอนุกรมรายเดือนและ Timeline
	// (เพื่อให้ข้อมูล Mock ดูสมจริง ไม่เปลี่ยนมั่วทุกครั้งที่เรียก)
	// ข้อมูลสะสมรายเดือนและรายปีจะเพิ่มขึ้นตามค่า Today ที่สุ่มมา

	// Constants for totals (assuming today is Nov 11, so we calculate month/year based on today's value)
	// ปริมาณสะสมของเดือน (Month) คือประมาณ 18-20 เท่าของ Today (11 วัน)
	domesticMonth := domesticToday * 18
	drinkingMonth := drinkingToday * 18
	// ปริมาณสะสมของปี (Year) คือประมาณ 100-110 เท่าของ Today
	domesticYear := domesticToday * 105
	drinkingYear := drinkingToday * 105

	// 4. คำนวณค่ารวม
	totalLitersToday := domesticToday + drinkingToday

	// 5. สร้าง Structs และคืนค่า

	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	return WaterMeterData{
		DeviceKey: "WATER-3078-A01",
		Timestamp: now,
		Domestic: DomesticData{
			PH:                domesticPH,
			FlowRateLpm:       flowRateLpm,
			ConsumptionLiters: domesticToday,
		},
		Drinking: DrinkingData{
			PH:                drinkingPH,
			TdsPpm:            tdsPpm,
			ConsumptionLiters: drinkingToday,
		},
		Totals: Totals{
			Domestic: TotalUsage{
				Today: domesticToday,
				Month: domesticMonth,
				Year:  domesticYear,
			},
			Drinking: TotalUsage{
				Today: drinkingToday,
				Month: drinkingMonth,
				Year:  drinkingYear,
			},
		},
		// ใช้ข้อมูลคงที่สำหรับกราฟรายเดือน/Timeline เพื่อความเรียบง่ายในการ Mock
		StackedSeries: StackedSeries{
			Categories: []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			Series: []SeriesData{
				{Name: "Domestic", Data: []int{820, 910, 960, 880, 1020, 990, 1040, 980, 950, 970, domesticToday, 0}},
				{Name: "Drinking", Data: []int{350, 370, 420, 380, 430, 410, 460, 420, 415, 430, drinkingToday, 0}},
				{Name: "Reclaimed", Data: []int{70, 60, 55, 62, 58, 60, 65, 63, 61, 59, 0, 0}},
			},
		},
		Radial: Radial{
			TotalLiters: totalLitersToday,
			Labels:      []string{"Domestic", "Drinking", "Reclaimed"},
			Values:      []int{domesticToday, drinkingToday, 0},
		},
		UsageTimeline: []UsageTimeline{
			{Name: "Domestic", Data: []int{60, 90, 120, 140, 250, 300, 260, 340, 360, 320, 380, 460}},
			{Name: "Drinking", Data: []int{360, 380, 420, 430, 450, 470, 440, 500, 520, 510, 530, 560}},
			{Name: "Reclaimed", Data: []int{540, 560, 590, 600, 650, 700, 660, 740, 780, 760, 800, 840}},
		},
	}
}
