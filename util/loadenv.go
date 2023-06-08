package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	envs := make(map[string]string)
	envs["K_USERID"] = os.Getenv("K_USERID")
	envs["K_PASSWORD"] = os.Getenv("K_PASSWORD")
	envs["DB"] = os.Getenv("DB")
	envs["DB_DIALECT"] = os.Getenv("DB_DIALECT")
	return envs

}
