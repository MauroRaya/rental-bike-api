package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PostgresHost      string
	PostgresUser      string
	PostgresPassword  string
	PostgresDB        string
	PostgresPort      string
	GooseDriver       string
	GooseDbstring     string
	GooseMigrationDir string
	DSN               string
	JwtSecret         string
	Port              string
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
		PostgresHost:      mustGet("POSTGRES_HOST"),
		PostgresUser:      mustGet("POSTGRES_USER"),
		PostgresPassword:  mustGet("POSTGRES_PASSWORD"),
		PostgresDB:        mustGet("POSTGRES_DB"),
		PostgresPort:      mustGet("POSTGRES_PORT"),
		GooseDriver:       mustGet("GOOSE_DRIVER"),
		GooseDbstring:     mustGet("GOOSE_DBSTRING"),
		GooseMigrationDir: mustGet("GOOSE_MIGRATION_DIR"),
		JwtSecret:         mustGet("JWT_SECRET"),
		Port:              mustGet("PORT"),
	}

	env.DSN = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.PostgresHost,
		env.PostgresPort,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
	)

	return env, nil
}
