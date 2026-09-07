package controller

import (
    TeamDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TeamDAO for database creation
//----------------------------------------------------------------------------
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Team model
	//----------------------------------------------------------------------------
	data := model.Team{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Team model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Team data access object to create
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.CreateTeam( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TeamDAO to find the relevant Team
//----------------------------------------------------------------------------
func GetTeam(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Team data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.GetTeam(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TeamDAO for database read of all Teams
//----------------------------------------------------------------------------
func GetAllTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Team data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.GetAllTeam()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TeamDAO for database save
//----------------------------------------------------------------------------
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Team model
	//----------------------------------------------------------------------------
	var data = model.Team{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Team model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Team data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.UpdateTeam(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TeamDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Team data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TeamDAO.DeleteTeam(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Agency on a Team
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAgencyToTeam(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agencyId,_ := strconv.ParseUint( vars["agencyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.AssignAgencyToTeam(teamId, agencyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Agency on a Team
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAgencyFromTeam( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.UnassignAgencyFromTeam(teamId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more usersIds as a Users to a Team
	//----------------------------------------------------------------------------
func AddUsersToTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.AddUsersToTeam(teamId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more usersIds as a Users from a Team
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUsersFromTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	usersIds,_ := vars["usersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.RemoveUsersFromTeam(teamId, usersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more adAccountsIds as a AdAccounts to a Team
	//----------------------------------------------------------------------------
func AddAdAccountsToTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.AddAdAccountsToTeam(teamId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adAccountsIds as a AdAccounts from a Team
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdAccountsFromTeam(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	teamId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Team DAO
	//----------------------------------------------------------------------------
	requestResult := TeamDAO.RemoveAdAccountsFromTeam(teamId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
