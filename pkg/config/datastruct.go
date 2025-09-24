package config

import (
	"time"
)

type (
	Key    string
	Secret string

	Provider interface {
		GetConfigClient() ConfigClient
		GetSecretClient() SecretClient
	}

	ConfigClient interface {
		GetValue(key Key) Value
	}

	SecretClient interface {
		GetSecret(secret Secret) Value
	}

	Value interface {
		Bool() bool
		Int() int
		String() string
		Duration() time.Duration
	}
)
