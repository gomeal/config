package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var _ SecretClient = &secretClientImpl{}

type (
	secretValue string

	secretClientImpl struct {
		secrets map[Secret]secretValue
	}
)

func newSecretClient() *secretClientImpl {
	s := &secretClientImpl{
		secrets: make(map[Secret]secretValue),
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("unable to os.Getwd() in config.newSecretClient: %w", err))
	}

	envPath, err := filepath.Abs(filepath.Join(wd, ".env"))
	if err != nil {
		panic(fmt.Errorf("unable to filepath.Abs() in config.NewProvider: %w", err))
	}

	envMap, err := godotenv.Read(envPath)
	if err != nil {
		panic(fmt.Errorf("unable to read env variables from env file, path: %v, err: %w", envPath, err))
	}

	for envName, envValue := range envMap {
		s.secrets[Secret(envName)] = secretValue(envValue)
	}

	return s
}

func (s *secretClientImpl) GetSecret(secret Secret) Value {
	raw, ok := s.secrets[secret]
	if !ok {
		panic(fmt.Errorf("unable to get secret value by key: %v", secret))
	}

	return secretValue(raw)
}

func (v secretValue) String() string {
	return string(v)
}

func (v secretValue) Int() int {
	i, _ := strconv.Atoi(string(v))
	return i
}

func (v secretValue) Bool() bool {
	b, _ := strconv.ParseBool(string(v))
	return b
}

func (v secretValue) Duration() time.Duration {
	d, _ := time.ParseDuration(string(v))
	return d
}
