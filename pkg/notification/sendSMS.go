package notification

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/ShopSphere/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
)

type NotifyClient interface {
	SendSMS(ctx context.Context, phone, message string) error
}

type notifyClient struct {
}

func (n *notifyClient) SendSMS(ctx context.Context, phone, message string) error {
	accountSid := config.AppConfig.Necessities.AccountSMSSid
	authToken := config.AppConfig.Necessities.AuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(config.AppConfig.Necessities.SetFROM)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("failed to send SMS to %s: %v", phone, err)
		return fmt.Errorf("failed to send SMS: %w", err)
	}

	log.Printf("SMS sent to %s: SID=%s, Status=%s", phone, *resp.Sid, *resp.Status)
	return nil
}

func NewNotifyClient() NotifyClient {
	return &notifyClient{}
}
