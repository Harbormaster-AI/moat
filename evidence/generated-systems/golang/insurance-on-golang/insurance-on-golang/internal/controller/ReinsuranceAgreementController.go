package controller

import (
    ReinsuranceAgreementDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ReinsuranceAgreementDAO for database creation
//----------------------------------------------------------------------------
func CreateReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ReinsuranceAgreement model
	//----------------------------------------------------------------------------
	data := model.ReinsuranceAgreement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ReinsuranceAgreement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement data access object to create
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.CreateReinsuranceAgreement( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ReinsuranceAgreementDAO to find the relevant ReinsuranceAgreement
//----------------------------------------------------------------------------
func GetReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ReinsuranceAgreement data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.GetReinsuranceAgreement(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ReinsuranceAgreementDAO for database read of all ReinsuranceAgreements
//----------------------------------------------------------------------------
func GetAllReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.GetAllReinsuranceAgreement()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ReinsuranceAgreementDAO for database save
//----------------------------------------------------------------------------
func UpdateReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ReinsuranceAgreement model
	//----------------------------------------------------------------------------
	var data = model.ReinsuranceAgreement{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ReinsuranceAgreement model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.UpdateReinsuranceAgreement(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ReinsuranceAgreementDAO for database deletion
//----------------------------------------------------------------------------
func DeleteReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ReinsuranceAgreement data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ReinsuranceAgreementDAO.DeleteReinsuranceAgreement(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Insurer on a ReinsuranceAgreement
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsurerToReinsuranceAgreement(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reinsuranceAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurerId,_ := strconv.ParseUint( vars["insurerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.AssignInsurerToReinsuranceAgreement(reinsuranceAgreementId, insurerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Insurer on a ReinsuranceAgreement
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsurerFromReinsuranceAgreement( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	reinsuranceAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.UnassignInsurerFromReinsuranceAgreement(reinsuranceAgreementId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more policiesIds as a Policies to a ReinsuranceAgreement
	//----------------------------------------------------------------------------
func AddPoliciesToReinsuranceAgreement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reinsuranceAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.AddPoliciesToReinsuranceAgreement(reinsuranceAgreementId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more policiesIds as a Policies from a ReinsuranceAgreement
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemovePoliciesFromReinsuranceAgreement(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	reinsuranceAgreementId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	policiesIds,_ := vars["policiesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the ReinsuranceAgreement DAO
	//----------------------------------------------------------------------------
	requestResult := ReinsuranceAgreementDAO.RemovePoliciesFromReinsuranceAgreement(reinsuranceAgreementId, policiesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
