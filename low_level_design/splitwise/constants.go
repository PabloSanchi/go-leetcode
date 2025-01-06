package constants

import "os"

var (
	VERSION     string = getEnv("version", "v1")
	PORT        string = getEnv("port", ":8080")
	AUTH_COOKIE string = getEnv("auth_cookie", "auth")
	SECRET_KEY  []byte = []byte(getEnv("secret", "secret-key"))
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
