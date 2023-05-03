package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(image multipart.File, handler multipart.FileHeader) {
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	os.Chdir("images")
	tempFile, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array

	fileBytes, err := io.ReadAll(image)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// return that we have successfully uploaded our file!
	fmt.Println("Successfully Uploaded File")
}
