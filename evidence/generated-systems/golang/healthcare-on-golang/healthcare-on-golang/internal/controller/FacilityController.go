package controller

import (
    FacilityDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to FacilityDAO for database creation
//----------------------------------------------------------------------------
func CreateFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Facility model
	//----------------------------------------------------------------------------
	data := model.Facility{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Facility model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Facility data access object to create
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.CreateFacility( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to FacilityDAO to find the relevant Facility
//----------------------------------------------------------------------------
func GetFacility(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Facility data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.GetFacility(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to FacilityDAO for database read of all Facilitys
//----------------------------------------------------------------------------
func GetAllFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Facility data access object to get all
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.GetAllFacility()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to FacilityDAO for database save
//----------------------------------------------------------------------------
func UpdateFacility(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Facility model
	//----------------------------------------------------------------------------
	var data = model.Facility{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Facility model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Facility data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.UpdateFacility(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to FacilityDAO for database deletion
//----------------------------------------------------------------------------
func DeleteFacility(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Facility data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := FacilityDAO.DeleteFacility(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a HealthSystem on a Facility
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignHealthSystemToFacility(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	healthSystemId,_ := strconv.ParseUint( vars["healthSystemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AssignHealthSystemToFacility(facilityId, healthSystemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a HealthSystem on a Facility
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignHealthSystemFromFacility( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.UnassignHealthSystemFromFacility(facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more departmentsIds as a Departments to a Facility
	//----------------------------------------------------------------------------
func AddDepartmentsToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddDepartmentsToFacility(facilityId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more departmentsIds as a Departments from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDepartmentsFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	departmentsIds,_ := vars["departmentsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemoveDepartmentsFromFacility(facilityId, departmentsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more careTeamsIds as a CareTeams to a Facility
	//----------------------------------------------------------------------------
func AddCareTeamsToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddCareTeamsToFacility(facilityId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more careTeamsIds as a CareTeams from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveCareTeamsFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	careTeamsIds,_ := vars["careTeamsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemoveCareTeamsFromFacility(facilityId, careTeamsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more laboratoriesIds as a Laboratories to a Facility
	//----------------------------------------------------------------------------
func AddLaboratoriesToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoriesIds,_ := vars["laboratoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddLaboratoriesToFacility(facilityId, laboratoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more laboratoriesIds as a Laboratories from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLaboratoriesFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	laboratoriesIds,_ := vars["laboratoriesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemoveLaboratoriesFromFacility(facilityId, laboratoriesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more imagingCentersIds as a ImagingCenters to a Facility
	//----------------------------------------------------------------------------
func AddImagingCentersToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingCentersIds,_ := vars["imagingCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddImagingCentersToFacility(facilityId, imagingCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingCentersIds as a ImagingCenters from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingCentersFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingCentersIds,_ := vars["imagingCentersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemoveImagingCentersFromFacility(facilityId, imagingCentersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more pharmaciesIds as a Pharmacies to a Facility
	//----------------------------------------------------------------------------
func AddPharmaciesToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pharmaciesIds,_ := vars["pharmaciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddPharmaciesToFacility(facilityId, pharmaciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more pharmaciesIds as a Pharmacies from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePharmaciesFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	pharmaciesIds,_ := vars["pharmaciesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemovePharmaciesFromFacility(facilityId, pharmaciesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more inventoryItemsIds as a InventoryItems to a Facility
	//----------------------------------------------------------------------------
func AddInventoryItemsToFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.AddInventoryItemsToFacility(facilityId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more inventoryItemsIds as a InventoryItems from a Facility
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveInventoryItemsFromFacility(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	facilityId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	inventoryItemsIds,_ := vars["inventoryItemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Facility DAO
	//----------------------------------------------------------------------------
	requestResult := FacilityDAO.RemoveInventoryItemsFromFacility(facilityId, inventoryItemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
