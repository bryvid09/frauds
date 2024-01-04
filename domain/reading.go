package domain

type Reading struct {
	Period string
	Reader int
	Client Client
}

type Readings []Reading