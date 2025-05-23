package config

type MariadbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadMySQLConfig() *MariadbConfig {
	return &MariadbConfig{
		User:     GetEnv("DB_USER", "root"),
		Password: GetEnv("DB_PASSWORD", "example"),
		Host:     GetEnv("DB_HOST", "192.168.106.3"),
		Port:     GetEnv("DB_PORT", "3306"),
		DBName:   GetEnv("DB_NAME", "yt_db"),
	}
}
