package models

import "time"

type Response struct {
	Timed_out bool    `json:"timed_out"`
	Hits      Hits    `json:"hits"`
	Took      float64 `json:"took"`
	Shards    Shards  `json:"_shards"`
}

type Total struct {
	Value float64 `json:"value"`
}

type Shards struct {
	Total      float64 `json:"total"`
	Successful float64 `json:"successful"`
	Skipped    float64 `json:"skipped"`
	Failed     float64 `json:"failed"`
}

type Hits struct {
	Total     Total     `json:"total"`
	Max_score float64   `json:"max_score"`
	Id        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp time.Time `json:"@timestamp"`
	Hits      []Hits    `json:"hits"`
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	Source    Source    `json:"_source"`
}

type Source struct {
	Message_ID                string    `json:"Message-ID"`
	Subject                   string    `json:"Subject"`
	To                        string    `json:"To"`
	From                      string    `json:"From"`
	Bcc                       string    `json:"Bcc"`
	Date                      string    `json:"Date"`
	Timestamp                 time.Time `json:"@timestamp"`
	Cc                        string    `json:"Cc"`
	Body                      string    `json:"Body"`
	Content_Transfer_Encoding string    `json:"Content-Transfer-Encoding"`
	Mime_Version              string    `json:"Mime-Version"`
	X_Folder                  string    `json:"X-Folder"`
	X_From                    string    `json:"X-From"`
	X_cc                      string    `json:"X-cc"`
	X_FileName                string    `json:"X-FileName"`
	X_Origin                  string    `json:"X-Origin"`
	X_To                      string    `json:"X-To"`
	X_bcc                     string    `json:"X-bcc"`
}
