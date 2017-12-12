package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"entities"

	"github.com/unrolled/render"
)

func validateNewMeeting(meeting *entities.Meeting) error {
	if len(meeting.Title) == 0 {
		return fmt.Errorf("Title should not be empty")
	}
	if len(meeting.ParticipatorsUsername) == 0 {
		return fmt.Errorf("Meeting must have participator")
	}

	return nil
}

func addMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var meeting entities.Meeting
		if err := json.NewDecoder(r.Body).Decode(&meeting); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		var badData struct {
			Msg string `json:"msg"`
		}

		if err := validateNewMeeting(&meeting); err != nil {
			badData.Msg = err.Error()
			formatter.JSON(w, http.StatusBadRequest, badData)
			return
		}

		user := checkIsLogin(r)
		meeting.SpeecherUsername = user.Username

		entities.MeetingServ.Add(&meeting)
		entities.MPServ.Add(&meeting)

		formatter.JSON(w, http.StatusCreated, struct{}{})
	}
}

func addMPHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody struct {
			ParticipatorsUsername string `json:"participators"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		title := mux.Vars(r)["title"]

		meeting := entities.Meeting{
			Title: title,
			ParticipatorsUsername: []string{reqBody.ParticipatorsUsername},
		}
		entities.MPServ.Add(&meeting)
		formatter.JSON(w, http.StatusOK, struct{}{})
	}
}

func deleteMPHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody struct {
			ParticipatorsUsername string `json:"participators"`
		}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		title := mux.Vars(r)["title"]

		meeting := entities.MeetingParticipator{
			Title:    title,
			Username: reqBody.ParticipatorsUsername,
		}
		entities.MPServ.Delete(&meeting)
		formatter.JSON(w, http.StatusOK, struct{}{})
	}
}

func findAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.JSON(w, http.StatusOK, entities.MeetingServ.FindAll())
	}
}
