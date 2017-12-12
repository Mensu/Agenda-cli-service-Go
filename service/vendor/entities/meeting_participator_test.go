package entities_test

import (
	"entities"
	"testing"
)

func TestMeetingParticipator(t *testing.T) {
	t.Log("[meetingparticipatortest] adding meeting participator")
	title := "test_agenda_title"
	m1 := &entities.Meeting{
		ParticipatorsUsername: []string{"test_agenda_username_1"},
	}
	m2 := &entities.MeetingParticipator{
		Title:    title,
		Username: "test_agenda_username_1",
	}
	entities.MPServ.Add(m1)
	t.Log("[meetingparticipatortest] deleting meeting participator")
	entities.MPServ.Delete(m2)
}
