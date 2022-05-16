package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type VideoGameRow struct {
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

const projectID = "bqconnect-349116"
const queryString = `
	SELECT * 
	FROM bq_video_games_dataset.bq_video_games_table
	WHERE genre = "Adventure" 
		AND platform = "GBA"
		AND year_of_release IS NOT NULL
	ORDER BY year_of_release
	LIMIT 100`

func main() {
	fmt.Println("Hello Game Advisor")

	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Bigquery client creation failed: %v", err)
	}
	defer client.Close()

	query := client.Query(queryString)

	iter, err := query.Read(ctx)
	if err != nil {
		log.Fatalf("Error while querying dataset: %v", err)
	}

	if err := printResults(os.Stdout, iter); err != nil {
		log.Fatal(err)
	}
}

func printResults(w io.Writer, rows *bigquery.RowIterator) error {
	for {
		var row VideoGameRow
		err := rows.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error iterating through results: %v", err)
		}
		fmt.Fprintf(w, "name: %s, year of release: %d\n", row.Name, row.Year_of_release.Int64)
	}
}
