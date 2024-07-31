package models

type ApiResponse struct {
	ApiLimit  int            `json:"api_limit"`
	ApiCount  int            `json:"api_count"`
	Page      int            `json:"page"`
	Results   int            `json:"results"`
	Links     Links          `json:"_links"`
	List      []DataResponse `json:"list"`
	Set       int            `json:"set"`
	Streaming int            `json:"streaming"`
	Region    int            `json:"region"`
	Date      string         `json:"date"`
	Type      int            `json:"type"`
	Country   int            `json:"country"`
	Year      int            `json:"year"`
	Genre     int            `json:"genre"`
	Company   int            `json:"company"`
	Filter    interface{}    `json:"filter"`
	Query     interface{}    `json:"query"`
	Grouping  int            `json:"grouping"`
	Top25     string         `json:"top25"`
	Updated   interface{}    `json:"updated"`
}

type Links struct {
	Self  Link `json:"self"`
	First Link `json:"first"`
	Last  Link `json:"last"`
	Next  Link `json:"next"`
}

type Link struct {
	Href string `json:"href"`
}

type DataResponse struct {
	Result      int    `json:"result"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Premiere    string `json:"premiere"` // 날짜가 문자열로 제공됨
	TypeID      int    `json:"type_id"`
	Type        string `json:"type"`
	CountryID   int    `json:"country_id"`
	Country     string `json:"country"`
	CompanyID   int    `json:"company_id"`
	Company     string `json:"company"`
	Key         string `json:"key"`
	Note        string `json:"note"`
	Region      int    `json:"region"` // 지역은 정수로 제공됨
	Ranking     int    `json:"ranking"`
	RankingLast int    `json:"ranking_last"`
	Value       int    `json:"value"`
	ValueLast   int    `json:"value_last"`
	ValueTotal  int    `json:"value_total"`
	Countries   int    `json:"countries"`
	Days        int    `json:"days"`
}

// PreferencesResponse represents the response structure for preferences API requests
type PreferencesResponse struct {
	ApiLimit  int              `json:"api_limit"`
	ApiCount  int              `json:"api_count"`
	Sum       int              `json:"sum"`
	Page      int              `json:"page"`
	Results   int              `json:"results"`
	List      []PreferenceItem `json:"list"`
	Set       int              `json:"set"`
	Streaming int              `json:"streaming"`
	Region    int              `json:"region"`
	Date      string           `json:"date"`
}

// PreferenceItem represents each item in the preferences response list
type PreferenceItem struct {
	Result int     `json:"result"`
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	GID    *int    `json:"gid"` // Nullable int, can be null
	Group  string  `json:"group"`
	Value  int     `json:"value"`
	Share  float64 `json:"share"`
}
