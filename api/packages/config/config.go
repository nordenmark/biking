package config

type ServerConfiguration struct {
	Port             int
	CorsAllowOrigins []string
}

func NewServerConfig() ServerConfiguration {
	return ServerConfiguration{
		Port:             getInt("PORT", 1323),
		CorsAllowOrigins: getStringList("CORS_ALLOW_ORIGINS", []string{"http://127.0.0.1:5173", "http://localhost:5173"}),
	}
}
