// internal/adapter/push/push.go
package push

import "github.com/putteror/iot-gateway/internal/app/schema"

// PushDataAdapter คือ Interface กลางสำหรับ Adapter ทั้งหมด
type PushDataAdapter interface {
	PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error
	PushEmergencyAlarmEventData(payload *schema.EmergencyAlarmEventSchema) error
	PushPM25Value(pm25Value float64, pm10Value float64, humidityValue float64, temperatureValue float64) error
	PushWaterSensorData() error
}
