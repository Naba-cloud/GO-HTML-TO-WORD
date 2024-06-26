package main

import (
	"fmt"
	"os"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

func main() {
	config, _ := models.NewConfiguration("config.json")
	wordsApi, ctx, _ := api.CreateWordsApi(config)
	doc, _ := os.Open("Input.html")
	defer doc.Close()
	options := map[string]interface{}{
		"permissions": "true",
		"password":    "123",
		"saveFormat":  "true",
	}
	request := &models.ConvertDocumentRequest{
		Document:  doc,
		Format:    ToStringPointer("docx"),
		Optionals: options,
	}
	convert, err := wordsApi.ConvertDocument(ctx, request)

	if err != nil {
		fmt.Println("Error converting document:", err)
		return
	}
	defer convert.Body.Close()
	outputFile, err := os.Create("Output3.docx")

	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// _, err = io.Copy(outputFile, convert.Body)
	_, err = outputFile.ReadFrom(convert.Body)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Document converted and saved successfully")

}

func ToStringPointer(s string) *string {
	return &s
}
