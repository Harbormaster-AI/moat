package controller

import (
    BOMItemDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to BOMItemDAO for database creation
//----------------------------------------------------------------------------
func CreateBOMItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BOMItem model
	//----------------------------------------------------------------------------
	data := model.BOMItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BOMItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem data access object to create
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.CreateBOMItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to BOMItemDAO to find the relevant BOMItem
//----------------------------------------------------------------------------
func GetBOMItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BOMItem data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.GetBOMItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to BOMItemDAO for database read of all BOMItems
//----------------------------------------------------------------------------
func GetAllBOMItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the BOMItem data access object to get all
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.GetAllBOMItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to BOMItemDAO for database save
//----------------------------------------------------------------------------
func UpdateBOMItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty BOMItem model
	//----------------------------------------------------------------------------
	var data = model.BOMItem{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a BOMItem model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.UpdateBOMItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to BOMItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteBOMItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the BOMItem data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := BOMItemDAO.DeleteBOMItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Bom on a BOMItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBomToBOMItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomId,_ := strconv.ParseUint( vars["bomId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem DAO
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.AssignBomToBOMItem(bOMItemId, bomId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Bom on a BOMItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBomFromBOMItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem DAO
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.UnassignBomFromBOMItem(bOMItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Component on a BOMItem
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignComponentToBOMItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	componentId,_ := strconv.ParseUint( vars["componentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem DAO
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.AssignComponentToBOMItem(bOMItemId, componentId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Component on a BOMItem
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignComponentFromBOMItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	bOMItemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the BOMItem DAO
	//----------------------------------------------------------------------------
	requestResult := BOMItemDAO.UnassignComponentFromBOMItem(bOMItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


