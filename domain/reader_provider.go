package domain

type ReaderProvider interface {
	GetReadingsForClient(client Client)
	GetAllReadings()
}