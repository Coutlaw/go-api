package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-api/models"
	u "go-api/utils"
	"net/http"
	"strconv"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContactById = func(w http.ResponseWriter, r *http.Request) {

	// Fetch the inline params
	vars := mux.Vars(r)
	contactIdParam := vars["contactId"]

	// Convert inline param to uint
	contactId, err := strconv.ParseUint(contactIdParam, 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Error with contactId param, could not be converted to uint"))
	}

	// pull User Id from context
	userId := r.Context().Value("user").(uint)

	data := models.GetContact(uint(contactId), userId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
