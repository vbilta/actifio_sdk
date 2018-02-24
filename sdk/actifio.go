package sdk

import (
	"fmt"
	"encoding/json"
	"net/url"
)

type ActConnection struct {
	Name, Password, VendorKey, Appliance string
	SessionId SessionId
}

type SessionId struct {
	SessionId string `json:sessionid`
}

func (conn ActConnection) BaseUrl() string  {

	url := fmt.Sprintf("https://%s/actifio/api", conn.Appliance)
	return url
}

func (conn ActConnection) Login()(SessionId, error)  {
	user := url.QueryEscape(conn.Name)
	password := url.QueryEscape(conn.Password)
	vendorKey := url.QueryEscape(conn.VendorKey)
	baseUrl := conn.BaseUrl()

	url := fmt.Sprintf("%s/login?name=%s&password=%s&vendorkey=%s", baseUrl, user, password, vendorKey)
	response, err := conn.httpPost(url, nil)
	if err != nil {
		return SessionId{}, err
	}
	bytes := []byte(response)

	var sessionId SessionId
	error := json.Unmarshal(bytes, &sessionId)
	conn.SessionId = sessionId
	return sessionId, error
}
