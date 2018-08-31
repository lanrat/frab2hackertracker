/*
 * Semi autogenerated with: https://mholt.github.io/json-to-go/
 */

package main

import (
	"bytes"
	"encoding/json"
)

type HackerTrackerArticle struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	Conference string `json:"conference"`
	Text       string `json:"text"`
}

type HackerTrackerEventType struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	Conference string `json:"conference"`
	Color      string `json:"color"`
}

type HackerTrackerEvent struct {
	StartDate   string `json:"start_date"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Location    int    `json:"location"`
	Link        string `json:"link"`
	Speakers    []int  `json:"speakers"`
	EndDate     string `json:"end_date"`
	Conference  string `json:"conference"`
	EventType   int    `json:"event_type"`
	Includes    string `json:"includes"`
	Title       string `json:"title"`
	UpdatedAt   string `json:"updated_at"`
}

type HackerTrackerFaq struct {
	Answer     string `json:"answer"`
	ID         int    `json:"id"`
	Question   string `json:"question"`
	UpdatedAt  string `json:"updated_at"`
	Conference string `json:"conference"`
}

type HackerTrackerLocation struct {
	Name       string `json:"name"`
	ID         int    `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	Conference string `json:"conference"`
}

type HackerTrackerSpeaker struct {
	Name        string `json:"name"`
	UpdatedAt   string `json:"updated_at"`
	Description string `json:"description"`
	Title       string `json:"title"`
	ID          int    `json:"id"`
	Twitter     string `json:"twitter"`
	Conference  string `json:"conference"`
	Link        string `json:"link"`
}

type HackerTrackerVendor struct {
	ID          int    `json:"id"`
	Partner     bool   `json:"partner"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Conference  string `json:"conference"`
	UpdatedAt   string `json:"updated_at"`
}

func HackerTrackerMarshal(name string, data interface{}) ([]byte, error) {
	m := make(map[string]interface{})
	m["updated_at"] = updatedAt
	m[name] = data
	//return json.MarshalIndent(m, "", "  ")

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(m)
	return buffer.Bytes(), err
}
