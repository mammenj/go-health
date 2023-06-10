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
	envs["DB_PROCS"] = os.Getenv("DB_PROCS")
	envs["DB_DIALECT_PROCS"] = os.Getenv("DB_DIALECT_PROCS")
	envs["DB_HOSPITAL"] = os.Getenv("DB_HOSPITAL")
	envs["DB_DIALECT_HOSPITAL"] = os.Getenv("DB_DIALECT_HOSPITAL")
	return envs

}
