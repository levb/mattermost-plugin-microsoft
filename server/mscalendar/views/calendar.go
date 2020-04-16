package views

import (
	"fmt"
	"sort"
	"time"

	"github.com/mattermost/mattermost-plugin-mscalendar/server/remote"
)

func RenderCalendarView(events []*remote.Event, timeZone string) (string, error) {
	if len(events) == 0 {
		return "You have no upcoming events.", nil
	}

	if timeZone != "" {
		for _, e := range events {
			e.Start = e.Start.In(timeZone)
			e.End = e.End.In(timeZone)
		}
	}

	resp := "Times are shown in " + events[0].Start.TimeZone + "\n"
	for _, group := range groupEventsByDate(events) {
		resp += group[0].Start.Time().Format("Monday January 02") + "\n"
		for _, e := range group {
			resp += fmt.Sprintf("* %s\n", renderEvent(e))
		}
	}

	return resp, nil
}

func renderEvent(event *remote.Event) string {
	start := event.Start.Time().Format(time.Kitchen)
	end := event.End.Time().Format(time.Kitchen)

	return fmt.Sprintf("%s - %s `%s`", start, end, event.Subject)
}

func groupEventsByDate(events []*remote.Event) [][]*remote.Event {
	groups := map[string][]*remote.Event{}

	for _, event := range events {
		date := event.Start.Time().Format("2006-01-02")
		_, ok := groups[date]
		if !ok {
			groups[date] = []*remote.Event{}
		}

		groups[date] = append(groups[date], event)
	}

	days := []string{}
	for k := range groups {
		days = append(days, k)
	}
	sort.Strings(days)

	result := [][]*remote.Event{}
	for _, day := range days {
		group := groups[day]
		result = append(result, group)
	}
	return result
}
