package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InterviewDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInterview - creates a new db entry
//----------------------------------------------------------------------------
func CreateInterview(obj model.Interview)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Interview with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Interview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInterview", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInterview - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInterview(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Interview

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Interview with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Interview using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Interview using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInterview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInterview - returns all
//----------------------------------------------------------------------------
func GetAllInterview()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Interview

	//----------------------------------------------------------------------------
	// Request the ORM to find all Interview
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Interview" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Interview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInterview", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInterview - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInterview(obj model.Interview)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Interview using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Interview using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInterview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInterview - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInterview(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInterview(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Interview)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Interview using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Interview using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInterview", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Requisition on a Interview
//----------------------------------------------------------------------------
func AssignRequisitionToInterview( interviewId uint64, requisitionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

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
			// assign the Requisition	to the Interview
			//----------------------------------------------------------------------------
			parentObj.Requisition = &childObj

			//----------------------------------------------------------------------------
			// save the Interview
			//----------------------------------------------------------------------------
			return UpdateInterview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requisition", requisitionId )
			return utils.RequestResult{false, msg, "assignRequisition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Requisition on a Interview
//----------------------------------------------------------------------------
func UnassignRequisitionFromInterview(interviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

		//----------------------------------------------------------------------------
		// assign an empty JobRequisition to the Requisition
		//----------------------------------------------------------------------------
		parentObj.Requisition = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Requisition
		//----------------------------------------------------------------------------
		parentObj.RequisitionId = nil;

		//----------------------------------------------------------------------------
		// save the Interview
		//----------------------------------------------------------------------------
		return UpdateInterview(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Candidate on a Interview
//----------------------------------------------------------------------------
func AssignCandidateToInterview( interviewId uint64, candidateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

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
			// assign the Candidate	to the Interview
			//----------------------------------------------------------------------------
			parentObj.Candidate = &childObj

			//----------------------------------------------------------------------------
			// save the Interview
			//----------------------------------------------------------------------------
			return UpdateInterview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidate", candidateId )
			return utils.RequestResult{false, msg, "assignCandidate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Candidate on a Interview
//----------------------------------------------------------------------------
func UnassignCandidateFromInterview(interviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

		//----------------------------------------------------------------------------
		// assign an empty Candidate to the Candidate
		//----------------------------------------------------------------------------
		parentObj.Candidate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Candidate
		//----------------------------------------------------------------------------
		parentObj.CandidateId = nil;

		//----------------------------------------------------------------------------
		// save the Interview
		//----------------------------------------------------------------------------
		return UpdateInterview(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more interviewersIds as a Interviewers to a Interview
//----------------------------------------------------------------------------
func AddInterviewersToInterview ( interviewId uint64, interviewersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewersIds, ",")

		for _, interviewersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching interviewersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Interviewers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviewers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviewers", interviewersId )
				return utils.RequestResult{false, msg, "unassignInterviewers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Interview from the gorm
		//----------------------------------------------------------------------------
		return GetInterview(interviewId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more interviewersIds as a Interviewers from a Interview
//----------------------------------------------------------------------------
func RemoveInterviewersFromInterview( interviewId uint64, interviewersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Interview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInterview(interviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Interview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Interview)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewersIds, ",")

		for _, interviewersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Employee

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Employee
			// with a matching interviewersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmployeeObj from the Interviewers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviewers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviewers", interviewersId )
				return utils.RequestResult{false, msg, "removeInterviewers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Interview from the gorm
		//----------------------------------------------------------------------------
		return GetInterview(interviewId)

	} else {
		return parentRequestResult
	}
}

