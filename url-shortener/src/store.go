package main

type Storage interface {
	CreateEntry(string, string) error
	GetEntry(string) (string, error)
	Close() error
}

func NewStorage() (Storage, error) {
	var s Storage
	if (conf.Backend == "memory") {
		s = nil
	} else {
		s, _ = NewDB()
	}
	return s, nil
}