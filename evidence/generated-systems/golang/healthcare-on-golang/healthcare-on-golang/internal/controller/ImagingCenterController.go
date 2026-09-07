package controller

import (
    ImagingCenterDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ImagingCenterDAO for database creation
//----------------------------------------------------------------------------
func CreateImagingCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingCenter model
	//----------------------------------------------------------------------------
	data := model.ImagingCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter data access object to create
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.CreateImagingCenter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ImagingCenterDAO to find the relevant ImagingCenter
//----------------------------------------------------------------------------
func GetImagingCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingCenter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.GetImagingCenter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ImagingCenterDAO for database read of all ImagingCenters
//----------------------------------------------------------------------------
func GetAllImagingCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.GetAllImagingCenter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ImagingCenterDAO for database save
//----------------------------------------------------------------------------
func UpdateImagingCenter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingCenter model
	//----------------------------------------------------------------------------
	var data = model.ImagingCenter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingCenter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.UpdateImagingCenter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ImagingCenterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteImagingCenter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingCenter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ImagingCenterDAO.DeleteImagingCenter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Facility on a ImagingCenter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignFacilityToImagingCenter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	facilityId,_ := strconv.ParseUint( vars["facilityId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.AssignFacilityToImagingCenter(imagingCenterId, facilityId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Facility on a ImagingCenter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignFacilityFromImagingCenter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.UnassignFacilityFromImagingCenter(imagingCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more imagingOrdersIds as a ImagingOrders to a ImagingCenter
	//----------------------------------------------------------------------------
func AddImagingOrdersToImagingCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.AddImagingOrdersToImagingCenter(imagingCenterId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingOrdersIds as a ImagingOrders from a ImagingCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingOrdersFromImagingCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrdersIds,_ := vars["imagingOrdersIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.RemoveImagingOrdersFromImagingCenter(imagingCenterId, imagingOrdersIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more imagingReportsIds as a ImagingReports to a ImagingCenter
	//----------------------------------------------------------------------------
func AddImagingReportsToImagingCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingReportsIds,_ := vars["imagingReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.AddImagingReportsToImagingCenter(imagingCenterId, imagingReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more imagingReportsIds as a ImagingReports from a ImagingCenter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveImagingReportsFromImagingCenter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	imagingCenterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingReportsIds,_ := vars["imagingReportsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ImagingCenter DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingCenterDAO.RemoveImagingReportsFromImagingCenter(imagingCenterId, imagingReportsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
