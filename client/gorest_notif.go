package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	ph "github.com/kitabisa/perkakas/v2/httpclient"
	"github.com/pushm0v/gorest/model"
)

type GorestNotif interface {
	SendEmail(m *model.EmailMessage) error
}

type gorestNotif struct {
	httpClient *ph.HttpClient
	notifUrl   string
}

func NewGorestNotif(notifUrl string) GorestNotif {
	conf := new(ph.HttpClientConf)
	conf.BackoffInterval = 2 * time.Millisecond       // 2ms
	conf.MaximumJitterInterval = 5 * time.Millisecond // 5ms
	conf.Timeout = 15000 * time.Millisecond           // 15s
	conf.RetryCount = 3                               // 3 times

	phClient := ph.NewHttpClient(conf)

	return &gorestNotif{
		httpClient: phClient,
		notifUrl:   notifUrl,
	}
}

func (c *gorestNotif) SendEmail(m *model.EmailMessage) error {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(m)
	resp, err := c.httpClient.Client.Post(fmt.Sprintf("%s/api/v1/email", c.notifUrl), payloadBuf, http.Header{})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 201 {
		return fmt.Errorf("Error from notif service : [%d] %v", resp.StatusCode, string(body))
	}

	return nil
}
