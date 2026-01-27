package push

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/putteror/iot-gateway/internal/app/schema"
	"github.com/putteror/iot-gateway/internal/config"
)

type ZytaPushDataAdapter interface {
	PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error
	PushEmergencyAlarmEventData(payload *schema.EmergencyAlarmEventSchema) error
	PushPM25Value(pm25Value float64, pm10Value float64, humidityValue float64, temperatureValue float64) error
	PushWaterSensorData() error
}

type ZytaPushDataServiceImpl struct {
}

func NewPushDataServiceImpl() ZytaPushDataAdapter {
	return &ZytaPushDataServiceImpl{}
}

func (s *ZytaPushDataServiceImpl) PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error {

	type ZytaFaceRecognitionEventSchema struct {
		Type        string `json:"type"`
		Severity    string `json:"severity"`
		TitleKey    string `json:"titleKey"`
		ImageBase64 string `json:"img"`
		OccurredAt  string `json:"occurredAt"`
		Meta        struct {
			Kind   string `json:"kind"`
			Person struct {
				FullName string `json:"fullName"`
				Gender   string `json:"gender"`
				Age      int    `json:"age"`
			} `json:"person"`
			RawID string `json:"rawId"`
		} `json:"meta"`
	}

	var sendPayload = new(ZytaFaceRecognitionEventSchema)
	sendPayload.Type = "info"
	sendPayload.Severity = "low"
	sendPayload.TitleKey = "notis.faceDetected"
	sendPayload.ImageBase64 = payload.ImageBase64
	sendPayload.OccurredAt = payload.StampDateTime.Format("2006-01-02T15:04:05Z")
	sendPayload.Meta.Kind = "face"
	// sendPayload.Meta.Person.FullName = payload.PersonInformation.FirstName + " " + payload.PersonInformation.LastName
	sendPayload.Meta.Person.FullName = "ตรวจพบใบหน้า"
	sendPayload.Meta.Person.Age = payload.PersonInformation.Age
	// sendPayload.Meta.Person.Gender = payload.PersonInformation.Gender
	sendPayload.Meta.Person.Gender = ""
	sendPayload.Meta.RawID = payload.PersonInformation.ID

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/webhooks/notis"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Push data to URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", payload.DeviceInformation.ID)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		log.Println("Webhook sent successfully.")
	}

	return nil
}

func (s *ZytaPushDataServiceImpl) PushEmergencyAlarmEventData(payload *schema.EmergencyAlarmEventSchema) error {

	type ZytaEmergencyAlarmEventSchema struct {
		Type       string `json:"type"`
		Severity   string `json:"severity"`
		TitleKey   string `json:"titleKey"`
		Title      string `json:"title"`
		SiteID     string `json:"siteId"`
		DeviceID   string `json:"deviceId"`
		OccurredAt string `json:"occurredAt"`
		Meta       struct {
		} `json:"meta"`
	}

	var sendPayload = new(ZytaEmergencyAlarmEventSchema)
	sendPayload.Type = "alert"
	sendPayload.TitleKey = "zytaNotis.sos"
	sendPayload.Title = "gateway send"
	sendPayload.SiteID = payload.SiteID
	sendPayload.DeviceID = payload.DeviceInformation.ID
	sendPayload.OccurredAt = payload.StampDateTime.Format("2006-01-02T15:04:05Z")

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/webhooks/notis"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Push data to URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", payload.DeviceInformation.ID)

	//// tentative //////////////
	// print request body, header
	//// tentative //////////////
	log.Println("Request Header:", req.Header)
	log.Println("Request Body:", string(convertedJsonPayload))

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body: %v\n", err)
		} else {
			// Cast ke string agar bisa dibaca di log
			//// tentative //////////////
			fmt.Printf("Response Body: %s\n", string(bodyBytes))
		}
		log.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		log.Println("Webhook alarm sent successfully.")
	}

	return nil
}

func (s *ZytaPushDataServiceImpl) PushPM25Value(pm25Value float64, pm10Value float64, humidityValue float64, temperatureValue float64) error {

	type ZytaAirSensorPayload struct {
		DeviceKey string    `json:"deviceKey"`
		Timestamp time.Time `json:"timestamp"`
		Source    string    `json:"source"`
		Metrics   struct {
			PM25        float64 `json:"pm25"`
			PM10        float64 `json:"pm10"`
			Humidity    float64 `json:"humidity"`
			Temperature float64 `json:"temperature"`
		} `json:"metrics"`
	}

	deviceId := "AIR-3078-S01"

	var sendPayload = new(ZytaAirSensorPayload)
	sendPayload.DeviceKey = deviceId
	sendPayload.Timestamp = time.Now()
	sendPayload.Source = "iaq-node-1"
	sendPayload.Metrics.PM25 = pm25Value
	sendPayload.Metrics.PM10 = pm10Value
	sendPayload.Metrics.Humidity = humidityValue
	sendPayload.Metrics.Temperature = temperatureValue

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/webhooks/air-sensors"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Push data to URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(sendPayload)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", deviceId)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		log.Println("Webhook sent successfully.")
	}

	return nil
}

func generateRandomFloat(min, max float64) float64 {
	diff := max - min
	randomValue := rand.Float64()*diff + min
	var roundedValue float64
	// Use Sscanf to ensure precise rounding/formatting similar to your PM25 example
	fmt.Sscanf(fmt.Sprintf("%.2f", randomValue), "%f", &roundedValue)
	return roundedValue
}
func (s *ZytaPushDataServiceImpl) PushWaterSensorData() error {

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

	// 1. สุ่มค่า Domestic (น้ำใช้ทั่วไป)
	domesticPH := generateRandomFloat(7.0, 7.5)    // pH 7.0 - 7.5 (เป็นกลางถึงด่างเล็กน้อย)
	flowRateLpm := generateRandomFloat(30.0, 45.0) // 30 - 45 ลิตร/นาที
	// ปริมาณการใช้น้ำใช้ในวันนี้ (Domestic Today)
	domesticToday := int(generateRandomFloat(800, 1100))

	// 2. สุ่มค่า Drinking (น้ำดื่ม)
	drinkingPH := generateRandomFloat(6.5, 7.0) // pH 6.5 - 7.0 (กรดอ่อนๆ ถึงกลาง)
	// TDS 80 - 150 ppm
	// timestampSecInt := (time.Now().Unix())
	// setCalValue := ((timestampSecInt - 1762836000) / 100)
	// domesticToday := int(setCalValue*100) / 7
	// tdsPpm := int(setCalValue*100) / 60
	// drinkingToday := int(setCalValue*100) / 15
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

	deviceId := "WATER-3078-A01"

	payloaddData := WaterMeterData{
		DeviceKey: deviceId,
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

	////////////////////////////
	/// Webhook 1 //////////////
	////////////////////////////

	webhookPath := "/api/webhooks/water-meters"
	webhookUrl := config.DESTINATION_HOST_ADDRESS + webhookPath
	log.Println("Push data to URL:", webhookUrl)

	convertedJsonPayload, err := json.Marshal(payloaddData)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return err
	}
	//fmt.Println(string(convertedJsonPayload))

	client := &http.Client{}

	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(convertedJsonPayload))
	if err != nil {
		log.Printf("Error creating HTTP request: %v\n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Key", deviceId)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending webhook POST request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Webhook responded with status code: %d\n", resp.StatusCode)
	} else {
		log.Println("Webhook sent successfully.")
	}

	return nil

}
