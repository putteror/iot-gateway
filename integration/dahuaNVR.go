package integration

import (
	"fmt"
	"io"
	"net/http"
)

func TestGetData() {
	// 1. กำหนด URL ของ API
	apiURL := "http://203.114.71.22/cgi-bin/configManager.cgi?action=getConfig&name=VideoColor"

	// 2. สร้าง HTTP GET Request
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Error sending GET request: %s\n", err)
		return
	}
	// ตรวจสอบให้แน่ใจว่าได้ปิด Response Body เมื่อสิ้นสุดการทำงาน
	defer response.Body.Close()

	// 3. ตรวจสอบสถานะ (Status Code) ของ Response
	if response.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", response.StatusCode)
		return
	}

	// 4. อ่าน Response Body
	// เนื่องจาก API นี้คืนค่าเป็นรูปแบบเฉพาะ (text-based config) เราจะอ่านเป็น string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// 5. แสดงผลลัพธ์
	fmt.Println("--- API Response Body ---")
	fmt.Println(string(body))
	fmt.Println("-------------------------")

	// หมายเหตุ: ในขั้นตอนถัดไป คุณจะต้องเขียนโค้ดเพื่อแยกวิเคราะห์ (Parse) string 'body'
	// ให้อยู่ในโครงสร้างข้อมูล (struct/map) ที่คุณต้องการ
}
