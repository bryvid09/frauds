package infrastructure

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bryvid09/frauds/domain"
)

type CSVAdapter struct{}

func (a *CSVAdapter) Read(file string) ([]domain.Reading, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // Allowing variable number of fields per record

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var readings []domain.Reading
	for _, record := range records {
		readInt, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}
		reading := domain.Reading{
			Period: record[1],
			Reader: readInt,
			Client: domain.Client(record[0]),
		}
		readings = append(readings, reading)
	}

	return readings, nil
}