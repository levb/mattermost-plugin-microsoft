package config

import "github.com/mattermost/mattermost-plugin-mscalendar/server/utils/bot"

// StoredConfig represents the data stored in and managed with the Mattermost
// config.
type StoredConfig struct {
	OAuth2Authority    string
	OAuth2ClientID     string
	OAuth2ClientSecret string

	EnableStatusSync   bool
	EnableDailySummary bool

	GoogleDomainVerifyKey string

	bot.Config
}

// Config represents the the metadata handed to all request runners (command,
// http).
type Config struct {
	PluginID               string
	BuildDate              string
	BuildHash              string
	BuildHashShort         string
	MattermostSiteHostname string
	MattermostSiteURL      string
	PluginURL              string
	PluginURLPath          string
	PluginVersion          string
	StoredConfig
}

func (c *Config) GetNotificationURL() string {
	return c.PluginURL + FullPathEventNotification
}
