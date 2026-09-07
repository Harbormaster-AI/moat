package controller

import (
    UserDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UserDAO for database creation
//----------------------------------------------------------------------------
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty User model
	//----------------------------------------------------------------------------
	data := model.User{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a User model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the User data access object to create
	//----------------------------------------------------------------------------
	requestResult := UserDAO.CreateUser( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UserDAO to find the relevant User
//----------------------------------------------------------------------------
func GetUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the User data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UserDAO.GetUser(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UserDAO for database read of all Users
//----------------------------------------------------------------------------
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the User data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UserDAO.GetAllUser()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UserDAO for database save
//----------------------------------------------------------------------------
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty User model
	//----------------------------------------------------------------------------
	var data = model.User{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a User model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the User data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UserDAO.UpdateUser(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UserDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the User data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UserDAO.DeleteUser(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Agency on a User
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAgencyToUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	agencyId,_ := strconv.ParseUint( vars["agencyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AssignAgencyToUser(userId, agencyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Agency on a User
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAgencyFromUser( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.UnassignAgencyFromUser(userId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more teamsIds as a Teams to a User
	//----------------------------------------------------------------------------
func AddTeamsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddTeamsToUser(userId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more teamsIds as a Teams from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTeamsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	teamsIds,_ := vars["teamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveTeamsFromUser(userId, teamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more adAccountsIds as a AdAccounts to a User
	//----------------------------------------------------------------------------
func AddAdAccountsToUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.AddAdAccountsToUser(userId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adAccountsIds as a AdAccounts from a User
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdAccountsFromUser(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	userId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the User DAO
	//----------------------------------------------------------------------------
	requestResult := UserDAO.RemoveAdAccountsFromUser(userId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
