package controller

import (
    ItemDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ItemDAO for database creation
//----------------------------------------------------------------------------
func CreateItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Item model
	//----------------------------------------------------------------------------
	data := model.Item{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Item model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Item data access object to create
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.CreateItem( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ItemDAO to find the relevant Item
//----------------------------------------------------------------------------
func GetItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Item data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.GetItem(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ItemDAO for database read of all Items
//----------------------------------------------------------------------------
func GetAllItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Item data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.GetAllItem()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ItemDAO for database save
//----------------------------------------------------------------------------
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Item model
	//----------------------------------------------------------------------------
	var data = model.Item{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Item model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Item data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.UpdateItem(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ItemDAO for database deletion
//----------------------------------------------------------------------------
func DeleteItem(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Item data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ItemDAO.DeleteItem(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a BusinessUnit on a Item
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignBusinessUnitToItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	businessUnitId,_ := strconv.ParseUint( vars["businessUnitId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AssignBusinessUnitToItem(itemId, businessUnitId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a BusinessUnit on a Item
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignBusinessUnitFromItem( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.UnassignBusinessUnitFromItem(itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more bomsIds as a Boms to a Item
	//----------------------------------------------------------------------------
func AddBomsToItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomsIds,_ := vars["bomsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AddBomsToItem(itemId, bomsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more bomsIds as a Boms from a Item
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveBomsFromItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	bomsIds,_ := vars["bomsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.RemoveBomsFromItem(itemId, bomsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more routingsIds as a Routings to a Item
	//----------------------------------------------------------------------------
func AddRoutingsToItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	routingsIds,_ := vars["routingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AddRoutingsToItem(itemId, routingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more routingsIds as a Routings from a Item
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRoutingsFromItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	routingsIds,_ := vars["routingsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.RemoveRoutingsFromItem(itemId, routingsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more suppliersIds as a Suppliers to a Item
	//----------------------------------------------------------------------------
func AddSuppliersToItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AddSuppliersToItem(itemId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more suppliersIds as a Suppliers from a Item
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSuppliersFromItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	suppliersIds,_ := vars["suppliersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.RemoveSuppliersFromItem(itemId, suppliersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more qualitySpecificationsIds as a QualitySpecifications to a Item
	//----------------------------------------------------------------------------
func AddQualitySpecificationsToItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	qualitySpecificationsIds,_ := vars["qualitySpecificationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AddQualitySpecificationsToItem(itemId, qualitySpecificationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more qualitySpecificationsIds as a QualitySpecifications from a Item
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveQualitySpecificationsFromItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	qualitySpecificationsIds,_ := vars["qualitySpecificationsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.RemoveQualitySpecificationsFromItem(itemId, qualitySpecificationsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a Item
	//----------------------------------------------------------------------------
func AddInventoryItemsToItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.AddInventoryItemsToItem(itemId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a Item
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromItem(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	itemId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Item DAO
	//----------------------------------------------------------------------------
	requestResult := ItemDAO.RemoveInventoryItemsFromItem(itemId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
