package importer

type Config struct {
	ImportAtStart   bool `yaml:"importAtStart"`
	ImportFrequency int  `yaml:"importFrequency"`
}
