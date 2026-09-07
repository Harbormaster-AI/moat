package controller

import (
    BOMDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BOMDAO for database creation
//----------------------------------------------------------------------------
func CreateBOM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BOM model
	//----------------------------------------------------------------------------
	data := model.BOM{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BOM model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BOM data access object to create
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.CreateBOM( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BOMDAO to find the relevant BOM
//----------------------------------------------------------------------------
func GetBOM(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BOM data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.GetBOM(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BOMDAO for database read of all BOMs
//----------------------------------------------------------------------------
func GetAllBOM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BOM data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.GetAllBOM()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BOMDAO for database save
//----------------------------------------------------------------------------
func UpdateBOM(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BOM model
	//----------------------------------------------------------------------------
	var data = model.BOM{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BOM model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BOM data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.UpdateBOM(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BOMDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBOM(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BOM data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BOMDAO.DeleteBOM(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ParentItem on a BOM
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignParentItemToBOM(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	parentItemId,_ := strconv.ParseUint( vars["parentItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOM DAO
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.AssignParentItemToBOM(bOMId, parentItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ParentItem on a BOM
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignParentItemFromBOM( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOM DAO
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.UnassignParentItemFromBOM(bOMId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more bomItemsIds as a BomItems to a BOM
	//----------------------------------------------------------------------------
func AddBomItemsToBOM(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bOMId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomItemsIds,_ := vars["bomItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BOM DAO
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.AddBomItemsToBOM(bOMId, bomItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more bomItemsIds as a BomItems from a BOM
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBomItemsFromBOM(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	bOMId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomItemsIds,_ := vars["bomItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the BOM DAO
	//----------------------------------------------------------------------------
	requestResult := BOMDAO.RemoveBomItemsFromBOM(bOMId, bomItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
