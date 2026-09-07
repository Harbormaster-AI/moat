package controller

import (
    UnderwriterDAO "insurance-on-golang/internal/dao"
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to UnderwriterDAO for database creation
//----------------------------------------------------------------------------
func CreateUnderwriter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Underwriter model
	//----------------------------------------------------------------------------
	data := model.Underwriter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Underwriter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter data access object to create
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.CreateUnderwriter( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to UnderwriterDAO to find the relevant Underwriter
//----------------------------------------------------------------------------
func GetUnderwriter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Underwriter data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.GetUnderwriter(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to UnderwriterDAO for database read of all Underwriters
//----------------------------------------------------------------------------
func GetAllUnderwriter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the Underwriter data access object to get all
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.GetAllUnderwriter()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to UnderwriterDAO for database save
//----------------------------------------------------------------------------
func UpdateUnderwriter(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty Underwriter model
	//----------------------------------------------------------------------------
	var data = model.Underwriter{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a Underwriter model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.UpdateUnderwriter(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to UnderwriterDAO for database deletion
//----------------------------------------------------------------------------
func DeleteUnderwriter(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the Underwriter data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := UnderwriterDAO.DeleteUnderwriter(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Insurer on a Underwriter
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignInsurerToUnderwriter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwriterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	insurerId,_ := strconv.ParseUint( vars["insurerId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.AssignInsurerToUnderwriter(underwriterId, insurerId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Insurer on a Underwriter
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignInsurerFromUnderwriter( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	underwriterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.UnassignInsurerFromUnderwriter(underwriterId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


	//----------------------------------------------------------------------------
	// adds one or more decisionsIds as a Decisions to a Underwriter
	//----------------------------------------------------------------------------
func AddDecisionsToUnderwriter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	underwriterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	decisionsIds,_ := vars["decisionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.AddDecisionsToUnderwriter(underwriterId, decisionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// removes one or more decisionsIds as a Decisions from a Underwriter
	// delegates via URI to an ORM handler
	//----------------------------------------------------------------------------
func RemoveDecisionsFromUnderwriter(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id and child ids
	//----------------------------------------------------------------------------
	underwriterId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	decisionsIds,_ := vars["decisionsIds"]

	//----------------------------------------------------------------------------
	// Delegate to the Underwriter DAO
	//----------------------------------------------------------------------------
	requestResult := UnderwriterDAO.RemoveDecisionsFromUnderwriter(underwriterId, decisionsIds)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}
		
