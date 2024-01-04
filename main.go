package main

import (
	"fmt"

	"github.com/bryvid09/frauds/infrastructure"
)

func main() {
// 	client, err := domain.NewClient("")
// 	if err != nil {
// 		panic(err)
// 	}
// reading := domain.Reading{
// 	Client: client,
// 	Period:   "2018-01",
// 	Reader:    100,
// }
xmlAdapter := &infrastructure.XMLAdapter{}
readingsXml, err := xmlAdapter.Read("2016-readings.xml")
if err != nil {
	panic(err)
}
for _, reading := range readingsXml {
    fmt.Printf("ClientID: %s, Periodo: %s, Lectura: %d\n", reading.Client, reading.Period, reading.Reader)
}

csvAdapter := &infrastructure.CSVAdapter{}

readingsCsv, err := csvAdapter.Read("2016-readings.csv")
if err != nil {
	panic(err)
}

for _, reading := range readingsCsv {
	fmt.Printf("ClientID: %s, Periodo: %s, Lectura: %d\n", reading.Client, reading.Period, reading.Reader)
}

}