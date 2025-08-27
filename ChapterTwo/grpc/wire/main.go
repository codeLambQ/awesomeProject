package main

type Config struct {
	DbIp   string
	DbPort string

	RedisIp   string
	RedisPort string
}

type MysqlClient struct{}

type RedisClient struct{}

func NewConfig() *Config {
	return &Config{}
}

func NewMysqlClient(c *Config) *MysqlClient {
	mysqlClient := &MysqlClient{}
	return mysqlClient
}

func NewRedisClient(c *Config) *RedisClient {
	redisClient := &RedisClient{}
	return redisClient
}

type App struct{}

func NewApp(m *MysqlClient, r *RedisClient) *App {
	return &App{}
}

func main() {
	app := WireApp()
	println(app)
}
