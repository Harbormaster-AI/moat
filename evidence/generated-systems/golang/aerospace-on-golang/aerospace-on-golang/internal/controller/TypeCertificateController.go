package controller

import (
    TypeCertificateDAO "aerospace-on-golang/internal/dao"
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//----------------------------------------------------------------------------
// Create controller, delegates to TypeCertificateDAO for database creation
//----------------------------------------------------------------------------
func CreateTypeCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TypeCertificate model
	//----------------------------------------------------------------------------
	data := model.TypeCertificate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TypeCertificate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TypeCertificate data access object to create
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.CreateTypeCertificate( data )
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Get controller, delegates to TypeCertificateDAO to find the relevant TypeCertificate
//----------------------------------------------------------------------------
func GetTypeCertificate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TypeCertificate data access object
	// find the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.GetTypeCertificate(ID)
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


//----------------------------------------------------------------------------
// GetAll controller, delegates to TypeCertificateDAO for database read of all TypeCertificates
//----------------------------------------------------------------------------
func GetAllTypeCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Delegate to the TypeCertificate data access object to get all
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.GetAllTypeCertificate()
	
	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res,_ := json.Marshal(requestResult)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Update controller, delegates to TypeCertificateDAO for database save
//----------------------------------------------------------------------------
func UpdateTypeCertificate(w http.ResponseWriter, r *http.Request) {
	//----------------------------------------------------------------------------
	// Initialize an empty TypeCertificate model
	//----------------------------------------------------------------------------
	var data = model.TypeCertificate{}
	
	//----------------------------------------------------------------------------
	// Parse the body into a TypeCertificate model structure
	//----------------------------------------------------------------------------
	utils.ParseBody(r, data)

	//----------------------------------------------------------------------------
	// Delegate to the TypeCertificate data access object
	// update the one with the matching identifier
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.UpdateTypeCertificate(data)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//----------------------------------------------------------------------------
// Delete controller, delegates to TypeCertificateDAO for database deletion
//----------------------------------------------------------------------------
func DeleteTypeCertificate(w http.ResponseWriter, r *http.Request) {
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
	// Delegate to the TypeCertificate data access object
	// delete the one with the matching identifier
	//----------------------------------------------------------------------------	
	requestResult := TypeCertificateDAO.DeleteTypeCertificate(ID)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// assigns a Program on a TypeCertificate
	// delegates to an ORM handler
	///----------------------------------------------------------------------------
func AssignProgramToTypeCertificate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	typeCertificateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)
	programId,_ := strconv.ParseUint( vars["programId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TypeCertificate DAO
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.AssignProgramToTypeCertificate(typeCertificateId, programId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

	//----------------------------------------------------------------------------
	// unassigns a Program on a TypeCertificate
	// delegates to the ORM handler
	//----------------------------------------------------------------------------
func UnassignProgramFromTypeCertificate( w http.ResponseWriter, r *http.Request ) {

	vars := mux.Vars(r)

	//----------------------------------------------------------------------------
	// Retrieve the id params
	//----------------------------------------------------------------------------
	typeCertificateId,_ := strconv.ParseUint( vars["parentId"], 10, 64)

	//----------------------------------------------------------------------------
	// Delegate to the TypeCertificate DAO
	//----------------------------------------------------------------------------
	requestResult := TypeCertificateDAO.UnassignProgramFromTypeCertificate(typeCertificateId)

	//----------------------------------------------------------------------------
	// Marshal the model into a JSON object
	//----------------------------------------------------------------------------
	res, _ := json.Marshal(requestResult)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


