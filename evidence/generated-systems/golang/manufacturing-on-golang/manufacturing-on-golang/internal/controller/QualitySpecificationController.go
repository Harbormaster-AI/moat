package controller

import (
    QualitySpecificationDAO "manufacturing-on-golang/internal/dao"
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to QualitySpecificationDAO for database creation
//----------------------------------------------------------------------------
func CreateQualitySpecification(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualitySpecification model
	//----------------------------------------------------------------------------
	data := model.QualitySpecification{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualitySpecification model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualitySpecification data access object to create
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.CreateQualitySpecification( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to QualitySpecificationDAO to find the relevant QualitySpecification
//----------------------------------------------------------------------------
func GetQualitySpecification(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualitySpecification data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.GetQualitySpecification(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to QualitySpecificationDAO for database read of all QualitySpecifications
//----------------------------------------------------------------------------
func GetAllQualitySpecification(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the QualitySpecification data access object to get all
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.GetAllQualitySpecification()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to QualitySpecificationDAO for database save
//----------------------------------------------------------------------------
func UpdateQualitySpecification(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty QualitySpecification model
	//----------------------------------------------------------------------------
	var data = model.QualitySpecification{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a QualitySpecification model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the QualitySpecification data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.UpdateQualitySpecification(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to QualitySpecificationDAO for database deletion
//----------------------------------------------------------------------------
func DeleteQualitySpecification(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the QualitySpecification data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := QualitySpecificationDAO.DeleteQualitySpecification(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Item on a QualitySpecification
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignItemToQualitySpecification(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualitySpecificationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	itemId,_ := strconv.ParseUint( vars["itemId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualitySpecification DAO
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.AssignItemToQualitySpecification(qualitySpecificationId, itemId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Item on a QualitySpecification
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignItemFromQualitySpecification( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	qualitySpecificationId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the QualitySpecification DAO
	//----------------------------------------------------------------------------
	requestResult := QualitySpecificationDAO.UnassignItemFromQualitySpecification(qualitySpecificationId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


