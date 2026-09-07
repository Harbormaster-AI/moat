package controller

import (
    DataBreachDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DataBreachDAO for database creation
//----------------------------------------------------------------------------
func CreateDataBreach(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataBreach model
	//----------------------------------------------------------------------------
	data := model.DataBreach{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataBreach model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach data access object to create
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.CreateDataBreach( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DataBreachDAO to find the relevant DataBreach
//----------------------------------------------------------------------------
func GetDataBreach(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataBreach data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.GetDataBreach(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DataBreachDAO for database read of all DataBreachs
//----------------------------------------------------------------------------
func GetAllDataBreach(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DataBreach data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.GetAllDataBreach()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DataBreachDAO for database save
//----------------------------------------------------------------------------
func UpdateDataBreach(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DataBreach model
	//----------------------------------------------------------------------------
	var data = model.DataBreach{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DataBreach model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.UpdateDataBreach(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DataBreachDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDataBreach(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DataBreach data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DataBreachDAO.DeleteDataBreach(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a DataBreach
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToDataBreach(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.AssignOrganizationToDataBreach(dataBreachId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a DataBreach
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromDataBreach( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.UnassignOrganizationFromDataBreach(dataBreachId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Matter on a DataBreach
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMatterToDataBreach(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	matterId,_ := strconv.ParseUint( vars["matterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.AssignMatterToDataBreach(dataBreachId, matterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Matter on a DataBreach
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMatterFromDataBreach( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.UnassignMatterFromDataBreach(dataBreachId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a DataBreach
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.AddProcessingActivitiesToDataBreach(dataBreachId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a DataBreach
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.RemoveProcessingActivitiesFromDataBreach(dataBreachId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataCategoriesIds as a DataCategories to a DataBreach
	//----------------------------------------------------------------------------
func AddDataCategoriesToDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.AddDataCategoriesToDataBreach(dataBreachId, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataCategoriesIds as a DataCategories from a DataBreach
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataCategoriesFromDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataCategoriesIds,_ := vars["dataCategoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.RemoveDataCategoriesFromDataBreach(dataBreachId, dataCategoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more thirdPartiesIds as a ThirdParties to a DataBreach
	//----------------------------------------------------------------------------
func AddThirdPartiesToDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.AddThirdPartiesToDataBreach(dataBreachId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more thirdPartiesIds as a ThirdParties from a DataBreach
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveThirdPartiesFromDataBreach(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dataBreachId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	thirdPartiesIds,_ := vars["thirdPartiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DataBreach DAO
	//----------------------------------------------------------------------------
	requestResult := DataBreachDAO.RemoveThirdPartiesFromDataBreach(dataBreachId, thirdPartiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
