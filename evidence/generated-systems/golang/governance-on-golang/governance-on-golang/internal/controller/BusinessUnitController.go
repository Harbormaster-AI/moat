package controller

import (
    BusinessUnitDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BusinessUnitDAO for database creation
//----------------------------------------------------------------------------
func CreateBusinessUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BusinessUnit model
	//----------------------------------------------------------------------------
	data := model.BusinessUnit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BusinessUnit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit data access object to create
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.CreateBusinessUnit( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BusinessUnitDAO to find the relevant BusinessUnit
//----------------------------------------------------------------------------
func GetBusinessUnit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BusinessUnit data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.GetBusinessUnit(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BusinessUnitDAO for database read of all BusinessUnits
//----------------------------------------------------------------------------
func GetAllBusinessUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.GetAllBusinessUnit()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BusinessUnitDAO for database save
//----------------------------------------------------------------------------
func UpdateBusinessUnit(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BusinessUnit model
	//----------------------------------------------------------------------------
	var data = model.BusinessUnit{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BusinessUnit model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.UpdateBusinessUnit(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BusinessUnitDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBusinessUnit(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BusinessUnit data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BusinessUnitDAO.DeleteBusinessUnit(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a BusinessUnit
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToBusinessUnit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.AssignOrganizationToBusinessUnit(businessUnitId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a BusinessUnit
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromBusinessUnit( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.UnassignOrganizationFromBusinessUnit(businessUnitId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more auditsIds as a Audits to a BusinessUnit
	//----------------------------------------------------------------------------
func AddAuditsToBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	auditsIds,_ := vars["auditsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.AddAuditsToBusinessUnit(businessUnitId, auditsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more auditsIds as a Audits from a BusinessUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAuditsFromBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	auditsIds,_ := vars["auditsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.RemoveAuditsFromBusinessUnit(businessUnitId, auditsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
