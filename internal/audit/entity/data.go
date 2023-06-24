package entity

import "time"

type Analytics struct {
	UploadedAt time.Time
	UserID     string
	Data       DataBody
}

type DataBody struct {
	Module string     `json:"module"`
	Type   string     `json:"type"`
	Event  string     `json:"event"`
	Name   string     `json:"name"`
	Data   ActionData `json:"data"`
}

type ActionData struct {
	Action string `json:"action"`
}

/*
"module" : "settings",
"type" : "alert",
"event" : "click",
"name" : "подтверждение выхода",
"data" : {"action" : "cancel"}
*/
