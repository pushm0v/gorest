package model

type EmailMessage struct {
	Subject         string `json:"subject"`
	Destination     string `json:"destination"`
	DestinationName string `json:"destination_name"`
	Body            string `json:"body"`
}
