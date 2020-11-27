package dingtalk

import "testing"

func TestDingTalkClient_SendMessage(t *testing.T) {
	message := &DingTalkMessage{

	}
	client := &DingTalkClient{
		Message: message,
	}

	client.SendMessage()
}
