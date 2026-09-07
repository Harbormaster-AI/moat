package controller

import (
    ThirdPartyDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ThirdPartyDAO for database creation
//----------------------------------------------------------------------------
func CreateThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdParty model
	//----------------------------------------------------------------------------
	data := model.ThirdParty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdParty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object to create
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.CreateThirdParty( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ThirdPartyDAO to find the relevant ThirdParty
//----------------------------------------------------------------------------
func GetThirdParty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdParty data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.GetThirdParty(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ThirdPartyDAO for database read of all ThirdPartys
//----------------------------------------------------------------------------
func GetAllThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.GetAllThirdParty()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ThirdPartyDAO for database save
//----------------------------------------------------------------------------
func UpdateThirdParty(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ThirdParty model
	//----------------------------------------------------------------------------
	var data = model.ThirdParty{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ThirdParty model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.UpdateThirdParty(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ThirdPartyDAO for database deletion
//----------------------------------------------------------------------------
func DeleteThirdParty(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ThirdParty data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ThirdPartyDAO.DeleteThirdParty(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a ThirdParty
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToThirdParty(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AssignOrganizationToThirdParty(thirdPartyId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a ThirdParty
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromThirdParty( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.UnassignOrganizationFromThirdParty(thirdPartyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more processingActivitiesIds as a ProcessingActivities to a ThirdParty
	//----------------------------------------------------------------------------
func AddProcessingActivitiesToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddProcessingActivitiesToThirdParty(thirdPartyId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more processingActivitiesIds as a ProcessingActivities from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	processingActivitiesIds,_ := vars["processingActivitiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveProcessingActivitiesFromThirdParty(thirdPartyId, processingActivitiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more assessmentsIds as a Assessments to a ThirdParty
	//----------------------------------------------------------------------------
func AddAssessmentsToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assessmentsIds,_ := vars["assessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddAssessmentsToThirdParty(thirdPartyId, assessmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assessmentsIds as a Assessments from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssessmentsFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assessmentsIds,_ := vars["assessmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveAssessmentsFromThirdParty(thirdPartyId, assessmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more contractsIds as a Contracts to a ThirdParty
	//----------------------------------------------------------------------------
func AddContractsToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddContractsToThirdParty(thirdPartyId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more contractsIds as a Contracts from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveContractsFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	contractsIds,_ := vars["contractsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveContractsFromThirdParty(thirdPartyId, contractsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more obligationsIds as a Obligations to a ThirdParty
	//----------------------------------------------------------------------------
func AddObligationsToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddObligationsToThirdParty(thirdPartyId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more obligationsIds as a Obligations from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveObligationsFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationsIds,_ := vars["obligationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveObligationsFromThirdParty(thirdPartyId, obligationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more dataBreachesIds as a DataBreaches to a ThirdParty
	//----------------------------------------------------------------------------
func AddDataBreachesToThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.AddDataBreachesToThirdParty(thirdPartyId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more dataBreachesIds as a DataBreaches from a ThirdParty
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDataBreachesFromThirdParty(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	thirdPartyId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	dataBreachesIds,_ := vars["dataBreachesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ThirdParty DAO
	//----------------------------------------------------------------------------
	requestResult := ThirdPartyDAO.RemoveDataBreachesFromThirdParty(thirdPartyId, dataBreachesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
