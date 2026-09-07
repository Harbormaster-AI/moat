package controller

import (
    InsuredObjectDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to InsuredObjectDAO for database creation
//----------------------------------------------------------------------------
func CreateInsuredObject(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsuredObject model
	//----------------------------------------------------------------------------
	data := model.InsuredObject{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsuredObject model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject data access object to create
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.CreateInsuredObject( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to InsuredObjectDAO to find the relevant InsuredObject
//----------------------------------------------------------------------------
func GetInsuredObject(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsuredObject data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.GetInsuredObject(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to InsuredObjectDAO for database read of all InsuredObjects
//----------------------------------------------------------------------------
func GetAllInsuredObject(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject data access object to get all
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.GetAllInsuredObject()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to InsuredObjectDAO for database save
//----------------------------------------------------------------------------
func UpdateInsuredObject(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty InsuredObject model
	//----------------------------------------------------------------------------
	var data = model.InsuredObject{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a InsuredObject model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.UpdateInsuredObject(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to InsuredObjectDAO for database deletion
//----------------------------------------------------------------------------
func DeleteInsuredObject(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the InsuredObject data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := InsuredObjectDAO.DeleteInsuredObject(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Policy on a InsuredObject
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToInsuredObject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insuredObjectId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject DAO
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.AssignPolicyToInsuredObject(insuredObjectId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a InsuredObject
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromInsuredObject( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	insuredObjectId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject DAO
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.UnassignPolicyFromInsuredObject(insuredObjectId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more coveragesIds as a Coverages to a InsuredObject
	//----------------------------------------------------------------------------
func AddCoveragesToInsuredObject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insuredObjectId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject DAO
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.AddCoveragesToInsuredObject(insuredObjectId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more coveragesIds as a Coverages from a InsuredObject
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCoveragesFromInsuredObject(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	insuredObjectId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	coveragesIds,_ := vars["coveragesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the InsuredObject DAO
	//----------------------------------------------------------------------------
	requestResult := InsuredObjectDAO.RemoveCoveragesFromInsuredObject(insuredObjectId, coveragesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
