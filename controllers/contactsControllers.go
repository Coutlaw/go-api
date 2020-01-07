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
		http.Error(w, "Error while decoding request body, your JSON is probably malformed", http.StatusBadRequest)
		return
	}

	contact.UserId = user
	resp := contact.Create()
	if resp["success"].(bool) != true {
		http.Error(w, resp["message"].(string), http.StatusBadRequest)
		return
	}
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
		http.Error(w, "Error with contactId param, could not be converted to uint", http.StatusBadRequest)
		return
	}

	// pull User Id from context
	userId := r.Context().Value("user").(uint)

	data := models.GetContact(uint(contactId), userId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
