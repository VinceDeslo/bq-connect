package domain

import "cloud.google.com/go/bigquery"

type Game struct {
	Name            bigquery.NullString
	Platform        bigquery.NullString
	Year_of_release bigquery.NullInt64
	Genre           bigquery.NullString
	Publisher       bigquery.NullString
	Na_sales        bigquery.NullFloat64
	Eu_sales        bigquery.NullFloat64
	Jp_sales        bigquery.NullFloat64
	Other_sales     bigquery.NullFloat64
	Global_sales    bigquery.NullFloat64
	Critic_score    bigquery.NullInt64
	Critic_count    bigquery.NullInt64
	User_score      bigquery.NullString
	User_count      bigquery.NullInt64
	Developer       bigquery.NullString
	Rating          bigquery.NullString
}
