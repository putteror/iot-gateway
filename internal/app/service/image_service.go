package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// ImageService is responsible for handling image-related utilities,
// such as converting files to base64.
type ImageService struct {
	// Path to the directory where images are stored (e.g., "./uploads")
	UploadsDirPath string
}

// NewImageService creates a new instance of ImageService.
func NewImageService(uploadsDirPath string) *ImageService {
	return &ImageService{
		UploadsDirPath: uploadsDirPath,
	}
}

// ConvertImageToBase64 reads an image file from the specified uploads directory
// and converts its content into a Base64 encoded string.
func (s *ImageService) ConvertImageToBase64(filename string) (string, error) {
	// 1. Construct the full file path
	filePath := filepath.Join(s.UploadsDirPath, filename)

	// 2. Read the entire file content using ioutil.ReadFile
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file %s: %w", filePath, err)
	}

	// 3. Encode the image bytes to a standard Base64 string
	base64String := base64.StdEncoding.EncodeToString(imageData)

	return base64String, nil
}
