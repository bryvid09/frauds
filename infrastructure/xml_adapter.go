package infrastructure

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/bryvid09/frauds/domain"
)

type XMLAdapter struct{}
type XMLReadings struct {
	XMLName  xml.Name     `xml:"readings"`
	Readings []XMLReading `xml:"reading"`
}
type XMLReading struct {
	XMLName  xml.Name `xml:"reading"`
	ClientID string   `xml:"clientID,attr"`
	Period   string   `xml:"period,attr"`
	Value    int      `xml:",chardata"`
}

func (a *XMLAdapter) Read(file string) ([]domain.Reading, error) {
	xmlFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var xmlContainer XMLReadings
	err = xml.Unmarshal(byteValue, &xmlContainer)
	if err != nil {
		return nil, err
	}

	var readings []domain.Reading
	for _, xmlReading := range xmlContainer.Readings {
		reading := domain.Reading{
			Period: xmlReading.Period,
			Reader: xmlReading.Value,
			Client: domain.Client(xmlReading.ClientID),
		}
		readings = append(readings, reading)
	}

	return readings, nil
}
