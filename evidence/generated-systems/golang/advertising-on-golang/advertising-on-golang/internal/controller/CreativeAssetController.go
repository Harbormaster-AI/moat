package controller

import (
    CreativeAssetDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to CreativeAssetDAO for database creation
//----------------------------------------------------------------------------
func CreateCreativeAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeAsset model
	//----------------------------------------------------------------------------
	data := model.CreativeAsset{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeAsset model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset data access object to create
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.CreateCreativeAsset( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to CreativeAssetDAO to find the relevant CreativeAsset
//----------------------------------------------------------------------------
func GetCreativeAsset(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeAsset data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.GetCreativeAsset(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to CreativeAssetDAO for database read of all CreativeAssets
//----------------------------------------------------------------------------
func GetAllCreativeAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset data access object to get all
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.GetAllCreativeAsset()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to CreativeAssetDAO for database save
//----------------------------------------------------------------------------
func UpdateCreativeAsset(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty CreativeAsset model
	//----------------------------------------------------------------------------
	var data = model.CreativeAsset{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a CreativeAsset model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.UpdateCreativeAsset(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to CreativeAssetDAO for database deletion
//----------------------------------------------------------------------------
func DeleteCreativeAsset(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the CreativeAsset data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := CreativeAssetDAO.DeleteCreativeAsset(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


	//----------------------------------------------------------------------------
	// adds one or more filesIds as a Files to a CreativeAsset
	//----------------------------------------------------------------------------
func AddFilesToCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	filesIds,_ := vars["filesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.AddFilesToCreativeAsset(creativeAssetId, filesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more filesIds as a Files from a CreativeAsset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveFilesFromCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	filesIds,_ := vars["filesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.RemoveFilesFromCreativeAsset(creativeAssetId, filesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more approvalsIds as a Approvals to a CreativeAsset
	//----------------------------------------------------------------------------
func AddApprovalsToCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.AddApprovalsToCreativeAsset(creativeAssetId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more approvalsIds as a Approvals from a CreativeAsset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveApprovalsFromCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	approvalsIds,_ := vars["approvalsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.RemoveApprovalsFromCreativeAsset(creativeAssetId, approvalsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more variationsIds as a Variations to a CreativeAsset
	//----------------------------------------------------------------------------
func AddVariationsToCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variationsIds,_ := vars["variationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.AddVariationsToCreativeAsset(creativeAssetId, variationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more variationsIds as a Variations from a CreativeAsset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveVariationsFromCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	variationsIds,_ := vars["variationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.RemoveVariationsFromCreativeAsset(creativeAssetId, variationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more lineItemsIds as a LineItems to a CreativeAsset
	//----------------------------------------------------------------------------
func AddLineItemsToCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.AddLineItemsToCreativeAsset(creativeAssetId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more lineItemsIds as a LineItems from a CreativeAsset
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLineItemsFromCreativeAsset(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	creativeAssetId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemsIds,_ := vars["lineItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the CreativeAsset DAO
	//----------------------------------------------------------------------------
	requestResult := CreativeAssetDAO.RemoveLineItemsFromCreativeAsset(creativeAssetId, lineItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
