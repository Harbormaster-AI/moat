package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WorkAuthorizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWorkAuthorization - creates a new db entry
//----------------------------------------------------------------------------
func CreateWorkAuthorization(obj model.WorkAuthorization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a WorkAuthorization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a WorkAuthorization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWorkAuthorization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWorkAuthorization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWorkAuthorization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.WorkAuthorization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a WorkAuthorization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a WorkAuthorization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a WorkAuthorization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWorkAuthorization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWorkAuthorization - returns all
//----------------------------------------------------------------------------
func GetAllWorkAuthorization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.WorkAuthorization

	//----------------------------------------------------------------------------
	// Request the ORM to find all WorkAuthorization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all WorkAuthorization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all WorkAuthorization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWorkAuthorization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWorkAuthorization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWorkAuthorization(obj model.WorkAuthorization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a WorkAuthorization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a WorkAuthorization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWorkAuthorization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWorkAuthorization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWorkAuthorization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the WorkAuthorization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWorkAuthorization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkAuthorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.WorkAuthorization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a WorkAuthorization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a WorkAuthorization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWorkAuthorization", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a WorkAuthorization
//----------------------------------------------------------------------------
func AssignEmployeeToWorkAuthorization( workAuthorizationId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkAuthorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkAuthorization(workAuthorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkAuthorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkAuthorization)

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
			// assign the Employee	to the WorkAuthorization
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the WorkAuthorization
			//----------------------------------------------------------------------------
			return UpdateWorkAuthorization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a WorkAuthorization
//----------------------------------------------------------------------------
func UnassignEmployeeFromWorkAuthorization(workAuthorizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkAuthorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkAuthorization(workAuthorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkAuthorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkAuthorization)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the WorkAuthorization
		//----------------------------------------------------------------------------
		return UpdateWorkAuthorization(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more documentsIds as a Documents to a WorkAuthorization
//----------------------------------------------------------------------------
func AddDocumentsToWorkAuthorization ( workAuthorizationId uint64, documentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkAuthorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkAuthorization(workAuthorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkAuthorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkAuthorization)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Document

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Document
			// with a matching documentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , documentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Documents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Documents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Documents", documentsId )
				return utils.RequestResult{false, msg, "unassignDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkAuthorization from the gorm
		//----------------------------------------------------------------------------
		return GetWorkAuthorization(workAuthorizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more documentsIds as a Documents from a WorkAuthorization
//----------------------------------------------------------------------------
func RemoveDocumentsFromWorkAuthorization( workAuthorizationId uint64, documentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkAuthorization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkAuthorization(workAuthorizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkAuthorization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkAuthorization)

		// slice the ids on comma with no spaces
		ids := strings.Split( documentsIds, ",")

		for _, documentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Document

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Document
			// with a matching documentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , documentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DocumentObj from the Documents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Documents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Documents", documentsId )
				return utils.RequestResult{false, msg, "removeDocuments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkAuthorization from the gorm
		//----------------------------------------------------------------------------
		return GetWorkAuthorization(workAuthorizationId)

	} else {
		return parentRequestResult
	}
}

