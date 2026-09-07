package controller

import (
    SalesRegionDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SalesRegionDAO for database creation
//----------------------------------------------------------------------------
func CreateSalesRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesRegion model
	//----------------------------------------------------------------------------
	data := model.SalesRegion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesRegion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion data access object to create
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.CreateSalesRegion( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SalesRegionDAO to find the relevant SalesRegion
//----------------------------------------------------------------------------
func GetSalesRegion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesRegion data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.GetSalesRegion(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SalesRegionDAO for database read of all SalesRegions
//----------------------------------------------------------------------------
func GetAllSalesRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.GetAllSalesRegion()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SalesRegionDAO for database save
//----------------------------------------------------------------------------
func UpdateSalesRegion(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesRegion model
	//----------------------------------------------------------------------------
	var data = model.SalesRegion{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesRegion model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.UpdateSalesRegion(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SalesRegionDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSalesRegion(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesRegion data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SalesRegionDAO.DeleteSalesRegion(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more operatorsIds as a Operators to a SalesRegion
	//----------------------------------------------------------------------------
func AddOperatorsToSalesRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatorsIds,_ := vars["operatorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion DAO
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.AddOperatorsToSalesRegion(salesRegionId, operatorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more operatorsIds as a Operators from a SalesRegion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveOperatorsFromSalesRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatorsIds,_ := vars["operatorsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion DAO
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.RemoveOperatorsFromSalesRegion(salesRegionId, operatorsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more salesCampaignsIds as a SalesCampaigns to a SalesRegion
	//----------------------------------------------------------------------------
func AddSalesCampaignsToSalesRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesCampaignsIds,_ := vars["salesCampaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion DAO
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.AddSalesCampaignsToSalesRegion(salesRegionId, salesCampaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more salesCampaignsIds as a SalesCampaigns from a SalesRegion
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSalesCampaignsFromSalesRegion(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesRegionId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	salesCampaignsIds,_ := vars["salesCampaignsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesRegion DAO
	//----------------------------------------------------------------------------
	requestResult := SalesRegionDAO.RemoveSalesCampaignsFromSalesRegion(salesRegionId, salesCampaignsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
