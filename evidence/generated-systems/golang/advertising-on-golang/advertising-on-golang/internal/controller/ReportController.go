package controller

import (
    ReportDAO "advertising-on-golang/internal/dao"
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ReportDAO for database creation
//----------------------------------------------------------------------------
func CreateReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Report model
	//----------------------------------------------------------------------------
	data := model.Report{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Report model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Report data access object to create
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.CreateReport( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ReportDAO to find the relevant Report
//----------------------------------------------------------------------------
func GetReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Report data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.GetReport(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ReportDAO for database read of all Reports
//----------------------------------------------------------------------------
func GetAllReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Report data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.GetAllReport()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ReportDAO for database save
//----------------------------------------------------------------------------
func UpdateReport(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Report model
	//----------------------------------------------------------------------------
	var data = model.Report{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Report model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Report data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UpdateReport(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ReportDAO for database deletion
//----------------------------------------------------------------------------
func DeleteReport(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Report data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ReportDAO.DeleteReport(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a AdAccount on a Report
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignAdAccountToReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	adAccountId,_ := strconv.ParseUint( vars["adAccountId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AssignAdAccountToReport(reportId, adAccountId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a AdAccount on a Report
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignAdAccountFromReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UnassignAdAccountFromReport(reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Campaign on a Report
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignCampaignToReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	campaignId,_ := strconv.ParseUint( vars["campaignId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AssignCampaignToReport(reportId, campaignId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Campaign on a Report
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignCampaignFromReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UnassignCampaignFromReport(reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a LineItem on a Report
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignLineItemToReport(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	lineItemId,_ := strconv.ParseUint( vars["lineItemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.AssignLineItemToReport(reportId, lineItemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a LineItem on a Report
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignLineItemFromReport( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reportId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Report DAO
	//----------------------------------------------------------------------------
	requestResult := ReportDAO.UnassignLineItemFromReport(reportId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


