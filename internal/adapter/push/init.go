// internal/adapter/push/push.go
package push

import "github.com/putteror/iot-gateway/internal/app/schema"

// PushDataAdapter คือ Interface กลางสำหรับ Adapter ทั้งหมด
type PushDataAdapter interface {
	PushFaceRecognitionEventData(payload *schema.FaceRecognitionEventSchema) error
}
