package mstdn

type Config struct {
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	AccessToken  string `yaml:"accessToken"`
}
