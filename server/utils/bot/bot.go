// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package bot

import (
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type Bot interface {
	Poster
	Logger
	Admin

	Ensure(stored *model.Bot, iconPath string) error
	WithConfig(BotConfig) Bot
	MattermostUserID() string
}

type bot struct {
	BotConfig
	pluginAPI        plugin.API
	pluginHelpers    plugin.Helpers
	mattermostUserID string
	displayName      string
	logContext       LogContext
}

func New(api plugin.API, helpers plugin.Helpers) Bot {
	return &bot{
		pluginAPI:     api,
		pluginHelpers: helpers,
	}
}

func (bot *bot) Ensure(stored *model.Bot, iconPath string) error {
	if bot.mattermostUserID != "" {
		// Already done
		return nil
	}

	botUserID, err := bot.pluginHelpers.EnsureBot(stored, plugin.ProfileImagePath(iconPath))
	if err != nil {
		return errors.Wrap(err, "failed to ensure bot account")
	}
	bot.mattermostUserID = botUserID
	bot.displayName = stored.DisplayName
	return nil
}

func (bot *bot) WithConfig(conf BotConfig) Bot {
	newbot := *bot
	newbot.BotConfig = conf
	return &newbot
}

func (bot *bot) MattermostUserID() string {
	return bot.mattermostUserID
}

func (bot *bot) String() string {
	return bot.displayName
}
