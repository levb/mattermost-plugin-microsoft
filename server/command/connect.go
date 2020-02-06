// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package command

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/config"
	"github.com/pkg/errors"
)

func (c *Command) connect(parameters ...string) (string, error) {
	ru, err := c.MSCalendar.GetRemoteUser(c.Args.UserId)
	if err == nil {
		return fmt.Sprintf("Your account is already connected to %s. Please run `/mscalendar disconnect`", ru.Mail), nil
	}

	out := fmt.Sprintf("[Click here to link your %s account.](%s/oauth2/connect)",
		config.ApplicationName,
		c.Config.PluginURL)
	return out, nil
}

func (c *Command) connectBot(parameters ...string) (string, error) {
	isAdmin, err := c.MSCalendar.IsAuthorizedAdmin(c.Args.UserId)
	if err != nil || !isAdmin {
		return "", errors.New("non-admin user attempting to connect bot account")
	}

	ru, err := c.MSCalendar.GetRemoteUser(c.Config.BotUserID)
	if err == nil {
		return fmt.Sprintf("Bot user already connected to %s. Please run `/mscalendar disconnect_bot`", ru.Mail), nil
	}

	out := fmt.Sprintf("[Click here to link the bot's %s account.](%s/oauth2/connect_bot)",
		config.ApplicationName,
		c.Config.PluginURL)
	return out, nil
}
