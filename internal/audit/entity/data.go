package entity

import "time"

type Analytics struct {
	UploadedAt time.Time
	UserID     string
	Data       DataUser
}

type DataUser struct {
	Headers map[string]string `json:"headers"`
	Body    BodyData          `json:"body"`
}

type BodyData struct {
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
