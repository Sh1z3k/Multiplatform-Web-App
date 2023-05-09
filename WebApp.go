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
	filePath := returnFilePath()
	archivePath := returnArchivePath()
	//NIBY WIEM, ALE KURWA NIE DO KOŃCA XDDDD
	http.Handle("/", http.FileServer(http.Dir("web")))
	// OPISAĆ CO TU SIĘ DZIEJE JAK SIĘ DOWIEM XDDDDDDDDDDD
	http.HandleFunc("/get-time", func(rw http.ResponseWriter, r *http.Request) {
		CurrentTime := Response{
			CurrentTime: time.Now().Format(time.RFC3339),
		}
		//Dodać comment jak ogarnę marshala
		ByteArray, err := json.Marshal(CurrentTime)

		//Printing Error if occurs when getting the CurrentTime via json.Marshal
		if err != nil {
			fmt.Println(err)
		}
		//Checking if logs file exists if not creating the file
		if _, err := os.Stat(filePath); err != nil {
			os.Create(filePath)
			fmt.Println("Logs.txt not found => creating a new logs file!")
		}

		rw.Write(ByteArray)
		textToFile := CurrentTime
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error while openning the logs file:", err)
			return
		}
		defer file.Close()
		// _, - 2 variables with the same name can
		_, err = fmt.Fprintln(file, strings.Join([]string{"Button clicked at:", textToFile.CurrentTime}, " "))
		if err != nil {
			fmt.Println("Error while writing to the logs file:", err)
			return
		}

		fmt.Println("Text added to the logs file")
	})

	http.HandleFunc("/clear-logs", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(filePath); err == nil {
			os.Remove("logs.txt")
			fmt.Printf("Logs cleared succesfully\n")
		} else {
			fmt.Printf("File does not exist\n")
		}
	})

	// TWORZY FAILE ZAMIAST DIRECTORY - POPRAWIĆ
	http.HandleFunc("/archive-logs", func(rw http.ResponseWriter, r *http.Request) {

		if _, err := os.Stat(archivePath); err == nil {
			os.MkdirAll(archivePath, os.ModePerm)
			fmt.Printf("Archive has been initialized\n Current logs file has been archivized\n")

		} else {
			fmt.Printf("Archive already exists\n Current logs file has been archivized\n")
		}
		os.Link(filePath, returnArchivePath())
	})

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}

}

// Checking on what OS the app is running by comparing a files path separators
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

// Returns the correct for the currently running OS path to the logs file where logs are stored/saved
func returnFilePath() string {
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

func returnArchivePath() string {
	archivePath := "NON DECLARED VALUE"
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	if checkOS() == "Windows" {
		archivePath := strings.Join([]string{path, "\\archive\\"}, "")
		return archivePath
	} else if checkOS() == "Linux" {
		archivePath := strings.Join([]string{path, "/archive/"}, "")
		return archivePath
	} else if checkOS() == "MacOS" {
		archivePath := strings.Join([]string{path, ":archive:"}, "")
		return archivePath
	} else {
		return archivePath
	}

}
