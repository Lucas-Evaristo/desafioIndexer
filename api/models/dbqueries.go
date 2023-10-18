package models

type Search struct {
	Search_type  string   `json:"search_type"`
	Search_query Query    `json:"query"`
	From         int      `json:"from"`
	Sort_fields  []string `json:"sort_fields"`
	Max_results  int      `json:"max_results"`
	Source       []string `json:"_source"`
}

type Query struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}
