package controller

import (
    ClaimReserveDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to ClaimReserveDAO for database creation
//----------------------------------------------------------------------------
func CreateClaimReserve(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClaimReserve model
	//----------------------------------------------------------------------------
	data := model.ClaimReserve{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClaimReserve model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve data access object to create
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.CreateClaimReserve( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to ClaimReserveDAO to find the relevant ClaimReserve
//----------------------------------------------------------------------------
func GetClaimReserve(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClaimReserve data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.GetClaimReserve(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to ClaimReserveDAO for database read of all ClaimReserves
//----------------------------------------------------------------------------
func GetAllClaimReserve(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve data access object to get all
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.GetAllClaimReserve()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to ClaimReserveDAO for database save
//----------------------------------------------------------------------------
func UpdateClaimReserve(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty ClaimReserve model
	//----------------------------------------------------------------------------
	var data = model.ClaimReserve{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a ClaimReserve model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.UpdateClaimReserve(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to ClaimReserveDAO for database deletion
//----------------------------------------------------------------------------
func DeleteClaimReserve(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the ClaimReserve data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := ClaimReserveDAO.DeleteClaimReserve(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Claim on a ClaimReserve
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignClaimToClaimReserve(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimReserveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	claimId,_ := strconv.ParseUint( vars["claimId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.AssignClaimToClaimReserve(claimReserveId, claimId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Claim on a ClaimReserve
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignClaimFromClaimReserve( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimReserveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.UnassignClaimFromClaimReserve(claimReserveId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

	//----------------------------------------------------------------------------
	// assigns a Exposure on a ClaimReserve
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignExposureToClaimReserve(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimReserveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	exposureId,_ := strconv.ParseUint( vars["exposureId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.AssignExposureToClaimReserve(claimReserveId, exposureId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Exposure on a ClaimReserve
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignExposureFromClaimReserve( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	claimReserveId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the ClaimReserve DAO
	//----------------------------------------------------------------------------
	requestResult := ClaimReserveDAO.UnassignExposureFromClaimReserve(claimReserveId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


