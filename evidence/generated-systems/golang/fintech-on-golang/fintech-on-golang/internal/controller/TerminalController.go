package controller

import (
    TerminalDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TerminalDAO for database creation
//----------------------------------------------------------------------------
func CreateTerminal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Terminal model
	//----------------------------------------------------------------------------
	data := model.Terminal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Terminal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Terminal data access object to create
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.CreateTerminal( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TerminalDAO to find the relevant Terminal
//----------------------------------------------------------------------------
func GetTerminal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Terminal data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.GetTerminal(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TerminalDAO for database read of all Terminals
//----------------------------------------------------------------------------
func GetAllTerminal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Terminal data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.GetAllTerminal()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TerminalDAO for database save
//----------------------------------------------------------------------------
func UpdateTerminal(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Terminal model
	//----------------------------------------------------------------------------
	var data = model.Terminal{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Terminal model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Terminal data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.UpdateTerminal(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TerminalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTerminal(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Terminal data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TerminalDAO.DeleteTerminal(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Merchant on a Terminal
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignMerchantToTerminal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	terminalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	merchantId,_ := strconv.ParseUint( vars["merchantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Terminal DAO
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.AssignMerchantToTerminal(terminalId, merchantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Merchant on a Terminal
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignMerchantFromTerminal( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	terminalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Terminal DAO
	//----------------------------------------------------------------------------
	requestResult := TerminalDAO.UnassignMerchantFromTerminal(terminalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


