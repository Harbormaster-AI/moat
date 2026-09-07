package controller

import (
    DocumentDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
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
	// assigns a Policy on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignPolicyToDocument(documentId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignPolicyFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Claim on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignClaimToDocument(documentId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignClaimFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Application on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApplicationToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationId,_ := strconv.ParseUint( vars["applicationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignApplicationToDocument(documentId, applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Application on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApplicationFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignApplicationFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Document
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToDocument(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.AssignCustomerToDocument(documentId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Document
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromDocument( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	documentId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Document DAO
	//----------------------------------------------------------------------------
	requestResult := DocumentDAO.UnassignCustomerFromDocument(documentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


