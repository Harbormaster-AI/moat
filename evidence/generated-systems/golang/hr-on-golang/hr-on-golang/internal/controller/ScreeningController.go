package controller

import (
    ScreeningDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ScreeningDAO for database creation
//----------------------------------------------------------------------------
func CreateScreening(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Screening model
	//----------------------------------------------------------------------------
	data := model.Screening{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Screening model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Screening data access object to create
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.CreateScreening( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ScreeningDAO to find the relevant Screening
//----------------------------------------------------------------------------
func GetScreening(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]
	
	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	//----------------------------------------------------------------------------
	// Delegate to the Screening data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.GetScreening(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ScreeningDAO for database read of all Screenings
//----------------------------------------------------------------------------
func GetAllScreening(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Screening data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.GetAllScreening()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ScreeningDAO for database save
//----------------------------------------------------------------------------
func UpdateScreening(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Screening model
	//----------------------------------------------------------------------------
	var data = model.Screening{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Screening model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Screening data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.UpdateScreening(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ScreeningDAO for database deletion
//----------------------------------------------------------------------------
func DeleteScreening(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Retrieve the parameter from the request using hte mux
	//----------------------------------------------------------------------------
	vars := mux.Vars(r)
	
	//----------------------------------------------------------------------------
	// Locate the value for the ID key
	//----------------------------------------------------------------------------	
	id := vars["id"]

	//----------------------------------------------------------------------------
	// Parse the value into an integer if provided as such
	//----------------------------------------------------------------------------	
	ID, err:= strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	//----------------------------------------------------------------------------
	// Delegate to the Screening data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ScreeningDAO.DeleteScreening(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Application on a Screening
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApplicationToScreening(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	screeningId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationId,_ := strconv.ParseUint( vars["applicationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Screening DAO
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.AssignApplicationToScreening(screeningId, applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Application on a Screening
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApplicationFromScreening( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	screeningId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Screening DAO
	//----------------------------------------------------------------------------
	requestResult := ScreeningDAO.UnassignApplicationFromScreening(screeningId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


