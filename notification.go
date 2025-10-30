package apple

import (
	"encoding/json"
	"github.com/xushuhui/apple/internal/storekit"
	"net/http"
)

const (
	kTestNotification = "/v1/notifications/test"
)

// RequestTestNotification https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
func (c *Client) RequestTestNotification() (result *TestNotificationResponse, err error) {
	err = c.request(http.MethodPost, c.BuildAPI(kTestNotification), nil, nil, &result)
	return result, err
}

// DecodeNotification 用于解析通知数据 https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2
//
// 关于接收到苹果服务器推送的通知之后，业务服务器如何响应参照：
// https://developer.apple.com/documentation/appstoreservernotifications/responding_to_app_store_server_notifications
func (c *Client) DecodeNotification(data []byte) (*Notification, error) {
	return DecodeNotification(data)
}

func DecodeNotification(data []byte) (*Notification, error) {
	var aux = struct {
		SignedPayload string `json:"signedPayload"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, err
	}

	var notification = &Notification{}
	if err := storekit.DecodeClaims(aux.SignedPayload, notification); err != nil {
		return nil, err
	}
	return notification, nil
}
