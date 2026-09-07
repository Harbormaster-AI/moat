package controller

import (
    EvidenceDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to EvidenceDAO for database creation
//----------------------------------------------------------------------------
func CreateEvidence(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Evidence model
	//----------------------------------------------------------------------------
	data := model.Evidence{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Evidence model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence data access object to create
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.CreateEvidence( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to EvidenceDAO to find the relevant Evidence
//----------------------------------------------------------------------------
func GetEvidence(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Evidence data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.GetEvidence(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to EvidenceDAO for database read of all Evidences
//----------------------------------------------------------------------------
func GetAllEvidence(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Evidence data access object to get all
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.GetAllEvidence()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to EvidenceDAO for database save
//----------------------------------------------------------------------------
func UpdateEvidence(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Evidence model
	//----------------------------------------------------------------------------
	var data = model.Evidence{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Evidence model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.UpdateEvidence(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to EvidenceDAO for database deletion
//----------------------------------------------------------------------------
func DeleteEvidence(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Evidence data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := EvidenceDAO.DeleteEvidence(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a ControlTest on a Evidence
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlTestToEvidence(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlTestId,_ := strconv.ParseUint( vars["controlTestId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.AssignControlTestToEvidence(evidenceId, controlTestId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a ControlTest on a Evidence
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlTestFromEvidence( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.UnassignControlTestFromEvidence(evidenceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Control on a Evidence
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignControlToEvidence(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	controlId,_ := strconv.ParseUint( vars["controlId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.AssignControlToEvidence(evidenceId, controlId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Control on a Evidence
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignControlFromEvidence( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.UnassignControlFromEvidence(evidenceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Obligation on a Evidence
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignObligationToEvidence(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	obligationId,_ := strconv.ParseUint( vars["obligationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.AssignObligationToEvidence(evidenceId, obligationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Obligation on a Evidence
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignObligationFromEvidence( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.UnassignObligationFromEvidence(evidenceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Workpaper on a Evidence
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignWorkpaperToEvidence(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	workpaperId,_ := strconv.ParseUint( vars["workpaperId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.AssignWorkpaperToEvidence(evidenceId, workpaperId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Workpaper on a Evidence
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignWorkpaperFromEvidence( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	evidenceId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Evidence DAO
	//----------------------------------------------------------------------------
	requestResult := EvidenceDAO.UnassignWorkpaperFromEvidence(evidenceId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


