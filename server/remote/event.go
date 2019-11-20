// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package remote

type Event struct {
	ID                string               `json:"id"`
	Subject           string               `json:"subject,omitempty"`
	BodyPreview       string               `json:"bodyPreview,omitempty"`
	Importance        string               `json:"importance,omitempty"`
	IsAllDay          bool                 `json:"isAllDay,omitempty"`
	IsCancelled       bool                 `json:"isCancelled,omitempty"`
	IsOrganizer       bool                 `json:"isOrganizer,omitempty"`
	ResponseRequested bool                 `json:"responseRequested,omitempty"`
	ShowAs            string               `json:"showAs,omitempty"`
	Weblink           string               `json:"weblink,omitempty"`
	Start             *DateTime            `json:"start,omitempty"`
	End               *DateTime            `json:"end,omitempty"`
	Location          *EventLocation       `json:"location,omitempty"`
	ResponseStatus    *EventResponseStatus `json:"responseStatus,omitempty"`
	Attendees         []*EventAttendee     `json:"attendees,omitempty"`
	Organizer         *EventAttendee       `json:"organizer,omitempty"`
}

type EventResponseStatus struct {
	Response string `json:"response,omitempty"`
	Time     string `json:"time,omitempty"`
}

type EventLocation struct {
	DisplayName string `json:"displayName,omitempty"`
}

type EventAttendee struct {
	Type         string               `json:"type,omitempty"`
	Status       *EventResponseStatus `json:"status,omitempty"`
	EmailAddress *EmailAddress        `json:"emailAddress,omitempty"`
}
