package wavecell

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Connection ...
type Connection struct {
	SubAccountID    string
	SubAccountNotif string
	SubAccountOTP   string
	Secret          string
	SenderName      string
}

var (
	defaultArea = "+62"
	defaultURL  = "https://api.wavecell.com/sms/v1"
	method      = "POST"
)

// Send ...
func (conn *Connection) Send(
	to string,
	message string,
	subAccount string,
) (string, error) {
	to = "+" + strings.Replace(to, "-", "", -1)
	from := url.QueryEscape(conn.SenderName)
	smsURL := defaultURL + "/" + subAccount + "/single"

	// Make payload
	payload := map[string]interface{}{
		"source":      from,
		"destination": to,
		"text":        message,
		"encoding":    "AUTO",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("Error when marshal the payload")
	}
	pBody := []byte(string(b))

	req, err := http.NewRequest(method, smsURL, bytes.NewBuffer(pBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+conn.Secret)
	req.Header.Set("Content-Type", "Application/Json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
