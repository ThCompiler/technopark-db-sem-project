package internal

type RepositoryConnections struct {
	DataBaseUrl string `toml:"database_url"`
}

type Config struct {
	LogLevel   string                `toml:"log_level"`
	LogAddr    string                `toml:"log_path"`
	Domen      string                `toml:"domen"`
	BindAddr   string                `toml:"bind_addr"`
	Repository RepositoryConnections `toml:"repository"`
}
