package controller

import (
    WorkAuthorizationDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to WorkAuthorizationDAO for database creation
//----------------------------------------------------------------------------
func CreateWorkAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkAuthorization model
	//----------------------------------------------------------------------------
	data := model.WorkAuthorization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkAuthorization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization data access object to create
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.CreateWorkAuthorization( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to WorkAuthorizationDAO to find the relevant WorkAuthorization
//----------------------------------------------------------------------------
func GetWorkAuthorization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkAuthorization data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.GetWorkAuthorization(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to WorkAuthorizationDAO for database read of all WorkAuthorizations
//----------------------------------------------------------------------------
func GetAllWorkAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization data access object to get all
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.GetAllWorkAuthorization()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to WorkAuthorizationDAO for database save
//----------------------------------------------------------------------------
func UpdateWorkAuthorization(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty WorkAuthorization model
	//----------------------------------------------------------------------------
	var data = model.WorkAuthorization{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a WorkAuthorization model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.UpdateWorkAuthorization(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to WorkAuthorizationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteWorkAuthorization(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the WorkAuthorization data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := WorkAuthorizationDAO.DeleteWorkAuthorization(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Employee on a WorkAuthorization
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToWorkAuthorization(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workAuthorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization DAO
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.AssignEmployeeToWorkAuthorization(workAuthorizationId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a WorkAuthorization
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromWorkAuthorization( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	workAuthorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization DAO
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.UnassignEmployeeFromWorkAuthorization(workAuthorizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more documentsIds as a Documents to a WorkAuthorization
	//----------------------------------------------------------------------------
func AddDocumentsToWorkAuthorization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workAuthorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization DAO
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.AddDocumentsToWorkAuthorization(workAuthorizationId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more documentsIds as a Documents from a WorkAuthorization
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDocumentsFromWorkAuthorization(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	workAuthorizationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	documentsIds,_ := vars["documentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the WorkAuthorization DAO
	//----------------------------------------------------------------------------
	requestResult := WorkAuthorizationDAO.RemoveDocumentsFromWorkAuthorization(workAuthorizationId, documentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
