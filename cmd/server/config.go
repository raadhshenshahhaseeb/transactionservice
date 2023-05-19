package server

type Config struct {
	Address   string `envconfig:"ADDRESS" default:"0.0.0.0:8000"`
	JwtSecret string `envconfig:"JWT_SECRET" required:"true"`
	Database  struct {
		Host       string `envconfig:"DB_HOST" required:"true"`
		Password   string `envconfig:"DB_PASSWORD" required:"true"`
		Port       string `envconfig:"DB_PORT" required:"true"`
		DBName     string `envconfig:"DB_NAME" required:"true"`
		DBUsername string `envconfig:"DB_USERNAME" required:"true"`
	}
	Logger struct {
		LogEnv   string `envconfig:"LOG_ENV" default:"local"`
		LogLevel string `envconfig:"LOG_LEVEL" default:"4"`
	}
	Chain struct {
		RPCEndpoint string `envconfig:"RPC_ENDPOINT"`
		ChainID     int64  `envconfig:"CHAIN_ID"`
	}
	PrivateKey string `envconfig:"PRIVATE_KEY"`
}
