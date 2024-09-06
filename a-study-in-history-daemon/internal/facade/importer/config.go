package importer

type Config struct {
	ImportAtStart      bool `yaml:"importAtStart"`
	ImportConcurrently bool `yaml:"importConcurrently"`
	ImportFrequency    int  `yaml:"importFrequency"`
}
