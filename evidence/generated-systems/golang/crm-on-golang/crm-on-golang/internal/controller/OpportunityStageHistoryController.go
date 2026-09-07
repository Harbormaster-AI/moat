package controller

import (
    OpportunityStageHistoryDAO "crm-on-golang/internal/dao"
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to OpportunityStageHistoryDAO for database creation
//----------------------------------------------------------------------------
func CreateOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OpportunityStageHistory model
	//----------------------------------------------------------------------------
	data := model.OpportunityStageHistory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OpportunityStageHistory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory data access object to create
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.CreateOpportunityStageHistory( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to OpportunityStageHistoryDAO to find the relevant OpportunityStageHistory
//----------------------------------------------------------------------------
func GetOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OpportunityStageHistory data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.GetOpportunityStageHistory(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to OpportunityStageHistoryDAO for database read of all OpportunityStageHistorys
//----------------------------------------------------------------------------
func GetAllOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory data access object to get all
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.GetAllOpportunityStageHistory()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to OpportunityStageHistoryDAO for database save
//----------------------------------------------------------------------------
func UpdateOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty OpportunityStageHistory model
	//----------------------------------------------------------------------------
	var data = model.OpportunityStageHistory{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a OpportunityStageHistory model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.UpdateOpportunityStageHistory(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to OpportunityStageHistoryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the OpportunityStageHistory data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := OpportunityStageHistoryDAO.DeleteOpportunityStageHistory(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Opportunity on a OpportunityStageHistory
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOpportunityToOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityStageHistoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	opportunityId,_ := strconv.ParseUint( vars["opportunityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.AssignOpportunityToOpportunityStageHistory(opportunityStageHistoryId, opportunityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Opportunity on a OpportunityStageHistory
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOpportunityFromOpportunityStageHistory( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityStageHistoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.UnassignOpportunityFromOpportunityStageHistory(opportunityStageHistoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ChangedBy on a OpportunityStageHistory
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignChangedByToOpportunityStageHistory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityStageHistoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	changedById,_ := strconv.ParseUint( vars["changedById"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.AssignChangedByToOpportunityStageHistory(opportunityStageHistoryId, changedById)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ChangedBy on a OpportunityStageHistory
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignChangedByFromOpportunityStageHistory( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	opportunityStageHistoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the OpportunityStageHistory DAO
	//----------------------------------------------------------------------------
	requestResult := OpportunityStageHistoryDAO.UnassignChangedByFromOpportunityStageHistory(opportunityStageHistoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


