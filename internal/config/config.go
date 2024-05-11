package config

type Config struct {
	Port   string
	DbPath string
}

var Settings = Config{
	Port:   ":8080",
	DbPath: "db.sqlite",
}
