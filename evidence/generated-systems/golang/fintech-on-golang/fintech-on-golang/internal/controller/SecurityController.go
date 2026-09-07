package controller

import (
    SecurityDAO "fintech-on-golang/internal/dao"
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SecurityDAO for database creation
//----------------------------------------------------------------------------
func CreateSecurity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Security model
	//----------------------------------------------------------------------------
	data := model.Security{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Security model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Security data access object to create
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.CreateSecurity( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SecurityDAO to find the relevant Security
//----------------------------------------------------------------------------
func GetSecurity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Security data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.GetSecurity(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SecurityDAO for database read of all Securitys
//----------------------------------------------------------------------------
func GetAllSecurity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Security data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.GetAllSecurity()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SecurityDAO for database save
//----------------------------------------------------------------------------
func UpdateSecurity(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Security model
	//----------------------------------------------------------------------------
	var data = model.Security{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Security model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Security data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.UpdateSecurity(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SecurityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSecurity(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Security data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SecurityDAO.DeleteSecurity(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more positionsIds as a Positions to a Security
	//----------------------------------------------------------------------------
func AddPositionsToSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.AddPositionsToSecurity(securityId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more positionsIds as a Positions from a Security
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePositionsFromSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	positionsIds,_ := vars["positionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.RemovePositionsFromSecurity(securityId, positionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more tradesIds as a Trades to a Security
	//----------------------------------------------------------------------------
func AddTradesToSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.AddTradesToSecurity(securityId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more tradesIds as a Trades from a Security
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveTradesFromSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	tradesIds,_ := vars["tradesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.RemoveTradesFromSecurity(securityId, tradesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more ordersIds as a Orders to a Security
	//----------------------------------------------------------------------------
func AddOrdersToSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.AddOrdersToSecurity(securityId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more ordersIds as a Orders from a Security
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOrdersFromSecurity(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	securityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	ordersIds,_ := vars["ordersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Security DAO
	//----------------------------------------------------------------------------
	requestResult := SecurityDAO.RemoveOrdersFromSecurity(securityId, ordersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
