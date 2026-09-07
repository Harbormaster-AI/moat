package controller

import (
    ImagingReportDAO "healthcare-on-golang/internal/dao"
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ImagingReportDAO for database creation
//----------------------------------------------------------------------------
func CreateImagingReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingReport model
	//----------------------------------------------------------------------------
	data := model.ImagingReport{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingReport model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport data access object to create
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.CreateImagingReport( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ImagingReportDAO to find the relevant ImagingReport
//----------------------------------------------------------------------------
func GetImagingReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingReport data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.GetImagingReport(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ImagingReportDAO for database read of all ImagingReports
//----------------------------------------------------------------------------
func GetAllImagingReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.GetAllImagingReport()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ImagingReportDAO for database save
//----------------------------------------------------------------------------
func UpdateImagingReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ImagingReport model
	//----------------------------------------------------------------------------
	var data = model.ImagingReport{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ImagingReport model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.UpdateImagingReport(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ImagingReportDAO for database deletion
//----------------------------------------------------------------------------
func DeleteImagingReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ImagingReport data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ImagingReportDAO.DeleteImagingReport(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ImagingOrder on a ImagingReport
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignImagingOrderToImagingReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingOrderId,_ := strconv.ParseUint( vars["imagingOrderId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.AssignImagingOrderToImagingReport(imagingReportId, imagingOrderId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ImagingOrder on a ImagingReport
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignImagingOrderFromImagingReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.UnassignImagingOrderFromImagingReport(imagingReportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Clinician on a ImagingReport
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClinicianToImagingReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	clinicianId,_ := strconv.ParseUint( vars["clinicianId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.AssignClinicianToImagingReport(imagingReportId, clinicianId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Clinician on a ImagingReport
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClinicianFromImagingReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.UnassignClinicianFromImagingReport(imagingReportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Encounter on a ImagingReport
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignEncounterToImagingReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	encounterId,_ := strconv.ParseUint( vars["encounterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.AssignEncounterToImagingReport(imagingReportId, encounterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Encounter on a ImagingReport
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignEncounterFromImagingReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.UnassignEncounterFromImagingReport(imagingReportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a ImagingCenter on a ImagingReport
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignImagingCenterToImagingReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	imagingCenterId,_ := strconv.ParseUint( vars["imagingCenterId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.AssignImagingCenterToImagingReport(imagingReportId, imagingCenterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ImagingCenter on a ImagingReport
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignImagingCenterFromImagingReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	imagingReportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ImagingReport DAO
	//----------------------------------------------------------------------------
	requestResult := ImagingReportDAO.UnassignImagingCenterFromImagingReport(imagingReportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


