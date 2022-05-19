package driven

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/VinceDeslo/bq-connect/internal/core/domain"
	"github.com/VinceDeslo/bq-connect/internal/core/queries"
	"google.golang.org/api/iterator"
)

type BQAdapter struct {
	context   context.Context
	projectID string
	Client    *bigquery.Client
}

func NewBQAdapter(gcpProjectID string) *BQAdapter {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, gcpProjectID)
	if err != nil {
		log.Fatalf("Bigquery client creation failed: %v", err)
	}

	return &BQAdapter{
		context:   ctx,
		projectID: gcpProjectID,
		Client:    client,
	}
}

func (adapter *BQAdapter) Close() error {
	return adapter.Client.Close()
}

func (adapter *BQAdapter) GetAll() ([]domain.Game, error) {
	qstring := queries.GetAllQuery

	iter, err := adapter.queryBq(qstring)
	if err != nil {
		log.Fatalf("Error while querying BigQuery: %v", err)
	}
	results, err := formatBqResponse(iter)
	if err != nil {
		log.Fatalf("Error while formatting response: %v", err)
	}

	return results, nil
}

func (adapter *BQAdapter) GetAdventure() ([]domain.Game, error) {
	qstring := queries.GetAdventureQuery

	iter, err := adapter.queryBq(qstring)
	if err != nil {
		log.Fatalf("Error while querying BigQuery: %v", err)
	}
	results, err := formatBqResponse(iter)
	if err != nil {
		log.Fatalf("Error while formatting response: %v", err)
	}

	return results, nil
}

func (adapter *BQAdapter) queryBq(qstring string) (*bigquery.RowIterator, error) {
	query := adapter.Client.Query(qstring)

	iter, err := query.Read(adapter.context)
	if err != nil {
		log.Fatalf("Error while reading query: %v", err)
	}
	return iter, nil
}

func formatBqResponse(iter *bigquery.RowIterator) ([]domain.Game, error) {
	results := []domain.Game{}
	for {
		var game domain.Game
		err := iter.Next(&game)
		if err == iterator.Done {
			return results, nil
		}
		if err != nil {
			return []domain.Game{}, fmt.Errorf("error iterating through BigQuery results: %v", err)
		}
		results = append(results, game)
	}
}
