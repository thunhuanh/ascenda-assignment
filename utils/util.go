package utils

import (
	"log"
	"strings"
)

var ErrLog *log.Logger
var InfoLog *log.Logger

func init() {
	ErrLog = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NonEmpty(a, b float64) float64 {
	if a != 0 {
		return a
	}
	return b
}

func RemoveDupStr(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[strings.ToLower(entry)]; !value {
			keys[strings.ToLower(entry)] = true
			list = append(list, entry)
		}
	}
	return list
}
