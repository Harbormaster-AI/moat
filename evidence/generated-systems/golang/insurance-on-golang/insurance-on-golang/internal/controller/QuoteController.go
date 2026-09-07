package controller

import (
    QuoteDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QuoteDAO for database creation
//----------------------------------------------------------------------------
func CreateQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Quote model
	//----------------------------------------------------------------------------
	data := model.Quote{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Quote model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Quote data access object to create
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.CreateQuote( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QuoteDAO to find the relevant Quote
//----------------------------------------------------------------------------
func GetQuote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Quote data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.GetQuote(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QuoteDAO for database read of all Quotes
//----------------------------------------------------------------------------
func GetAllQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Quote data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.GetAllQuote()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QuoteDAO for database save
//----------------------------------------------------------------------------
func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Quote model
	//----------------------------------------------------------------------------
	var data = model.Quote{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Quote model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Quote data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.UpdateQuote(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QuoteDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQuote(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Quote data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QuoteDAO.DeleteQuote(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Application on a Quote
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignApplicationToQuote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	applicationId,_ := strconv.ParseUint( vars["applicationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.AssignApplicationToQuote(quoteId, applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Application on a Quote
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignApplicationFromQuote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.UnassignApplicationFromQuote(quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Policy on a Quote
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPolicyToQuote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policyId,_ := strconv.ParseUint( vars["policyId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.AssignPolicyToQuote(quoteId, policyId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Policy on a Quote
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPolicyFromQuote( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.UnassignPolicyFromQuote(quoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more underwritingDecisionsIds as a UnderwritingDecisions to a Quote
	//----------------------------------------------------------------------------
func AddUnderwritingDecisionsToQuote(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	underwritingDecisionsIds,_ := vars["underwritingDecisionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.AddUnderwritingDecisionsToQuote(quoteId, underwritingDecisionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more underwritingDecisionsIds as a UnderwritingDecisions from a Quote
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveUnderwritingDecisionsFromQuote(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	quoteId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	underwritingDecisionsIds,_ := vars["underwritingDecisionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Quote DAO
	//----------------------------------------------------------------------------
	requestResult := QuoteDAO.RemoveUnderwritingDecisionsFromQuote(quoteId, underwritingDecisionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
