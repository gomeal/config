package config

var _ Provider = &providerImpl{}

type providerImpl struct {
	c *configClientImpl
	s *secretClientImpl
}

func NewProvider(configPath string) *providerImpl {
	p := &providerImpl{
		c: newConfigClient(configPath),
		s: newSecretClient(),
	}

	return p
}

func (p *providerImpl) GetConfigClient() ConfigClient {
	return p.c
}

func (p *providerImpl) GetSecretClient() SecretClient {
	return p.s
}
