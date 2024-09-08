package twitter

type Config struct {
	ConsumerKey    string `yaml:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret"`
	Token          string `yaml:"token"`
	TokenSecret    string `yaml:"tokenSecret"`
}
