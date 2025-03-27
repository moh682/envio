package filemanager

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"strings"
)

type file struct{}

type FileManager interface {
	SaveFileFromString(dataURI string, path string) error
}

func NewFileManager() FileManager {
	return &file{}
}

// Save a file from stringuri to a file
func (f *file) SaveFileFromString(dataURI string, path string) error {

	// 1. Verify it starts with the expected prefix (mimeType might vary).
	prefix := fmt.Sprintf("data:application/pdf;filename=%s;base64,", path)
	if !strings.HasPrefix(dataURI, prefix) {
		return fmt.Errorf("Data URI does not appear to be a PDF in base64 format.")
	}

	postfix := strings.TrimPrefix(".pdf", path)
	if !strings.HasSuffix(path, postfix) {
		return fmt.Errorf("Data URI does not appear to be a PDF in base64 format. was given %s", path)
	}
	log.Println("Saving file to:", path)

	// 2. Extract the actual base64 content.
	base64Data := strings.TrimPrefix(dataURI, prefix)

	// 3. Decode the base64 string.
	pdfData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("Error decoding base64:", err)
	}

	// 4. Write the decoded data to a PDF file.
	err = os.WriteFile(path, pdfData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file:", err)
	}

	fmt.Printf("PDF file created: %s\n", path)
	return nil
}
