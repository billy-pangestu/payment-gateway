package mandrill

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"payment-gateway-backend/pkg/file"

	"github.com/keighl/mandrill"
)

// Credential ...
type Credential struct {
	Key      string
	FromMail string
	FromName string
}

// Attachment ...
type Attachment struct {
	Name string
	URL  string
}

type template struct {
	Attribute []templateDetail `json:"attribute"`
}

type templateDetail struct {
	Code string `json:"code"`
}

var (
	getTemplateURL = "https://mandrillapp.com/api/1.0/templates/list.json"

	/* Mandrill Template list */

	// EmailVerification ...
	EmailVerification = "email_verification"
	// AdminVerification ...
	AdminVerification = "admin_verification"
	// SendEmailNotif ...
	SendEmailNotif = "email_sign_notif"
	// SendEmailConfirm ...
	SendEmailConfirm = "email_sign_confirm"

	// RegisterReminderOne ...
	RegisterReminderOne = "register_reminder_one"
	// RegisterReminderTwo ...
	RegisterReminderTwo = "register_reminder_two"
	// RegisterReminderThree ...
	RegisterReminderThree = "register_reminder_three"
	// RegisterReminderFour ...
	RegisterReminderFour = "register_reminder_four"
	// SignatureReminderOne ...
	SignatureReminderOne = "signature_reminder_one"
	// SignatureReminderTwo ...
	SignatureReminderTwo = "signature_reminder_two"
	// TransactionReminderOne ...
	TransactionReminderOne = "transaction_reminder_one"
	// TransactionReminderTwo ...
	TransactionReminderTwo = "transaction_reminder_two"
	// ProfitTakingReminderOne ...
	ProfitTakingReminderOne = "profit_taking_reminder_one"

	// RepairDataUser ...
	RepairDataUser = "user_repair_data"
	// RepairDataAdmin ...
	RepairDataAdmin = "admin_repair_data"
	// EmailSignNotif sent when document already generated to privy and need user sign
	EmailSignNotif = "email_sign_notif"
	// AdminReported sent when reported by admin
	AdminReported = "admin_reported"

	// MarginSingleSubmission ...
	MarginSingleSubmission = "margin_single_submission"
	// MarginNotifSpouse ...
	MarginNotifSpouse = "margin_notif_spouse"
	// MarginSpouseVerification ...
	MarginSpouseVerification = "margin_spouse_verification"
	// MarginSpouseContinueLink ...
	MarginSpouseContinueLink = "margin_spouse_continue_link"
	// MarginSpouseAccepted ...
	MarginSpouseAccepted = "margin_spouse_accepted"
	// MarginSpouseRejected ...
	MarginSpouseRejected = "margin_spouse_rejected"
	// MarginSpouseFinal ...
	MarginSpouseFinal = "margin_spouse_final"
	// MarginWoSpouseApproval ...
	MarginWoSpouseApproval = "margin_wo_spouse_approval"
	// MarginApproved ...
	MarginApproved = "margin_approved"
	// MarginSuccess ...
	MarginSuccess = "margin_success"

	// MarginOfferOne ...
	MarginOfferOne = "margin_offer_one"
	// MarginOfferTwo ...
	MarginOfferTwo = "margin_offer_two"
	// MarginOfferThree ...
	MarginOfferThree = "margin_offer_three"

	// EmailActive ...
	EmailActive = "email_active"
	// EmailActiveMargin ...
	EmailActiveMargin = "email_active_margin"
	// Default ...
	Default = "default"
)

// Send ...
func (cred *Credential) Send(to, toName, subject, html, text string) error {
	client := mandrill.ClientWithKey(cred.Key)

	message := &mandrill.Message{}
	message.AddRecipient(to, toName, "to")
	message.FromEmail = cred.FromMail
	message.FromName = cred.FromName
	message.Subject = subject
	message.HTML = html
	message.Text = text

	_, err := client.MessagesSend(message)

	return err
}

// SendMulti ...
func (cred *Credential) SendMulti(to []string, subject, html, text string) error {
	// Check to empty
	if len(to) < 1 {
		return errors.New("Empty Recipient")
	}

	client := mandrill.ClientWithKey(cred.Key)

	message := &mandrill.Message{}
	for _, t := range to {
		message.AddRecipient(t, t, "to")
	}
	message.FromEmail = cred.FromMail
	message.FromName = cred.FromName
	message.Subject = subject
	message.HTML = html
	message.Text = text

	_, err := client.MessagesSend(message)

	return err
}

// SendAttachment ...
func (cred *Credential) SendAttachment(to, toName, subject, html, text string, attachment []Attachment) error {
	client := mandrill.ClientWithKey(cred.Key)

	message := &mandrill.Message{}
	message.AddRecipient(to, toName, "to")
	message.FromEmail = cred.FromMail
	message.FromName = cred.FromName
	message.Subject = subject
	message.HTML = html
	message.Text = text

	// Add attachment
	var mandrillAttachment []*mandrill.Attachment
	for _, a := range attachment {
		// Get file type and base64 string
		content, types := file.PathToBase64(a.URL)

		// Append to mandrill attachment struct
		mandrillAttachment = append(mandrillAttachment, &mandrill.Attachment{
			Type:    types,
			Name:    a.Name,
			Content: content,
		})
	}
	message.Attachments = mandrillAttachment

	_, err := client.MessagesSend(message)

	return err
}

// GetTemplate ...
func (cred *Credential) GetTemplate(name string) (string, error) {
	jsonStr := []byte(`{"key": "` + cred.Key + `","label": "` + name + `"}`)
	req, err := http.NewRequest("POST", getTemplateURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

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
	data := template{}
	json.Unmarshal(body, &data.Attribute)

	if len(data.Attribute) <= 0 {
		return "", errors.New("empty")
	}

	return data.Attribute[0].Code, nil
}
