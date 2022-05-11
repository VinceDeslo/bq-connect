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

func main() {
	fmt.Println("Hello Go")

	ctx := context.Background()

	projectID := "bqconnect-349116"
	queryString := `
		SELECT *
		FROM bq_plants_dataset.bq_plants_table 
		LIMIT 25;`

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Bigquery client creation failed: %v", err)
	}
	defer client.Close()

	query := client.Query(queryString)

	rows, err := query.Read(ctx)
	if err != nil {
		log.Fatalf("Error while querying dataset: %v", err)
	}

	if err := printRowResults(os.Stdout, rows); err != nil {
		log.Fatal(err)
	}
}

func printRowResults(w io.Writer, rows *bigquery.RowIterator) error {
	for {
		var row []bigquery.Value
		err := rows.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, row[0])
	}
	fmt.Fprintf(w, "Row Count: %d\n", rows.TotalRows)
	return nil
}
