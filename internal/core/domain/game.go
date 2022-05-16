package domain

type Game struct {
	Name            string
	Platform        string
	Year_of_release int64
	Genre           string
	Publisher       string
	Na_sales        float64
	Eu_sales        float64
	Jp_sales        float64
	Other_sales     float64
	Global_sales    float64
	Critic_score    int64
	Critic_count    int64
	User_score      string
	User_count      int64
	Developer       string
	Rating          string
}
