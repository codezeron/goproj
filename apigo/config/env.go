package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost 		string
	Port       		string
	DBUser     		string
	DBPasswd   		string
	DBAddress  		string
	DBName     		string
	JWTExpiration int64
	JWTSecret     string
}

// serve para nao inicializar toda vez que for chamado
var Envies = initConfig()

// set default value
func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: 		getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       		getEnv("PORT", "8080"),
		DBUser:     		getEnv("DB_USER", "root"),
		DBPasswd:   		getEnv("DB_PASSWD", "password"),
		DBAddress:  		fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), 		getEnv("DB_PORT", "3306")),
		DBName:     		getEnv("DB_NAME", "apigo"),
		JWTSecret: 			getEnv("JWT_SECRET", "not-secret?"),
		JWTExpiration:  getEnvInt("JWT_EXP", 3600 * 24 *7),//7 days
	}
}

func getEnv(key, content string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return content
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok{
		num, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return num //retorna interiro
	}
	return fallback
}