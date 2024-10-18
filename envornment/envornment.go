package envornment

import (
	"fmt"
	"os"
)

func GetMongoURI() string {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return ""
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return ""
	}

	clientURI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.yud8r.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", name, password)
	return clientURI
}

func GetJwtKey() string {
	return os.Getenv("JWT_KEY")
}

func GetUserCollection() string {
	return os.Getenv("USER_COLLECTION")
}

func GetTaskCollection() string {
	return os.Getenv("TASK_COLLECTION")
}
