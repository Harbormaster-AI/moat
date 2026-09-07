package controller

import (
    DocumentDAO "hr-on-golang/internal/dao"
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DocumentDAO for database creation
//----------------------------------------------------------------------------
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Document model
	//----------------------------------------------------------------------------
	data := model.Document{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Document model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Document data access object to create
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.CreateDocument( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DocumentDAO to find the relevant Document
//----------------------------------------------------------------------------
func GetDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Document data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.GetDocument(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DocumentDAO for database read of all Documents
//----------------------------------------------------------------------------
func GetAllDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Document data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.GetAllDocument()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DocumentDAO for database save
//----------------------------------------------------------------------------
func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Document model
	//----------------------------------------------------------------------------
	var data = model.Document{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Document model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Document data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UpdateDocument(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DocumentDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Document data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DocumentDAO.DeleteDocument(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Candidate on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCandidateToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	candidateId,_ := strconv.ParseUint( vars["candidateId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignCandidateToDocument(documentId, candidateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Candidate on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCandidateFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignCandidateFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Employee on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEmployeeToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	employeeId,_ := strconv.ParseUint( vars["employeeId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignEmployeeToDocument(documentId, employeeId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Employee on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEmployeeFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignEmployeeFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


