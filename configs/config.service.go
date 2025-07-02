package configs

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func Load() {
	envFile := getEnvFilePath()

	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}
}

func getEnvFileName() string {
	fileName := ".env"

	env := Get("APP_ENV")
	if len(env) > 0 && env != "prod" {
		fileName = fmt.Sprintf("%s.%s", fileName, env)
	}

	return fileName
}

func getEnvFileDir() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(execPath)
}

func getEnvFilePath() string {
	execDir := getEnvFileDir()
	fileName := getEnvFileName()
	filePath := filepath.Join(execDir, fileName)

	return filePath
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetOrPanic(key string) string {
	value := Get(key)
	if len(value) == 0 {
		message := fmt.Sprintf("Not found env variable with Key: %s", key)
		panic(errors.New(message))
	}

	return value
}
