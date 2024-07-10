package config

var Load config

type config struct {
	MBAccessKey   string `yaml:"mb_access_key"`
	MBWorkspaceID string `yaml:"mb_workspace_id"`
	MBChannelID   string `yaml:"mb_channel_id"`
}
