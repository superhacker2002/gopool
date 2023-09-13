package elasticclient

import (
	"bytes"
	"encoding/json"
	"ex00/entity"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
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
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
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
	buf := entitiesToBulk(restaurants, index)
	res, err := e.es.Bulk(bytes.NewReader(buf.Bytes()), e.es.Bulk.WithIndex(index))
	if err != nil {
		return errors.Errorf("failed to send Bulk request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.Errorf("bulk request failed: %s", res.String())
	}

	return nil
}

func entitiesToBulk(restaurants []entity.Restaurant, index string) *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})

	for _, r := range restaurants {
		jsonData, err := json.Marshal(entityToDTO(r))
		if err != nil {
			fmt.Printf("failed to marshal document: %s\n", err)
			return nil
		}

		actionLine := fmt.Sprintf(`{ "index" : { "_index" : "%s", "_id" : "%d" } }`, index, r.ID)
		buf.WriteString(actionLine)
		buf.WriteString("\n")

		buf.Write(jsonData)
		buf.WriteString("\n")
	}

	return buf
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
