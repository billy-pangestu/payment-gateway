package helper

var (
	// InternalServer internal server error
	InternalServer = "internal_server"
	// InvalidBody error body on validator.v9 validation
	InvalidBody = "invalid_body"
	// CountryCodeNotExist country code doesn't exist in db
	CountryCodeNotExist = "country_code_not_exist"
	// InvalidPhoneRegionLength phone region length doesn't meet the requirement
	InvalidPhoneRegionLength = "invalid_phone_region_length"
	// InvalidPhoneLength phone length doesn't meet the requirement
	InvalidPhoneLength = "invalid_phone_length"
	// InvalidPhoneFormat invalid format
	InvalidPhoneFormat = "invalid_phone_format"
	// PhoneTooMuch too much duplication on phone number
	PhoneTooMuch = "phone_too_much"
	// EmailExist email already registered
	EmailExist = "email_exist"
	// DBSubmit error on create new record
	DBSubmit = "db_submit"
	// OTP generate otp fail
	OTP = "otp"
	// OtpStillRunning ...
	OtpStillRunning = "otp_still_running"
	// JWT generate jwt fail
	JWT = "jwt"
	// SendSms send sms fail (record to rabbitmq)
	SendSms = "send_sms"
	// RecordNotExist no record found from db
	RecordNotExist = "record_not_exist"
	// SamePhone New phone same with the old one
	SamePhone = "same_phone"
	// DBUpdate error on update record
	DBUpdate = "db_update"
	// ExpOTP expired otp
	ExpOTP = "exp_otp"
	// WrongOTP wrong otp
	WrongOTP = "wrong_otp"
	// MaxSendOtp Max Send otp
	MaxSendOtp = "max_send_otp"
	// ActivationMail send activation mail fail (record to rabbitmq)
	ActivationMail = "activation_mail"
	// SendMail send activation mail fail
	SendMail = "send_mail"
	// InvalidEmailKey wrong key to activate email
	InvalidEmailKey = "invalid_email_key"
	// SameEmail New email same with the old one
	SameEmail = "same_email"
	// UserNotFound User not found when login
	UserNotFound = "user_not_found"
	// FileTooBig uploaded file size is too big
	FileTooBig = "file_too_big"
	// InvalidFileType invalid file type
	InvalidFileType = "invalid_file_type"
	// FileError fail to read file
	FileError = "file_error"
	// UploadFileError fail to upload file
	UploadFileError = "upload_file_error"
	// InvalidIdentityImage can't find identity image in db
	InvalidIdentityImage = "invalid_identity_image"
	// InvalidOcr identity number different with identity image
	InvalidOcr = "invalid_ocr"
	// InvalidSelfieImage can't find selfie image in db
	InvalidSelfieImage = "invalid_selfie_image"
	// InvalidFaceRecognition face different between identity and selfie image
	InvalidFaceRecognition = "invalid_face_recognition"
	// GenderNotFound gender not found when input dummy data after insert identity
	GenderNotFound = "gender_not_found"
	// GetDukcapilData get dukcapil data fail
	GetDukcapilData = "get_dukcapil_data"
	// EmptyDukcapilData empty dukcapil data
	EmptyDukcapilData = "empty_dukcapil_data"
	// InvalidName invalid name when pick from randomize dummy name
	InvalidName = "invalid_name"
	// MaxDetail max detail is 3 records
	MaxDetail = "max_detail"
	// InvalidSignatureImage can't find signature image in db
	InvalidSignatureImage = "invalid_signature_image"
	// InvalidNpwpImage can't find npwp image in db
	InvalidNpwpImage = "invalid_npwp_image"
	// InvalidNpwpOcr npwp number different with npwp image
	InvalidNpwpOcr = "invalid_npwp_ocr"
	// CountryTaxDetail Invalid country in tax details
	CountryTaxDetail = "country_tax_detail"
	// BaFormat Invalid bank account number format
	BaFormat = "ba_format"
	// MailingAddressType Invalid mailing address type
	MailingAddressType = "mailing_address_type"
	// JobStartAt Invalid job start at
	JobStartAt = "job_start_at"
	// InvalidBirthDate invalid birth date format
	InvalidBirthDate = "invalid_birth_date"
	// DukcapilComplete all dukcapil data is complete
	DukcapilComplete = "dukcapil_complete"
	// FillAllData You must fill all your data first
	FillAllData = "fill_all_data"
	// UmTCP error on um tcp request
	UmTCP = "um_tcp"
	// EmailActivated email already registered
	EmailActivated = "email_activated"
	// PasswordLength password length must 6-10 characters
	PasswordLength = "password_length"
	// PasswordFormat Password must consists 6-10 characters and includes at least 3 criterias : uppercase letters, lowercase letters, numbers, symbols
	PasswordFormat = "password_format"
	// OAOFile failed when upload file to oao folder
	OAOFile = "oao_file"
	// UserLocked user locked
	UserLocked = "user_locked"
	// NpwpFormat invalid npwp number format
	NpwpFormat = "npwp_format"
	// RedirectToLogin when user register using existing email
	RedirectToLogin = "redirect_to_login"
	// Recaptcha recaptcha error
	Recaptcha = "recaptcha"
	// InvalidSpouseName invalid or empty spouse name
	InvalidSpouseName = "invalid_spouse_name"
	// InvalidSpousePhone invalid or empty spouse phone
	InvalidSpousePhone = "invalid_spouse_phone"
	// InvalidSpouseEmail invalid or empty spouse email
	InvalidSpouseEmail = "invalid_spouse_email"
	// UnfinishedApplication have ducplicate on going opening account application
	UnfinishedApplication = "unfinished_application"
	// PlafondTooBig margin request bigger than plafond
	PlafondTooBig = "plafond_too_big"
	// PlafondTooSmall margin request smaller than min plafond
	PlafondTooSmall = "plafond_too_small"
	// NoteAlreadyFilled note already filled by another admin
	NoteAlreadyFilled = "note_already_filled"
	// InvalidStatus invalid status
	InvalidStatus = "invalid_status"
	// MarginConfirmationMail push margin confirmation mail request to queue
	MarginConfirmationMail = "margin_confirmation_mail"
	// InvalidReminderCount invalid reminder count
	InvalidReminderCount = "invalid_reminder_count"
	// ExpKey ...
	ExpKey = "exp_key"
	// InvalidDate ...
	InvalidDate = "invalid_date"
	// SpouseEmailSameWithUser ...
	SpouseEmailSameWithUser = "spouse_email_same_with_user"
	// MaxSpouseRejected ...
	MaxSpouseRejected = "max_spouse_rejected"
	// InvalidImageType ...
	InvalidImageType = "invalid_image_type"
	// RepairUserData push mail request to queue ...
	RepairUserData = "repair_user_data_mail"
	// ReviewDataAdmin ...
	ReviewDataAdmin = "Not found repair data"
	// PhoneEmailExist ...
	PhoneEmailExist = "phone_email_exist"
	// PrivyAlreadyRegistered ...
	PrivyAlreadyRegistered = "privy_already_registered"
	// CodeExist ...
	CodeExist = "code_exist"
	// DefaultPartnerExist ...
	DefaultPartnerExist = "default_partner_exist"
	// DocumentAlreadyGenerated ...
	DocumentAlreadyGenerated = "document_already_generated"
	// AlreadySubmitMargin ...
	AlreadySubmitMargin = "already_submit_margin"
	// LockAdmin ...
	LockAdmin = "lock_admin"
	// AlreadyGetZoloz ...
	AlreadyGetZoloz = "already_get_zoloz"
	// InvalidRole ...
	InvalidRole = "invalid_role"
	// ZolozError ...
	ZolozError = "zoloz_error"
	// DocumentUnsigned ...
	DocumentUnsigned = "document_unsigned"
	// ExpPassword expired password
	ExpPassword = "exp_password"
	// MaxSendEmail Max Send email
	MaxSendEmail = "max_send_email"
	// AutoVerifyTurnedOff ...
	AutoVerifyTurnedOff = "auto_verify_turned_off"
	// OneAccountMax ...
	OneAccountMax = "one_account_max"
	// PasswordAlreadyUsed ...
	PasswordAlreadyUsed = "password_already_used"
	// PinLength ...
	PinLength = "pin_length"
	// DuplicateEmail ...
	DuplicateEmail = "duplicate_email"
	// DuplicatePhoneNumber ...
	DuplicatePhoneNumber = "duplicate_phone_number"
	// InvalidPhoneNumber ...
	InvalidPhoneNumber = "invalid_phone_number"
	// InvalidCredentials ...
	InvalidCredentials = "invalid_credentials"
	// InvalidAuth ...
	InvalidAuth = "invalid_auth"
	// StoreFailed ...
	StoreFailed = "store_failed"
	// InactiveAdmin ...
	InactiveAdmin = "inactive_admin"
	// InvalidAmount ...
	InvalidAmount = "invalid_amount"
	// TransactionNotFound ...
	TransactionNotFound = "transaction_not_found"
	// TagNotFound ...
	TagNotFound = "tag_not_found"
	//TagsTooManyInput ...
	TagsTooManyInput = "tags_too_many"
	//PinSame ...
	PinSame = "pin_same"
	//InvalidOldPin ...
	InvalidOldPin = "invalid_old_pin"
	//InvalidLogin ...
	InvalidLogin = "phone_or_pin_not_match"
	//InvalidPin ...
	InvalidPin = "pin_not_match"
	//PinAlreadySet ...
	PinAlreadySet = "pin_already_set"
	//PinUnset ...
	PinUnset = "pin_unset"
	//UserAlreadyValidated ...
	UserAlreadyValidated = "user_already_validated"
	// BankNotFound ...
	BankNotFound = "bank_not_found"
	// PhoneNotValidated ...
	PhoneNotValidated = "phone_not_validated"
)
