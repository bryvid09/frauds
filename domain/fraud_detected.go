package domain

type FraudDetected struct {
	Period   string
	Value    int
	Client   Client
}