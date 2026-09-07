package controller

import (
    CreativeApprovalDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CreativeApprovalDAO for database creation
//----------------------------------------------------------------------------
func CreateCreativeApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeApproval model
	//----------------------------------------------------------------------------
	data := model.CreativeApproval{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeApproval model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval data access object to create
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.CreateCreativeApproval( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CreativeApprovalDAO to find the relevant CreativeApproval
//----------------------------------------------------------------------------
func GetCreativeApproval(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeApproval data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.GetCreativeApproval(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CreativeApprovalDAO for database read of all CreativeApprovals
//----------------------------------------------------------------------------
func GetAllCreativeApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.GetAllCreativeApproval()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CreativeApprovalDAO for database save
//----------------------------------------------------------------------------
func UpdateCreativeApproval(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeApproval model
	//----------------------------------------------------------------------------
	var data = model.CreativeApproval{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeApproval model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.UpdateCreativeApproval(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CreativeApprovalDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCreativeApproval(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeApproval data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CreativeApprovalDAO.DeleteCreativeApproval(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a CreativeAsset on a CreativeApproval
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCreativeAssetToCreativeApproval(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeApprovalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	creativeAssetId,_ := strconv.ParseUint( vars["creativeAssetId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.AssignCreativeAssetToCreativeApproval(creativeApprovalId, creativeAssetId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a CreativeAsset on a CreativeApproval
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCreativeAssetFromCreativeApproval( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeApprovalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.UnassignCreativeAssetFromCreativeApproval(creativeApprovalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Publisher on a CreativeApproval
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPublisherToCreativeApproval(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeApprovalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	publisherId,_ := strconv.ParseUint( vars["publisherId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.AssignPublisherToCreativeApproval(creativeApprovalId, publisherId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Publisher on a CreativeApproval
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPublisherFromCreativeApproval( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	creativeApprovalId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeApproval DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeApprovalDAO.UnassignPublisherFromCreativeApproval(creativeApprovalId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


