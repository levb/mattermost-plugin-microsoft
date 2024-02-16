package main

import (
	mattermostplugin "github.com/mattermost/mattermost/server/public/plugin"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/config"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/mscalendar"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/plugin"
)

var BuildHash string
var BuildHashShort string
var BuildDate string

func main() {
	mattermostplugin.ClientMain(
		plugin.NewWithEnv(
			mscalendar.Env{
				Config: &config.Config{
					PluginID:       manifest.Id,
					PluginVersion:  manifest.Version,
					BuildHash:      BuildHash,
					BuildHashShort: BuildHashShort,
					BuildDate:      BuildDate,
				},
				Dependencies: &mscalendar.Dependencies{},
			}))
}
