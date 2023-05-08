package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	CurrentTime string
}

func main() {
	//filePath := "C:\\Users\\kacpe\\Desktop\\Web App\\logs.txt"
	filePath := setFilePath()

	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/get-time", func(rw http.ResponseWriter, r *http.Request) {
		CurrentTime := Response{
			CurrentTime: time.Now().Format(time.RFC3339),
		}

		textToFile := CurrentTime

		ByteArray, err := json.Marshal(CurrentTime)
		if err != nil {
			fmt.Println(err)
		}
		rw.Write(ByteArray)

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error while openning the logs file:", err)
			return
		}
		defer file.Close()

		_, err = fmt.Fprintln(file, strings.Join([]string{"Button clicked at:", textToFile.CurrentTime}, " "))
		if err != nil {
			fmt.Println("Error while writing to the logs file:", err)
			return
		}

		fmt.Println("Text added to the logs file")
	})

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}

}
func checkOS() string {
	OS := "NON DECLARED VALUE"
	if os.PathSeparator == '/' {
		OS = "Linux"
		return OS
	} else if os.PathSeparator == '\\' {
		OS = "Windows"
		return OS
	} else if os.PathSeparator == ':' {
		OS = "MacOS"
		return OS
	}
	return OS
}

func setFilePath() string {
	filePath := "NON DECLARED VALUE"
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if checkOS() == "Windows" {
		filePath := strings.Join([]string{path, "\\logs.txt"}, "")
		return filePath
	} else if checkOS() == "Linux" {
		filePath := strings.Join([]string{path, "/logs.txt"}, "")
		return filePath
	} else if checkOS() == "MacOS" {
		filePath := strings.Join([]string{path, ":logs.txt"}, "")
		return filePath
	} else {
		return filePath
	}

}
