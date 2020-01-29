// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package api

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/config"
	"github.com/mattermost/mattermost-plugin-mscalendar/server/mscalendar"
)

func (api *api) preprocessAction(w http.ResponseWriter, req *http.Request) (mscalendar.MSCalendar, *mscalendar.User, string, string) {
	mattermostUserID := req.Header.Get("Mattermost-User-ID")

	request := model.PostActionIntegrationRequestFromJson(req.Body)
	if request == nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return nil, nil, "", ""
	}

	eventID, ok := request.Context[config.EventIDKey].(string)
	if !ok {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return nil, nil, "", ""
	}
	option, _ := request.Context["selected_option"].(string)
	mscal := mscalendar.New(api.Env, mattermostUserID)

	return mscal, mscalendar.NewUser(mattermostUserID), eventID, option
}

func (api *api) postActionAccept(w http.ResponseWriter, req *http.Request) {
	mscalendar, user, eventID, _ := api.preprocessAction(w, req)
	if eventID == "" {
		return
	}
	err := mscalendar.AcceptEvent(user, eventID)
	if err != nil {
		api.Logger.Warnf(err.Error())
		http.Error(w, "Failed to accept event: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (api *api) postActionDecline(w http.ResponseWriter, req *http.Request) {
	mscalendar, user, eventID, _ := api.preprocessAction(w, req)
	if eventID == "" {
		return
	}
	err := mscalendar.DeclineEvent(user, eventID)
	if err != nil {
		http.Error(w, "Failed to decline event: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (api *api) postActionTentative(w http.ResponseWriter, req *http.Request) {
	mscalendar, user, eventID, _ := api.preprocessAction(w, req)
	if eventID == "" {
		return
	}
	err := mscalendar.TentativelyAcceptEvent(user, eventID)
	if err != nil {
		http.Error(w, "Failed to tentatively accept event: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func (api *api) postActionRespond(w http.ResponseWriter, req *http.Request) {
	mscalendar, user, eventID, option := api.preprocessAction(w, req)
	if eventID == "" {
		return
	}
	err := mscalendar.RespondToEvent(user, eventID, option)
	if err != nil {
		// h.LogWarn(err.Error())
		http.Error(w, "Failed to respond to event: "+err.Error(), http.StatusBadRequest)
		return
	}
}
