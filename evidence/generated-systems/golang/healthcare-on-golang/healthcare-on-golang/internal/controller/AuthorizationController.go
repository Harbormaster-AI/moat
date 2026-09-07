package controller

import (
    AuthorizationDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to AuthorizationDAO for database creation
//----------------------------------------------------------------------------
func CreateAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Authorization model
	//----------------------------------------------------------------------------
	data := model.Authorization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Authorization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization data access object to create
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.CreateAuthorization( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to AuthorizationDAO to find the relevant Authorization
//----------------------------------------------------------------------------
func GetAuthorization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Authorization data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.GetAuthorization(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to AuthorizationDAO for database read of all Authorizations
//----------------------------------------------------------------------------
func GetAllAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Authorization data access object to get all
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.GetAllAuthorization()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to AuthorizationDAO for database save
//----------------------------------------------------------------------------
func UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Authorization model
	//----------------------------------------------------------------------------
	var data = model.Authorization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Authorization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.UpdateAuthorization(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to AuthorizationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteAuthorization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Authorization data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := AuthorizationDAO.DeleteAuthorization(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Coverage on a Authorization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCoverageToAuthorization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	authorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coverageId,_ := strconv.ParseUint( vars["coverageId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization DAO
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.AssignCoverageToAuthorization(authorizationId, coverageId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Coverage on a Authorization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCoverageFromAuthorization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	authorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization DAO
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.UnassignCoverageFromAuthorization(authorizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Order on a Authorization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrderToAuthorization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	authorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	orderId,_ := strconv.ParseUint( vars["orderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization DAO
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.AssignOrderToAuthorization(authorizationId, orderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Order on a Authorization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrderFromAuthorization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	authorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Authorization DAO
	//----------------------------------------------------------------------------
	requestResult := AuthorizationDAO.UnassignOrderFromAuthorization(authorizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


