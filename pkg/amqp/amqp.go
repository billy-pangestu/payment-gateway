package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// IQueue ...
type IQueue interface {
	PushQueue(data map[string]interface{}, types string) error
	PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error)
	PushDLQueueReconnect(url string, data map[string]interface{}, types string) (*amqp.Connection, *amqp.Channel, error)
}

const (
	// ActivationMailExchange ...
	ActivationMailExchange = "activation_mail.exchange"
	// ActivationMail ...
	ActivationMail = "activation_mail.incoming.queue"
	// ActivationMailDeadLetter ...
	ActivationMailDeadLetter = "activation_mail.deadletter.queue"

	// RegistrationOtpExchange ...
	RegistrationOtpExchange = "registration_otp.exchange"
	// RegistrationOtp ...
	RegistrationOtp = "registration_otp.incoming.queue"
	// RegistrationOtpDeadLetter ...
	RegistrationOtpDeadLetter = "registration_otp.deadletter.queue"

	// GetAzureFaceIDExchange ...
	GetAzureFaceIDExchange = "get_azure_face_id.exchange"
	// GetAzureFaceID ...
	GetAzureFaceID = "get_azure_face_id.incoming.queue"
	// GetAzureFaceIDDeadLetter ...
	GetAzureFaceIDDeadLetter = "get_azure_face_id.deadletter.queue"

	// DataSubmitedMailExchange ...
	DataSubmitedMailExchange = "data_submited_mail.exchange"
	// DataSubmitedMail ...
	DataSubmitedMail = "data_submited_mail.incoming.queue"
	// DataSubmitedMailDeadLetter ...
	DataSubmitedMailDeadLetter = "data_submited_mail.deadletter.queue"

	// GetIdentityOcrExchange ...
	GetIdentityOcrExchange = "get_identity_ocr.exchange"
	// GetIdentityOcr ...
	GetIdentityOcr = "get_identity_ocr.incoming.queue"
	// GetIdentityOcrDeadLetter ...
	GetIdentityOcrDeadLetter = "get_identity_ocr.deadletter.queue"

	// MarginConfirmationExchange ...
	MarginConfirmationExchange = "margin_confirmation.exchange"
	// MarginConfirmation ...
	MarginConfirmation = "margin_confirmation.incoming.queue"
	// MarginConfirmationDeadLetter ...
	MarginConfirmationDeadLetter = "margin_confirmation.deadletter.queue"

	// RegistrationReminderExchange ...
	RegistrationReminderExchange = "registration_reminder.exchange"
	// RegistrationReminder ...
	RegistrationReminder = "registration_reminder.incoming.queue"
	// RegistrationReminderDeadLetter ...
	RegistrationReminderDeadLetter = "registration_reminder.deadletter.queue"

	// SignDocumentReminderExchange ...
	SignDocumentReminderExchange = "sign_document_reminder.exchange"
	// SignDocumentReminder ...
	SignDocumentReminder = "sign_document_reminder.incoming.queue"
	// SignDocumentReminderDeadLetter ...
	SignDocumentReminderDeadLetter = "sign_document_reminder.deadletter.queue"

	// StartTransactionReminderExchange ...
	StartTransactionReminderExchange = "start_transaction_reminder.exchange"
	// StartTransactionReminder ...
	StartTransactionReminder = "start_transaction_reminder.incoming.queue"
	// StartTransactionReminderDeadLetter ...
	StartTransactionReminderDeadLetter = "start_transaction_reminder.deadletter.queue"

	// ProfitTakingReminderExchange ...
	ProfitTakingReminderExchange = "profit_taking_reminder.exchange"
	// ProfitTakingReminder ...
	ProfitTakingReminder = "profit_taking_reminder.incoming.queue"
	// ProfitTakingReminderDeadLetter ...
	ProfitTakingReminderDeadLetter = "profit_taking_reminder.deadletter.queue"

	// MarginUserSpouseOtpExchange ...
	MarginUserSpouseOtpExchange = "margin_user_spouse_otp.exchange"
	// MarginUserSpouseOtp ...
	MarginUserSpouseOtp = "margin_user_spouse_otp.incoming.queue"
	// MarginUserSpouseDeadLetter ...
	MarginUserSpouseDeadLetter = "margin_user_spouse_otp.deadletter.queue"

	// AdminReviewExchange ...
	AdminReviewExchange = "admin_review.exchange"
	// AdminReview ...
	AdminReview = "admin_review.incoming.queue"
	// AdminReviewDeadLetter ...
	AdminReviewDeadLetter = "admin_review.deadletter.queue"
	// AdminReviewRepairDataUser ...
	AdminReviewRepairDataUser = "repair_data_user"
	// AdminReviewRepairDataAdmin ...
	AdminReviewRepairDataAdmin = "repair_data_admin"
	// AdminReviewReported ...
	AdminReviewReported = "admin_reported"
	// AdminReviewSignDocument ...
	AdminReviewSignDocument = "user_sign_document"

	// MarginWhitelistExchange ...
	MarginWhitelistExchange = "margin"
	// MarginWhitelist ...
	MarginWhitelist = "eligible"
	// MarginWhitelistDeadLetter ...
	MarginWhitelistDeadLetter = "client"

	// MarginApprovalExchange ...
	MarginApprovalExchange = "margin"
	// MarginApproval ...
	MarginApproval = "newportfolio"
	// MarginApprovalDeadLetter ...
	MarginApprovalDeadLetter = "portfolio"

	// MarginOfferExchange ...
	MarginOfferExchange = "margin_offer.exchange"
	// MarginOffer ...
	MarginOffer = "margin_offer.incoming.queue"
	// MarginOfferDeadLetter ...
	MarginOfferDeadLetter = "margin_offer.deadletter.queue"

	// PrivyRegistrationExchange ...
	PrivyRegistrationExchange = "privy_registration.exchange"
	// PrivyRegistration ...
	PrivyRegistration = "privy_registration.incoming.queue"
	// PrivyRegistrationDeadLetter ...
	PrivyRegistrationDeadLetter = "privy_registration.deadletter.queue"

	// PrivyCallbackExchange ...
	PrivyCallbackExchange = "privy_callback.exchange"
	// PrivyCallback ...
	PrivyCallback = "privy_callback.incoming.queue"
	// PrivyCallbackDeadLetter ...
	PrivyCallbackDeadLetter = "privy_callback.deadletter.queue"

	// UserAutoVerificationExchange ...
	UserAutoVerificationExchange = "user_auto_verification.exchange"
	// UserAutoVerification ...
	UserAutoVerification = "user_auto_verification.incoming.queue"
	// UserAutoVerificationDeadLetter ...
	UserAutoVerificationDeadLetter = "user_auto_verification.deadletter.queue"

	// APIUserSubmitExchange ...
	APIUserSubmitExchange = "api_user_submit.exchange"
	// APIUserSubmit ...
	APIUserSubmit = "api_user_submit.incoming.queue"
	// APIUserSubmitDeadLetter ...
	APIUserSubmitDeadLetter = "api_user_submit.deadletter.queue"

	// APIUserCallbackExchange ...
	APIUserCallbackExchange = "api_user_callback.exchange"
	// APIUserCallback ...
	APIUserCallback = "api_user_callback.incoming.queue"
	// APIUserCallbackDeadLetter ...
	APIUserCallbackDeadLetter = "api_user_callback.deadletter.queue"

	// ApproveCMExchange ...
	ApproveCMExchange = "mscm"
	// ApproveCM ...
	ApproveCM = "newclient"
	// ApproveCMExchangeDeadLetter ...
	ApproveCMExchangeDeadLetter = "client"

	// SubmitSecuritiesExchange ...
	SubmitSecuritiesExchange = "submit_securities.exchange"
	// SubmitSecurities ...
	SubmitSecurities = "submit_securities.incoming.queue"
	// SubmitSecuritiesDeadLetter ...
	SubmitSecuritiesDeadLetter = "submit_securities.deadletter.queue"

	// MailActiveExchange ...
	MailActiveExchange = "email_active.exchange"
	// MailActive ...
	MailActive = "email_active.incoming.queue"
	// MailActiveDeadLetter ...
	MailActiveDeadLetter = "email_active.deadletter.queue"

	// BatchFileSubmitExchange ...
	BatchFileSubmitExchange = "batch_file_submit.exchange"
	// BatchFileSubmit ...
	BatchFileSubmit = "batch_file_submit.incoming.queue"
	// BatchFileSubmitDeadLetter ...
	BatchFileSubmitDeadLetter = "batch_file_submit.deadletter.queue"

	// APIUserEmailVerificationExchange ...
	APIUserEmailVerificationExchange = "api_user_email_verification.exchange"
	// APIUserEmailVerification ...
	APIUserEmailVerification = "api_user_email_verification.incoming.queue"
	// APIUserEmailVerificationDeadLetter ...
	APIUserEmailVerificationDeadLetter = "api_user_email_verification.deadletter.queue"

	// PrivyDocumentReuploadExchange ...
	PrivyDocumentReuploadExchange = "privy_document_reupload.exchange"
	// PrivyDocumentReupload ...
	PrivyDocumentReupload = "privy_document_reupload.incoming.queue"
	// PrivyDocumentReuploadDeadLetter ...
	PrivyDocumentReuploadDeadLetter = "privy_document_reupload.deadletter.queue"
)

// queue ...
type queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewQueue ...
func NewQueue(conn *amqp.Connection, channel *amqp.Channel) IQueue {
	return &queue{
		Connection: conn,
		Channel:    channel,
	}
}

// PushQueue ...
func (m queue) PushQueue(data map[string]interface{}, types string) error {
	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return err
}

// PushQueueReconnect ...
func (m queue) PushQueueReconnect(url string, data map[string]interface{}, types, deadLetterKey string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": deadLetterKey,
	}

	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, args)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return m.Connection, m.Channel, err
}

// PushDLQueueReconnect ...
func (m queue) PushDLQueueReconnect(url string, data map[string]interface{}, types string) (*amqp.Connection, *amqp.Channel, error) {
	if m.Connection != nil {
		if m.Connection.IsClosed() {
			c := Connection{
				URL: url,
			}
			newConn, newChannel, err := c.Connect()
			if err != nil {
				return nil, nil, err
			}
			m.Connection = newConn
			m.Channel = newChannel
		}
	} else {
		c := Connection{
			URL: url,
		}
		newConn, newChannel, err := c.Connect()
		if err != nil {
			return nil, nil, err
		}
		m.Connection = newConn
		m.Channel = newChannel
	}

	queue, err := m.Channel.QueueDeclare(types, true, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, nil, nil
	}

	err = m.Channel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	return m.Connection, m.Channel, err
}
