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
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "example"),
		Host:     getEnv("DB_HOST", "192.168.106.3"),
		Port:     getEnv("DB_PORT", "3306"),
		DBName:   getEnv("DB_NAME", "yt_db"),
	}
}
