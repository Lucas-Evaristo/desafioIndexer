package model

type BulkEmail struct {
	Index  string              `json:"Index"`
	Emails []map[string]string `json:"Records"`
}

type Email struct {
	MessageID       string `json:"Message-ID"`
	Date            string `json:"Date"`
	From            string `json:"From"`
	To              string `json:"To"`
	Subject         string `json:"Subject"`
	Cc              string `json:"Cc"`
	MimeV           string `json:"Mime-Version"`
	ContentType     string `json:"Content-Type"`
	ContentEncoding string `json:"Content-Transfer-Encoding"`
	Bcc             string `json:"Bcc"`
	XFrom           string `json:"X-From"`
	XTo             string `json:"X-To"`
	Xcc             string `json:"X-cc"`
	Xbcc            string `json:"X-bcc"`
	XFolder         string `json:"X-Folder"`
	XOrigin         string `json:"X-Origin"`
	XFileName       string `json:"X-Filename"`
	Body            string `json:"Body"`
}

// type EmailFields struct {
// 	fields map[string]string
// }

func NewEmailFields() map[string]string {
	ef := make(map[string]string)

	ef["Message-ID"] = "MessageID"
	ef["Date"] = "Date"
	ef["From"] = "From"
	ef["To"] = "To"
	ef["Subject"] = "Subject"
	ef["Cc"] = "Cc"
	ef["Mime-Version"] = "MimeV"
	ef["Content-Type"] = "ContentType"
	ef["Content-Transfer-Encoding"] = "ContentEncoding"
	ef["Bcc"] = "Bcc"
	ef["X-From"] = "XFrom"
	ef["X-To"] = "XTo"
	ef["X-cc"] = "Xcc"
	ef["X-bcc"] = "Xbcc"
	ef["X-Folder"] = "XFolder"
	ef["X-Origin"] = "XOrigin"
	ef["X-Filename"] = "XFileName"

	return ef
}
