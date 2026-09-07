package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DocumentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDocument - creates a new db entry
//----------------------------------------------------------------------------
func CreateDocument(obj model.Document)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a Document with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Document", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDocument", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDocument - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDocument(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Document

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Document with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Document using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Document using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDocument - returns all
//----------------------------------------------------------------------------
func GetAllDocument()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Document

	//----------------------------------------------------------------------------
	// Request the ORM to find all Document
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Document" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Document", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDocument", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDocument - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDocument(obj model.Document)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a Document using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Document using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDocument - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDocument(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDocument(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Document using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Document using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDocument", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Candidate on a Document
//----------------------------------------------------------------------------
func AssignCandidateToDocument( documentId uint64, candidateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Candidate

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Candidate with a
		// matching candidateId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, candidateId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Candidate	to the Document
			//----------------------------------------------------------------------------
			parentObj.Candidate = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidate", candidateId )
			return utils.RequestResult{false, msg, "assignCandidate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Candidate on a Document
//----------------------------------------------------------------------------
func UnassignCandidateFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Candidate to the Candidate
		//----------------------------------------------------------------------------
		parentObj.Candidate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Candidate
		//----------------------------------------------------------------------------
		parentObj.CandidateId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a Document
//----------------------------------------------------------------------------
func AssignEmployeeToDocument( documentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the Document
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a Document
//----------------------------------------------------------------------------
func UnassignEmployeeFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}


