package controller

import (
    ApplicationDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ApplicationDAO for database creation
//----------------------------------------------------------------------------
func CreateApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Application model
	//----------------------------------------------------------------------------
	data := model.Application{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Application model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Application data access object to create
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.CreateApplication( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ApplicationDAO to find the relevant Application
//----------------------------------------------------------------------------
func GetApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Application data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.GetApplication(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ApplicationDAO for database read of all Applications
//----------------------------------------------------------------------------
func GetAllApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Application data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.GetAllApplication()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ApplicationDAO for database save
//----------------------------------------------------------------------------
func UpdateApplication(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Application model
	//----------------------------------------------------------------------------
	var data = model.Application{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Application model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Application data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.UpdateApplication(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ApplicationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Application data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ApplicationDAO.DeleteApplication(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Customer on a Application
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCustomerToApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	customerId,_ := strconv.ParseUint( vars["customerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.AssignCustomerToApplication(applicationId, customerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Customer on a Application
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCustomerFromApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.UnassignCustomerFromApplication(applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Product on a Application
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProductToApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productId,_ := strconv.ParseUint( vars["productId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.AssignProductToApplication(applicationId, productId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Product on a Application
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProductFromApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.UnassignProductFromApplication(applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Distributor on a Application
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignDistributorToApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	distributorId,_ := strconv.ParseUint( vars["distributorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.AssignDistributorToApplication(applicationId, distributorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Distributor on a Application
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignDistributorFromApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.UnassignDistributorFromApplication(applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a SelectedQuote on a Application
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignSelectedQuoteToApplication(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	selectedQuoteId,_ := strconv.ParseUint( vars["selectedQuoteId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.AssignSelectedQuoteToApplication(applicationId, selectedQuoteId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a SelectedQuote on a Application
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignSelectedQuoteFromApplication( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.UnassignSelectedQuoteFromApplication(applicationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a Application
	//----------------------------------------------------------------------------
func AddQuotesToApplication(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.AddQuotesToApplication(applicationId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a Application
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromApplication(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	applicationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Application DAO
	//----------------------------------------------------------------------------
	requestResult := ApplicationDAO.RemoveQuotesFromApplication(applicationId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
