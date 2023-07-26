package main

import (
	mattermostplugin "github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/config"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/mscalendar"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/plugin"
)

var BuildHash string
var BuildHashShort string
var BuildDate string
var CalendarProvider string
var CalendarProviderDisplayName string

func main() {
	config.Provider = config.ProviderConfig{
		Name:               CalendarProvider,
		DisplayName:        CalendarProviderDisplayName,
		Repository:         "",
		CommandTrigger:     CalendarProvider,
		TelemetryShortName: CalendarProvider,
		BotUsername:        CalendarProvider,
		BotDisplayName:     CalendarProviderDisplayName,
	}

	mattermostplugin.ClientMain(
		plugin.NewWithEnv(
			mscalendar.Env{
				Config: &config.Config{
					PluginID:       manifest.ID,
					PluginVersion:  manifest.Version,
					BuildHash:      BuildHash,
					BuildHashShort: BuildHashShort,
					BuildDate:      BuildDate,
					Provider:       config.Provider,
				},
				Dependencies: &mscalendar.Dependencies{},
			}))
}
