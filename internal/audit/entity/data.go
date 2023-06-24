package entity

import "time"

type Analytics struct {
	UploadedAt time.Time
	UserID     string
	Data       DataBody
}

type DataBody struct {
	header string
	body   string
}
