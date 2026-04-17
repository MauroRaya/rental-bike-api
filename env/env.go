package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresPort     string
	DSN              string
	Port             string
}

func mustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("missing required env: %s", key))
	}
	return v
}

func Load() (Env, error) {
	err := godotenv.Load()
	if err != nil {
		return Env{}, err
	}

	env := Env{
		PostgresUser:     mustGet("POSTGRES_USER"),
		PostgresPassword: mustGet("POSTGRES_PASSWORD"),
		PostgresDB:       mustGet("POSTGRES_DB"),
		PostgresPort:     mustGet("POSTGRES_PORT"),
		Port:             mustGet("PORT"),
	}

	env.DSN = fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.PostgresPort,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
	)

	return env, nil
}
