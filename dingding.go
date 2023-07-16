package dingding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DingDing struct {
	MsgType     string
	ContentType string
	WebhookUrl  string
	Client      *http.Client
}

func NewDingDing(msgType, contentType, webhookUrl string) *DingDing {
	return &DingDing{
		MsgType:     msgType,
		ContentType: contentType,
		WebhookUrl:  webhookUrl,
		Client:      &http.Client{},
	}
}

func (s *DingDing) SendAlarmMessage(msg string) error {
	content := map[string]string{"content": msg}
	data := map[string]interface{}{
		"msgtype": s.MsgType,
		"text":    content,
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	resp, err := s.Client.Post(s.WebhookUrl, s.ContentType, bytes.NewBuffer(byteData))
	if err != nil {
		return fmt.Errorf("failed to send HTTP POST request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Close ok!")
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Println(string(body))
	return nil
}
