package controller

import (
    ProductionLineDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ProductionLineDAO for database creation
//----------------------------------------------------------------------------
func CreateProductionLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionLine model
	//----------------------------------------------------------------------------
	data := model.ProductionLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine data access object to create
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.CreateProductionLine( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ProductionLineDAO to find the relevant ProductionLine
//----------------------------------------------------------------------------
func GetProductionLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionLine data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.GetProductionLine(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ProductionLineDAO for database read of all ProductionLines
//----------------------------------------------------------------------------
func GetAllProductionLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.GetAllProductionLine()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ProductionLineDAO for database save
//----------------------------------------------------------------------------
func UpdateProductionLine(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ProductionLine model
	//----------------------------------------------------------------------------
	var data = model.ProductionLine{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ProductionLine model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.UpdateProductionLine(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ProductionLineDAO for database deletion
//----------------------------------------------------------------------------
func DeleteProductionLine(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ProductionLine data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ProductionLineDAO.DeleteProductionLine(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Plant on a ProductionLine
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignPlantToProductionLine(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	plantId,_ := strconv.ParseUint( vars["plantId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.AssignPlantToProductionLine(productionLineId, plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Plant on a ProductionLine
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignPlantFromProductionLine( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	productionLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.UnassignPlantFromProductionLine(productionLineId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more workCentersIds as a WorkCenters to a ProductionLine
	//----------------------------------------------------------------------------
func AddWorkCentersToProductionLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productionLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCentersIds,_ := vars["workCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.AddWorkCentersToProductionLine(productionLineId, workCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workCentersIds as a WorkCenters from a ProductionLine
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkCentersFromProductionLine(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	productionLineId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCentersIds,_ := vars["workCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ProductionLine DAO
	//----------------------------------------------------------------------------
	requestResult := ProductionLineDAO.RemoveWorkCentersFromProductionLine(productionLineId, workCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
