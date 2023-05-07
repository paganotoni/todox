package envor

import "os"

func Get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
