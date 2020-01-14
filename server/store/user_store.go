// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package store

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/remote"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/utils/kvstore"
	"golang.org/x/oauth2"
)

type UserStore interface {
	LoadUser(mattermostUserId string) (*User, error)
	LoadMattermostUserId(remoteUserId string) (string, error)
	LoadUserIndex() (UserIndex, error)
	StoreUser(user *User) error
	DeleteUser(mattermostUserId string) error
}

type UserIndex []*UserShort

type UserShort struct {
	MattermostUserID string `json:"mm_id"`
	RemoteID         string `json:"remote_id"`
	Email            string `json:"email"`
}

type User struct {
	PluginVersion    string
	Remote           *remote.User
	MattermostUserID string
	OAuth2Token      *oauth2.Token
	Settings         Settings `json:"mattermostSettings,omitempty"`
}

type Settings struct {
	EventSubscriptionID string
}

func (settings Settings) String() string {
	sub := "no subscription"
	if settings.EventSubscriptionID != "" {
		sub = "subscription ID: " + settings.EventSubscriptionID
	}
	return fmt.Sprintf(" - %s", sub)
}

func (s *pluginStore) LoadUser(mattermostUserId string) (*User, error) {
	user := User{}
	err := kvstore.LoadJSON(s.userKV, mattermostUserId, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *pluginStore) LoadMattermostUserId(remoteUserId string) (string, error) {
	data, err := s.userKV.Load(remoteUserId)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *pluginStore) LoadUserIndex() (UserIndex, error) {
	users := UserIndex{}
	err := kvstore.LoadJSON(s.userIndexKV, "", &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *pluginStore) StoreUser(user *User) error {
	err := kvstore.StoreJSON(s.userKV, user.MattermostUserID, user)
	if err != nil {
		return err
	}

	err = s.mattermostUserIDKV.Store(user.Remote.ID, []byte(user.MattermostUserID))
	if err != nil {
		_ = s.userKV.Delete(user.MattermostUserID)
		return err
	}

	var userIndex []*UserShort
	err = kvstore.LoadJSON(s.userIndexKV, "", &userIndex)
	if err != nil {
		userIndex = []*UserShort{}
	}

	newUser := &UserShort{
		MattermostUserID: user.MattermostUserID,
		RemoteID:         user.Remote.ID,
		Email:            user.Remote.Mail,
	}

	found := false
	filtered := []*UserShort{}
	for _, u := range userIndex {
		if u.MattermostUserID == user.MattermostUserID && u.RemoteID == user.Remote.ID {
			found = true
			filtered = append(filtered, newUser)
		} else {
			filtered = append(filtered, u)
		}
	}

	if !found {
		filtered = append(filtered, newUser)
	}

	err = kvstore.StoreJSON(s.userIndexKV, "", &filtered)
	if err != nil {
		return err
	}

	return nil
}

func (s *pluginStore) DeleteUser(mattermostUserID string) error {
	u, err := s.LoadUser(mattermostUserID)
	if err != nil {
		return err
	}
	err = s.userKV.Delete(mattermostUserID)
	if err != nil {
		return err
	}
	err = s.mattermostUserIDKV.Delete(u.Remote.ID)
	if err != nil {
		return err
	}
	return nil
}

func (index UserIndex) ByMattermostID() map[string]*UserShort {
	result := map[string]*UserShort{}

	for _, u := range index {
		result[u.MattermostUserID] = u
	}

	return result
}

func (index UserIndex) ByRemoteID() map[string]*UserShort {
	result := map[string]*UserShort{}

	for _, u := range index {
		result[u.RemoteID] = u
	}

	return result
}

func (index UserIndex) ByEmail() map[string]*UserShort {
	result := map[string]*UserShort{}

	for _, u := range index {
		result[u.Email] = u
	}

	return result
}

func (index UserIndex) GetMattermostUserIDs() []string {
	result := []string{}

	for _, u := range index {
		result = append(result, u.MattermostUserID)
	}

	return result
}
