package main

import (
	"fmt"
	"time"
)

type LogEntry struct {
	IP        string
	HTTPCode  int
	Timestamp time.Time
}

func filterErrors(logs []LogEntry) []LogEntry {
	var filtered []LogEntry
	for _, entry := range logs {
		if entry.HTTPCode >= 400 && entry.HTTPCode <= 599 {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}
func main() {
	logs := []LogEntry{
		{IP: "192.168.1.1", HTTPCode: 200, Timestamp: time.Now()},
		{IP: "192.168.1.2", HTTPCode: 404, Timestamp: time.Now()},
		{IP: "192.168.1.3", HTTPCode: 500, Timestamp: time.Now()},
	}
	errors := filterErrors(logs)
	fmt.Printf("Найдено ошибок: %d\n", len(errors))
}
