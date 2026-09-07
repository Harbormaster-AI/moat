package controller

import (
    RecordsRepositoryDAO "governance-on-golang/internal/dao"
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to RecordsRepositoryDAO for database creation
//----------------------------------------------------------------------------
func CreateRecordsRepository(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RecordsRepository model
	//----------------------------------------------------------------------------
	data := model.RecordsRepository{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RecordsRepository model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository data access object to create
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.CreateRecordsRepository( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to RecordsRepositoryDAO to find the relevant RecordsRepository
//----------------------------------------------------------------------------
func GetRecordsRepository(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RecordsRepository data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.GetRecordsRepository(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to RecordsRepositoryDAO for database read of all RecordsRepositorys
//----------------------------------------------------------------------------
func GetAllRecordsRepository(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository data access object to get all
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.GetAllRecordsRepository()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to RecordsRepositoryDAO for database save
//----------------------------------------------------------------------------
func UpdateRecordsRepository(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty RecordsRepository model
	//----------------------------------------------------------------------------
	var data = model.RecordsRepository{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a RecordsRepository model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.UpdateRecordsRepository(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to RecordsRepositoryDAO for database deletion
//----------------------------------------------------------------------------
func DeleteRecordsRepository(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the RecordsRepository data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := RecordsRepositoryDAO.DeleteRecordsRepository(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Organization on a RecordsRepository
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignOrganizationToRecordsRepository(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	organizationId,_ := strconv.ParseUint( vars["organizationId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.AssignOrganizationToRecordsRepository(recordsRepositoryId, organizationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Organization on a RecordsRepository
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignOrganizationFromRecordsRepository( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.UnassignOrganizationFromRecordsRepository(recordsRepositoryId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more recordsIds as a Records to a RecordsRepository
	//----------------------------------------------------------------------------
func AddRecordsToRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.AddRecordsToRecordsRepository(recordsRepositoryId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more recordsIds as a Records from a RecordsRepository
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRecordsFromRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	recordsIds,_ := vars["recordsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.RemoveRecordsFromRecordsRepository(recordsRepositoryId, recordsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more systemsIds as a Systems to a RecordsRepository
	//----------------------------------------------------------------------------
func AddSystemsToRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	systemsIds,_ := vars["systemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.AddSystemsToRecordsRepository(recordsRepositoryId, systemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more systemsIds as a Systems from a RecordsRepository
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveSystemsFromRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	systemsIds,_ := vars["systemsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.RemoveSystemsFromRecordsRepository(recordsRepositoryId, systemsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more retentionSchedulesIds as a RetentionSchedules to a RecordsRepository
	//----------------------------------------------------------------------------
func AddRetentionSchedulesToRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionSchedulesIds,_ := vars["retentionSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.AddRetentionSchedulesToRecordsRepository(recordsRepositoryId, retentionSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more retentionSchedulesIds as a RetentionSchedules from a RecordsRepository
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveRetentionSchedulesFromRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	retentionSchedulesIds,_ := vars["retentionSchedulesIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.RemoveRetentionSchedulesFromRecordsRepository(recordsRepositoryId, retentionSchedulesIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
	//----------------------------------------------------------------------------
	// adds one or more legalHoldsIds as a LegalHolds to a RecordsRepository
	//----------------------------------------------------------------------------
func AddLegalHoldsToRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.AddLegalHoldsToRecordsRepository(recordsRepositoryId, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more legalHoldsIds as a LegalHolds from a RecordsRepository
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveLegalHoldsFromRecordsRepository(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	recordsRepositoryId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	legalHoldsIds,_ := vars["legalHoldsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the RecordsRepository DAO
	//----------------------------------------------------------------------------
	requestResult := RecordsRepositoryDAO.RemoveLegalHoldsFromRecordsRepository(recordsRepositoryId, legalHoldsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
