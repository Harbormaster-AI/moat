package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BackgroundCheckDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBackgroundCheck - creates a new db entry
//----------------------------------------------------------------------------
func CreateBackgroundCheck(obj model.BackgroundCheck)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BackgroundCheck with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BackgroundCheck", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBackgroundCheck", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBackgroundCheck - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBackgroundCheck(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BackgroundCheck

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BackgroundCheck with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BackgroundCheck using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BackgroundCheck using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBackgroundCheck", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBackgroundCheck - returns all
//----------------------------------------------------------------------------
func GetAllBackgroundCheck()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BackgroundCheck

	//----------------------------------------------------------------------------
	// Request the ORM to find all BackgroundCheck
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BackgroundCheck" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BackgroundCheck", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBackgroundCheck", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBackgroundCheck - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBackgroundCheck(obj model.BackgroundCheck)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BackgroundCheck using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BackgroundCheck using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBackgroundCheck", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBackgroundCheck - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBackgroundCheck(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBackgroundCheck(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BackgroundCheck using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BackgroundCheck using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBackgroundCheck", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Candidate on a BackgroundCheck
//----------------------------------------------------------------------------
func AssignCandidateToBackgroundCheck( backgroundCheckId uint64, candidateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

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
			// assign the Candidate	to the BackgroundCheck
			//----------------------------------------------------------------------------
			parentObj.Candidate = &childObj

			//----------------------------------------------------------------------------
			// save the BackgroundCheck
			//----------------------------------------------------------------------------
			return UpdateBackgroundCheck(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidate", candidateId )
			return utils.RequestResult{false, msg, "assignCandidate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Candidate on a BackgroundCheck
//----------------------------------------------------------------------------
func UnassignCandidateFromBackgroundCheck(backgroundCheckId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// assign an empty Candidate to the Candidate
		//----------------------------------------------------------------------------
		parentObj.Candidate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Candidate
		//----------------------------------------------------------------------------
		parentObj.CandidateId = nil;

		//----------------------------------------------------------------------------
		// save the BackgroundCheck
		//----------------------------------------------------------------------------
		return UpdateBackgroundCheck(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Requisition on a BackgroundCheck
//----------------------------------------------------------------------------
func AssignRequisitionToBackgroundCheck( backgroundCheckId uint64, requisitionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobRequisition

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobRequisition with a
		// matching requisitionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, requisitionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Requisition	to the BackgroundCheck
			//----------------------------------------------------------------------------
			parentObj.Requisition = &childObj

			//----------------------------------------------------------------------------
			// save the BackgroundCheck
			//----------------------------------------------------------------------------
			return UpdateBackgroundCheck(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requisition", requisitionId )
			return utils.RequestResult{false, msg, "assignRequisition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Requisition on a BackgroundCheck
//----------------------------------------------------------------------------
func UnassignRequisitionFromBackgroundCheck(backgroundCheckId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// assign an empty JobRequisition to the Requisition
		//----------------------------------------------------------------------------
		parentObj.Requisition = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Requisition
		//----------------------------------------------------------------------------
		parentObj.RequisitionId = nil;

		//----------------------------------------------------------------------------
		// save the BackgroundCheck
		//----------------------------------------------------------------------------
		return UpdateBackgroundCheck(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Report on a BackgroundCheck
//----------------------------------------------------------------------------
func AssignReportToBackgroundCheck( backgroundCheckId uint64, reportId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Document

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Document with a
		// matching reportId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, reportId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Report	to the BackgroundCheck
			//----------------------------------------------------------------------------
			parentObj.Report = &childObj

			//----------------------------------------------------------------------------
			// save the BackgroundCheck
			//----------------------------------------------------------------------------
			return UpdateBackgroundCheck(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Report", reportId )
			return utils.RequestResult{false, msg, "assignReport", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Report on a BackgroundCheck
//----------------------------------------------------------------------------
func UnassignReportFromBackgroundCheck(backgroundCheckId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BackgroundCheck with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBackgroundCheck(backgroundCheckId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BackgroundCheck so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BackgroundCheck)

		//----------------------------------------------------------------------------
		// assign an empty Document to the Report
		//----------------------------------------------------------------------------
		parentObj.Report = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Report
		//----------------------------------------------------------------------------
		parentObj.ReportId = nil;

		//----------------------------------------------------------------------------
		// save the BackgroundCheck
		//----------------------------------------------------------------------------
		return UpdateBackgroundCheck(parentObj)

	} else {
		return parentRequestResult
	}

}


