package config

type Bleve struct {
	Index     string `mapstructure:"index"`
	SetupPath string `mapstructure:"setup_path"`
	SyncCron  string `mapstructure:"sync_cron"`
}
