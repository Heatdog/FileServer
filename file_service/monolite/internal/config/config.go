package config

type Config struct {
	FileStorage FileStorage
}

type FileStorage struct {
	Location string
}
