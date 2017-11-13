package stackserver

//ServerConfig server configuration
type ServerConfig struct {
	Port int
}

//DefaultConfig returns a default configuration
func DefaultConfig() *ServerConfig {
	return &ServerConfig{9009}
}
