package entities_test

import (
	"entities"
	"testing"
)

func TestMeeting(t *testing.T) {
	t.Log("[meetingtest] adding meeting")
	title := "test_agenda_title"
	m := &entities.Meeting{
		Title:     title,
		Speecher:  entities.User{Username: "test_agenda_username"},
		StartTime: "2017-01-01/12:00:00",
		EndTime:   "2017-01-01/12:00:00",
	}
	entities.MeetingServ.Add(m)
	t.Log("[meetingtest] finding meeting")
	if entities.MeetingServ.FindByTitle(title) == nil {
		t.Fatalf("could not find meeting '%s' that was just added", title)
	}
	t.Log("[meetingtest] deleting meeting")
	entities.MeetingServ.Delete(m)
}
