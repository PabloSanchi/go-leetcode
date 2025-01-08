package constants

import "os"

var (
	VERSION     = getEnv("version", "v1")
	PORT        = getEnv("port", ":8080")
	AUTH_COOKIE = getEnv("auth_cookie", "auth")
	SECRET_KEY  = []byte(getEnv("secret", "secret-key"))
	CSRF_KEY    = []byte(getEnv("csrf_key", "csrf-key"))
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
