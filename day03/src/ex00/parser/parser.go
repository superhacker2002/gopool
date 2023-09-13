package parser

import (
	"encoding/csv"
	"errors"
	"ex00/entity"
	"fmt"
	"io"
	"strconv"
)

var ErrReadFailed = errors.New("failed to read record from .csv file")

func ReadAll(r io.Reader) ([]entity.Restaurant, error) {
	var restaurants []entity.Restaurant
	reader := csv.NewReader(r)
	reader.Comma = '\t'

	record, err := reader.Read()
	if err != nil {
		if err != io.EOF {
			return nil, fmt.Errorf("%w: %v", ErrReadFailed, err)
		}
	}

	for {
		record, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("%w: %v", ErrReadFailed, err)
		}
		restaurants = append(restaurants, recordToEntity(record))
	}

	return restaurants, nil
}

func recordToEntity(record []string) entity.Restaurant {
	id, _ := strconv.Atoi(record[0])
	lon, _ := strconv.ParseFloat(record[4], 64)
	lat, _ := strconv.ParseFloat(record[5], 64)

	return entity.New(id, record[1], record[2], record[3], lon, lat)
}
