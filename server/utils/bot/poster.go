// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package bot

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

type Poster interface {
	// DM posts a simple Direct Message to the specified user
	DM(mattermostUserID, format string, args ...interface{}) error

	// DMWithAttachments posts a Direct Message that contains Slack attachments.
	// Often used to include post actions.
	DMWithAttachments(mattermostUserID string, attachments ...*model.SlackAttachment) (string, error)

	// Ephemeral sends an ephemeral message to a user
	Ephemeral(mattermostUserID, channelID, format string, args ...interface{})

	// DeletePost deletes a single post
	DeletePost(postID string) error

	// DMUpdatePost substitute one post with another
	UpdatePost(post *model.Post) error
}

// DM posts a simple Direct Message to the specified user
func (bot *bot) DM(mattermostUserID, format string, args ...interface{}) error {
	_, err := bot.dm(mattermostUserID, &model.Post{
		Message: fmt.Sprintf(format, args...),
	})
	if err != nil {
		return err
	}
	return nil
}

// DMWithAttachments posts a Direct Message that contains Slack attachments.
// Often used to include post actions.
func (bot *bot) DMWithAttachments(mattermostUserID string, attachments ...*model.SlackAttachment) (string, error) {
	post := model.Post{}
	model.ParseSlackAttachment(&post, attachments)
	return bot.dm(mattermostUserID, &post)
}

func (bot *bot) dm(mattermostUserID string, post *model.Post) (string, error) {
	channel, err := bot.pluginAPI.GetDirectChannel(mattermostUserID, bot.mattermostUserID)
	if err != nil {
		bot.pluginAPI.LogInfo("Couldn't get bot's DM channel", "user_id", mattermostUserID)
		return "", err
	}
	post.ChannelId = channel.Id
	post.UserId = bot.mattermostUserID
	sentPost, err := bot.pluginAPI.CreatePost(post)
	if err != nil {
		return "", err
	}
	return sentPost.Id, nil
}

// DM posts a simple Direct Message to the specified user
func (bot *bot) dmAdmins(format string, args ...interface{}) error {
	for _, id := range strings.Split(bot.AdminUserIDs, ",") {
		_, err := bot.dm(id, &model.Post{
			Message: fmt.Sprintf(format, args...),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Ephemeral sends an ephemeral message to a user
func (bot *bot) Ephemeral(userId, channelId, format string, args ...interface{}) {
	post := &model.Post{
		UserId:    bot.mattermostUserID,
		ChannelId: channelId,
		Message:   fmt.Sprintf(format, args...),
	}
	_ = bot.pluginAPI.SendEphemeralPost(userId, post)
}

func (bot *bot) DeletePost(postID string) error {
	appErr := bot.pluginAPI.DeletePost(postID)
	if appErr != nil {
		return appErr
	}
	return nil
}

func (bot *bot) UpdatePost(post *model.Post) error {
	_, appErr := bot.pluginAPI.UpdatePost(post)
	if appErr != nil {
		return appErr
	}
	return nil
}
