package elasticclient

import (
	"bytes"
	"context"
	"encoding/json"
	"ex00/entity"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/olivere/elastic"
	"golang.org/x/sync/errgroup"
	"strconv"

	"strings"
)

type Schema struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name     Type `json:"name"`
	Address  Type `json:"address"`
	Phone    Type `json:"phone"`
	Location Type `json:"location"`
	Id       Type `json:"id"`
}

type Restaurant struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

type Type struct {
	D string `json:"type"`
}

type ElasticClient struct {
	es *elasticsearch.Client
}

func New() (ElasticClient, error) {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return ElasticClient{}, fmt.Errorf("failed to create elastic search client: %w", err)
	}
	return ElasticClient{es: esClient}, nil
}

func (e ElasticClient) CreateIndex(name string) error {
	index, err := e.es.Indices.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create elastic search index: %w", err)
	}

	err = index.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to close elastic search response: %w", err)
	}

	return nil
}

func (e ElasticClient) PutMapping(index, doc string) error {
	var buf bytes.Buffer
	b := Schema{Properties{
		Name:     Type{"text"},
		Address:  Type{"text"},
		Phone:    Type{"text"},
		Id:       Type{"long"},
		Location: Type{"geo_point"},
	}}

	err := json.NewEncoder(&buf).Encode(b)
	if err != nil {
		return fmt.Errorf("failed to encode mapping structure to json: %w", err)
	}

	res, err := e.es.Indices.PutMapping(
		strings.NewReader(buf.String()),
		e.es.Indices.PutMapping.WithIndex(index),
		e.es.Indices.PutMapping.WithDocumentType(doc),
	)

	err = res.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to close elastic search response: %w", err)
	}

	return nil
}

func (e ElasticClient) SendData(restaurants []entity.Restaurant, index string) error {
	var g errgroup.Group

	for i, r := range restaurants {
		i, r := i, r
		g.Go(func() error {
			res, err := json.Marshal(entityToDTO(r))
			if err != nil {
				return err
			}
			request := esapi.IndexRequest{
				Index:      index,
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(string(res)),
				Refresh:    "true",
			}

			response, err := request.Do(context.Background(), e.es)
			if err != nil {
				return err
			}
			if response.IsError() {
				return fmt.Errorf("error while indexing document: %w", err)
			}
			err = response.Body.Close()
			if err != nil {
				return err
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func entityToDTO(r entity.Restaurant) Restaurant {
	return Restaurant{
		ID:       r.ID,
		Name:     r.Name,
		Address:  r.Address,
		Phone:    r.Phone,
		Location: elastic.GeoPoint{Lat: r.Latitude, Lon: r.Longitude},
	}
}
