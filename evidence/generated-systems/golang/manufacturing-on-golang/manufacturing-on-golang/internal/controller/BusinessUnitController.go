package controller

import (
    BusinessUnitDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
	// assigns a Enterprise on a BusinessUnit
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEnterpriseToBusinessUnit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterpriseId,_ := strconv.ParseUint( vars["enterpriseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.AssignEnterpriseToBusinessUnit(businessUnitId, enterpriseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Enterprise on a BusinessUnit
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEnterpriseFromBusinessUnit( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.UnassignEnterpriseFromBusinessUnit(businessUnitId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more itemsIds as a Items to a BusinessUnit
	//----------------------------------------------------------------------------
func AddItemsToBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.AddItemsToBusinessUnit(businessUnitId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more itemsIds as a Items from a BusinessUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveItemsFromBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemsIds,_ := vars["itemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.RemoveItemsFromBusinessUnit(businessUnitId, itemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more plantsIds as a Plants to a BusinessUnit
	//----------------------------------------------------------------------------
func AddPlantsToBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.AddPlantsToBusinessUnit(businessUnitId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more plantsIds as a Plants from a BusinessUnit
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePlantsFromBusinessUnit(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	businessUnitId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantsIds,_ := vars["plantsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BusinessUnit DAO
	//----------------------------------------------------------------------------
	requestResult := BusinessUnitDAO.RemovePlantsFromBusinessUnit(businessUnitId, plantsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
