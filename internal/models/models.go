package models

type DataResponse struct {
	Result      int    `json:"result"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Premiere    int    `json:"premiere"`
	TypeID      int    `json:"type_id"`
	Type        string `json:"type"`
	CountryID   int    `json:"country_id"`
	Country     string `json:"country"`
	CompanyID   int    `json:"company_id"`
	Company     string `json:"company"`
	Key         string `json:"key"`
	Note        string `json:"note"`
	Region      string `json:"region"`
	Ranking     int    `json:"ranking"`
	RankingLast int    `json:"ranking_last"`
	Value       int    `json:"value"`
	ValueLast   int    `json:"value_last"`
	ValueTotal  int    `json:"value_total"`
	Countries   int    `json:"countries"`
	Days        int    `json:"days"`
}

type DemographicsResponse struct {
	Result      int     `json:"result"`
	MovieID     *int    `json:"movie_id"`
	Movie       *string `json:"movie"`
	FranchiseID *int    `json:"franchise_id"`
	Franchise   *string `json:"franchise"`
	URL         string  `json:"url"`
	CountryID   int     `json:"country_id"`
	Country     string  `json:"country"`
	Generation  int     `json:"generation"`
	Gender      int     `json:"gender"`
	Value       int     `json:"value"`
	Share       float64 `json:"share"`
}

type PreferencesResponse struct {
	Sum    int `json:"sum"`
	Result []struct {
		ID    int     `json:"id"`
		Name  string  `json:"name"`
		GID   int     `json:"gid"`
		Group string  `json:"group"`
		Value int     `json:"value"`
		Share float64 `json:"share"`
	} `json:"result"`
}

type SearchResponse struct {
	Result     int    `json:"result"`
	ID         int    `json:"id"`
	IMDB       string `json:"imdb"`
	TMDB       int    `json:"tmdb"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	API        string `json:"api"`
	Premiere   string `json:"premiere"`
	Popularity int    `json:"popularity"`
}

type TitleResponse struct {
	ID        int    `json:"id"`
	IMDB      string `json:"imdb"`
	TMDB      int    `json:"tmdb"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Premiere  string `json:"premiere"`
	TypeID    int    `json:"type_id"`
	Type      string `json:"type"`
	CountryID int    `json:"country_id"`
	Country   string `json:"country"`
	GenreID   int    `json:"genre_id"`
	Genre     string `json:"genre"`
	CompanyID int    `json:"company_id"`
	Company   string `json:"company"`
	Result    []struct {
		Key         string `json:"key"`
		Streaming   int    `json:"streaming"`
		Region      int    `json:"region"`
		From        string `json:"from"`
		To          string `json:"to"`
		Ranking     int    `json:"ranking"`
		RankingLast int    `json:"ranking_last"`
		Value       int    `json:"value"`
		ValueLast   int    `json:"value_last"`
		ValueTotal  int    `json:"value_total"`
		Countries   int    `json:"countries"`
		Days        int    `json:"days"`
	} `json:"result"`
}

type TrendingResponse struct {
	Result      int     `json:"result"`
	MovieID     *int    `json:"movie_id"`
	Movie       *string `json:"movie"`
	FranchiseID *int    `json:"franchise_id"`
	Franchise   *string `json:"franchise"`
	URL         string  `json:"url"`
	Region      int     `json:"region"`
	Value       int     `json:"value"`
	From        string  `json:"from"`
	To          string  `json:"to"`
}
