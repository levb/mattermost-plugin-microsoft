package store

import "fmt"

const (
	UpdateStatusPropertyName              = "update_status"
	GetConfirmationPropertyName           = "get_confirmation"
	ReceiveNotificationsDuringMeetingName = "receive_notifications_during_meetings"
	SubscribePropertyName                 = "subscribe"
	AutoRespondPropertyName               = "auto_respond"
	ReceiveUpcomingEventReminderName      = "receive_reminder"
)

func (s *pluginStore) SetProperty(userID, propertyName string, value bool) error {
	user, err := s.LoadUser(userID)
	if err != nil {
		return err
	}

	switch propertyName {
	case UpdateStatusPropertyName:
		user.Settings.UpdateStatus = value
		s.Tracker.TrackAutomaticStatusUpdate(userID, value, "flow")
	case GetConfirmationPropertyName:
		user.Settings.GetConfirmation = value
	case AutoRespondPropertyName:
		user.Settings.AutoRespond = value
	case ReceiveUpcomingEventReminderName:
		user.Settings.ReceiveReminders = value
	case ReceiveNotificationsDuringMeetingName:
		user.Settings.ReceiveNotificationsDuringMeeting = value
	default:
		return fmt.Errorf("property %s not found", propertyName)
	}

	err = s.StoreUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *pluginStore) SetPostID(userID, propertyName, postID string) error {
	user, err := s.LoadUser(userID)
	if err != nil {
		return err
	}

	if user.WelcomeFlowStatus.PostIDs == nil {
		user.WelcomeFlowStatus.PostIDs = make(map[string]string)
	}

	user.WelcomeFlowStatus.PostIDs[propertyName] = postID

	err = s.StoreUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *pluginStore) GetPostID(userID, propertyName string) (string, error) {
	user, err := s.LoadUser(userID)
	if err != nil {
		return "", err
	}

	return user.WelcomeFlowStatus.PostIDs[propertyName], nil
}

func (s *pluginStore) RemovePostID(userID, propertyName string) error {
	user, err := s.LoadUser(userID)
	if err != nil {
		return err
	}

	delete(user.WelcomeFlowStatus.PostIDs, propertyName)

	err = s.StoreUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *pluginStore) GetCurrentStep(userID string) (int, error) {
	user, err := s.LoadUser(userID)
	if err != nil {
		return 0, err
	}

	return user.WelcomeFlowStatus.Step, nil
}
func (s *pluginStore) SetCurrentStep(userID string, step int) error {
	user, err := s.LoadUser(userID)
	if err != nil {
		return err
	}

	user.WelcomeFlowStatus.Step = step

	err = s.StoreUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *pluginStore) DeleteCurrentStep(userID string) error {
	user, err := s.LoadUser(userID)
	if err != nil {
		return err
	}

	user.WelcomeFlowStatus.Step = 0

	err = s.StoreUser(user)
	if err != nil {
		return err
	}

	return nil
}
