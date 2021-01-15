package service

import (
	"fmt"

	"github.com/pushm0v/gorest/client"
	"github.com/pushm0v/gorest/model"
)

type NotifService interface {
	SendEmailToCustomerCreated(cust *model.Customer) (err error)
}

type notifService struct {
	gorestClient client.GorestNotif
}

func NewNotifService(gorestClient client.GorestNotif) NotifService {
	return &notifService{
		gorestClient: gorestClient,
	}
}

func (n *notifService) SendEmailToCustomerCreated(cust *model.Customer) (err error) {
	var m = new(model.EmailMessage)
	m.Destination = cust.Email
	m.DestinationName = cust.Name
	m.Subject = "Welcome new customer!"
	m.Body = fmt.Sprintf("Hi %s !, welcome to my shop.", cust.Name)

	return n.gorestClient.SendEmail(m)
}
