package cassandra

type Config struct {
	ConnectionString string `yaml:"connectionString"`
	ClientID         string `yaml:"clientId"`
	ClientSecret     string `yaml:"clientSecret"`
	Token            string `yaml:"token"`
	Keyspace string `yaml:"keyspace"`
}
