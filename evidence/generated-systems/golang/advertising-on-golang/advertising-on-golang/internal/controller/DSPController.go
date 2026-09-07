package controller

import (
    DSPDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to DSPDAO for database creation
//----------------------------------------------------------------------------
func CreateDSP(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DSP model
	//----------------------------------------------------------------------------
	data := model.DSP{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DSP model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DSP data access object to create
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.CreateDSP( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to DSPDAO to find the relevant DSP
//----------------------------------------------------------------------------
func GetDSP(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DSP data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.GetDSP(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to DSPDAO for database read of all DSPs
//----------------------------------------------------------------------------
func GetAllDSP(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the DSP data access object to get all
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.GetAllDSP()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to DSPDAO for database save
//----------------------------------------------------------------------------
func UpdateDSP(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty DSP model
	//----------------------------------------------------------------------------
	var data = model.DSP{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a DSP model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the DSP data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.UpdateDSP(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to DSPDAO for database deletion
//----------------------------------------------------------------------------
func DeleteDSP(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the DSP data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := DSPDAO.DeleteDSP(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more adAccountsIds as a AdAccounts to a DSP
	//----------------------------------------------------------------------------
func AddAdAccountsToDSP(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dSPId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DSP DAO
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.AddAdAccountsToDSP(dSPId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more adAccountsIds as a AdAccounts from a DSP
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAdAccountsFromDSP(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	dSPId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountsIds,_ := vars["adAccountsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the DSP DAO
	//----------------------------------------------------------------------------
	requestResult := DSPDAO.RemoveAdAccountsFromDSP(dSPId, adAccountsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
