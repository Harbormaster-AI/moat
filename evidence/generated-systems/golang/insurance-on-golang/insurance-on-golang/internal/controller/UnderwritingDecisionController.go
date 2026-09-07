package controller

import (
    UnderwritingDecisionDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UnderwritingDecisionDAO for database creation
//----------------------------------------------------------------------------
func CreateUnderwritingDecision(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UnderwritingDecision model
	//----------------------------------------------------------------------------
	data := model.UnderwritingDecision{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UnderwritingDecision model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision data access object to create
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.CreateUnderwritingDecision( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UnderwritingDecisionDAO to find the relevant UnderwritingDecision
//----------------------------------------------------------------------------
func GetUnderwritingDecision(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UnderwritingDecision data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.GetUnderwritingDecision(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UnderwritingDecisionDAO for database read of all UnderwritingDecisions
//----------------------------------------------------------------------------
func GetAllUnderwritingDecision(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.GetAllUnderwritingDecision()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UnderwritingDecisionDAO for database save
//----------------------------------------------------------------------------
func UpdateUnderwritingDecision(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty UnderwritingDecision model
	//----------------------------------------------------------------------------
	var data = model.UnderwritingDecision{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a UnderwritingDecision model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.UpdateUnderwritingDecision(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UnderwritingDecisionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUnderwritingDecision(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the UnderwritingDecision data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UnderwritingDecisionDAO.DeleteUnderwritingDecision(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Quote on a UnderwritingDecision
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignQuoteToUnderwritingDecision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwritingDecisionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quoteId,_ := strconv.ParseUint( vars["quoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.AssignQuoteToUnderwritingDecision(underwritingDecisionId, quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Quote on a UnderwritingDecision
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignQuoteFromUnderwritingDecision( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwritingDecisionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.UnassignQuoteFromUnderwritingDecision(underwritingDecisionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Underwriter on a UnderwritingDecision
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignUnderwriterToUnderwritingDecision(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwritingDecisionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	underwriterId,_ := strconv.ParseUint( vars["underwriterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.AssignUnderwriterToUnderwritingDecision(underwritingDecisionId, underwriterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Underwriter on a UnderwritingDecision
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignUnderwriterFromUnderwritingDecision( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwritingDecisionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the UnderwritingDecision DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwritingDecisionDAO.UnassignUnderwriterFromUnderwritingDecision(underwritingDecisionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


