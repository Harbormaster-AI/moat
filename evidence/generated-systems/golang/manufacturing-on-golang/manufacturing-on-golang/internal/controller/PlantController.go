package controller

import (
    PlantDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to PlantDAO for database creation
//----------------------------------------------------------------------------
func CreatePlant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Plant model
	//----------------------------------------------------------------------------
	data := model.Plant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Plant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Plant data access object to create
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.CreatePlant( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to PlantDAO to find the relevant Plant
//----------------------------------------------------------------------------
func GetPlant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Plant data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.GetPlant(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to PlantDAO for database read of all Plants
//----------------------------------------------------------------------------
func GetAllPlant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Plant data access object to get all
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.GetAllPlant()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to PlantDAO for database save
//----------------------------------------------------------------------------
func UpdatePlant(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Plant model
	//----------------------------------------------------------------------------
	var data = model.Plant{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Plant model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Plant data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.UpdatePlant(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to PlantDAO for database deletion
//----------------------------------------------------------------------------
func DeletePlant(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Plant data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := PlantDAO.DeletePlant(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Enterprise on a Plant
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEnterpriseToPlant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	enterpriseId,_ := strconv.ParseUint( vars["enterpriseId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AssignEnterpriseToPlant(plantId, enterpriseId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Enterprise on a Plant
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEnterpriseFromPlant( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.UnassignEnterpriseFromPlant(plantId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more productionLinesIds as a ProductionLines to a Plant
	//----------------------------------------------------------------------------
func AddProductionLinesToPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionLinesIds,_ := vars["productionLinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AddProductionLinesToPlant(plantId, productionLinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productionLinesIds as a ProductionLines from a Plant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductionLinesFromPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionLinesIds,_ := vars["productionLinesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.RemoveProductionLinesFromPlant(plantId, productionLinesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more workCentersIds as a WorkCenters to a Plant
	//----------------------------------------------------------------------------
func AddWorkCentersToPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCentersIds,_ := vars["workCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AddWorkCentersToPlant(plantId, workCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more workCentersIds as a WorkCenters from a Plant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWorkCentersFromPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workCentersIds,_ := vars["workCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.RemoveWorkCentersFromPlant(plantId, workCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more warehousesIds as a Warehouses to a Plant
	//----------------------------------------------------------------------------
func AddWarehousesToPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehousesIds,_ := vars["warehousesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AddWarehousesToPlant(plantId, warehousesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more warehousesIds as a Warehouses from a Plant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveWarehousesFromPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	warehousesIds,_ := vars["warehousesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.RemoveWarehousesFromPlant(plantId, warehousesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more assetsIds as a Assets to a Plant
	//----------------------------------------------------------------------------
func AddAssetsToPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetsIds,_ := vars["assetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AddAssetsToPlant(plantId, assetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more assetsIds as a Assets from a Plant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveAssetsFromPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	assetsIds,_ := vars["assetsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.RemoveAssetsFromPlant(plantId, assetsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more productionSchedulesIds as a ProductionSchedules to a Plant
	//----------------------------------------------------------------------------
func AddProductionSchedulesToPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionSchedulesIds,_ := vars["productionSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.AddProductionSchedulesToPlant(plantId, productionSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more productionSchedulesIds as a ProductionSchedules from a Plant
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveProductionSchedulesFromPlant(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	plantId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	productionSchedulesIds,_ := vars["productionSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Plant DAO
	//----------------------------------------------------------------------------
	requestResult := PlantDAO.RemoveProductionSchedulesFromPlant(plantId, productionSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
