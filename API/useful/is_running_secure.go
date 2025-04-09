package useful

import (
	"os"
)

func IsRunningSecure() bool {
	LoadEnv()
	feUrl := os.Getenv("FRONTEND_URL")
	// Check if frontend URL starts with https
	return len(feUrl) > 5 && feUrl[:5] != "https"
}
