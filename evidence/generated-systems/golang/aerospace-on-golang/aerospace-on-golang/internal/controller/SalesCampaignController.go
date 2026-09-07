package controller

import (
    SalesCampaignDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to SalesCampaignDAO for database creation
//----------------------------------------------------------------------------
func CreateSalesCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesCampaign model
	//----------------------------------------------------------------------------
	data := model.SalesCampaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesCampaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign data access object to create
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.CreateSalesCampaign( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to SalesCampaignDAO to find the relevant SalesCampaign
//----------------------------------------------------------------------------
func GetSalesCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesCampaign data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.GetSalesCampaign(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to SalesCampaignDAO for database read of all SalesCampaigns
//----------------------------------------------------------------------------
func GetAllSalesCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign data access object to get all
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.GetAllSalesCampaign()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to SalesCampaignDAO for database save
//----------------------------------------------------------------------------
func UpdateSalesCampaign(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty SalesCampaign model
	//----------------------------------------------------------------------------
	var data = model.SalesCampaign{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a SalesCampaign model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.UpdateSalesCampaign(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to SalesCampaignDAO for database deletion
//----------------------------------------------------------------------------
func DeleteSalesCampaign(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the SalesCampaign data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := SalesCampaignDAO.DeleteSalesCampaign(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Region on a SalesCampaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignRegionToSalesCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	regionId,_ := strconv.ParseUint( vars["regionId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.AssignRegionToSalesCampaign(salesCampaignId, regionId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Region on a SalesCampaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignRegionFromSalesCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.UnassignRegionFromSalesCampaign(salesCampaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Operator on a SalesCampaign
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOperatorToSalesCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	operatorId,_ := strconv.ParseUint( vars["operatorId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.AssignOperatorToSalesCampaign(salesCampaignId, operatorId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Operator on a SalesCampaign
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOperatorFromSalesCampaign( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.UnassignOperatorFromSalesCampaign(salesCampaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more quotesIds as a Quotes to a SalesCampaign
	//----------------------------------------------------------------------------
func AddQuotesToSalesCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.AddQuotesToSalesCampaign(salesCampaignId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more quotesIds as a Quotes from a SalesCampaign
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQuotesFromSalesCampaign(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	salesCampaignId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	quotesIds,_ := vars["quotesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the SalesCampaign DAO
	//----------------------------------------------------------------------------
	requestResult := SalesCampaignDAO.RemoveQuotesFromSalesCampaign(salesCampaignId, quotesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
