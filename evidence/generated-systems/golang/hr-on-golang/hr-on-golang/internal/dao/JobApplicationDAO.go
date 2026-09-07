package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing JobApplicationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateJobApplication - creates a new db entry
//----------------------------------------------------------------------------
func CreateJobApplication(obj model.JobApplication)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a JobApplication with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a JobApplication", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateJobApplication", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetJobApplication - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetJobApplication(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.JobApplication

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a JobApplication with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a JobApplication using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a JobApplication using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetJobApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllJobApplication - returns all
//----------------------------------------------------------------------------
func GetAllJobApplication()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.JobApplication

	//----------------------------------------------------------------------------
	// Request the ORM to find all JobApplication
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all JobApplication" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all JobApplication", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllJobApplication", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateJobApplication - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateJobApplication(obj model.JobApplication)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a JobApplication using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a JobApplication using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateJobApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteJobApplication - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteJobApplication(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetJobApplication(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.JobApplication)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a JobApplication using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a JobApplication using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteJobApplication", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Candidate on a JobApplication
//----------------------------------------------------------------------------
func AssignCandidateToJobApplication( jobApplicationId uint64, candidateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

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
			// assign the Candidate	to the JobApplication
			//----------------------------------------------------------------------------
			parentObj.Candidate = &childObj

			//----------------------------------------------------------------------------
			// save the JobApplication
			//----------------------------------------------------------------------------
			return UpdateJobApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidate", candidateId )
			return utils.RequestResult{false, msg, "assignCandidate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Candidate on a JobApplication
//----------------------------------------------------------------------------
func UnassignCandidateFromJobApplication(jobApplicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

		//----------------------------------------------------------------------------
		// assign an empty Candidate to the Candidate
		//----------------------------------------------------------------------------
		parentObj.Candidate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Candidate
		//----------------------------------------------------------------------------
		parentObj.CandidateId = nil;

		//----------------------------------------------------------------------------
		// save the JobApplication
		//----------------------------------------------------------------------------
		return UpdateJobApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Requisition on a JobApplication
//----------------------------------------------------------------------------
func AssignRequisitionToJobApplication( jobApplicationId uint64, requisitionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

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
			// assign the Requisition	to the JobApplication
			//----------------------------------------------------------------------------
			parentObj.Requisition = &childObj

			//----------------------------------------------------------------------------
			// save the JobApplication
			//----------------------------------------------------------------------------
			return UpdateJobApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requisition", requisitionId )
			return utils.RequestResult{false, msg, "assignRequisition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Requisition on a JobApplication
//----------------------------------------------------------------------------
func UnassignRequisitionFromJobApplication(jobApplicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

		//----------------------------------------------------------------------------
		// assign an empty JobRequisition to the Requisition
		//----------------------------------------------------------------------------
		parentObj.Requisition = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Requisition
		//----------------------------------------------------------------------------
		parentObj.RequisitionId = nil;

		//----------------------------------------------------------------------------
		// save the JobApplication
		//----------------------------------------------------------------------------
		return UpdateJobApplication(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more screeningsIds as a Screenings to a JobApplication
//----------------------------------------------------------------------------
func AddScreeningsToJobApplication ( jobApplicationId uint64, screeningsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Screening

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Screening
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Screenings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "unassignScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobApplication from the gorm
		//----------------------------------------------------------------------------
		return GetJobApplication(jobApplicationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more screeningsIds as a Screenings from a JobApplication
//----------------------------------------------------------------------------
func RemoveScreeningsFromJobApplication( jobApplicationId uint64, screeningsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobApplication(jobApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobApplication)

		// slice the ids on comma with no spaces
		ids := strings.Split( screeningsIds, ",")

		for _, screeningsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Screening

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Screening
			// with a matching screeningsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , screeningsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ScreeningObj from the Screenings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Screenings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Screenings", screeningsId )
				return utils.RequestResult{false, msg, "removeScreenings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobApplication from the gorm
		//----------------------------------------------------------------------------
		return GetJobApplication(jobApplicationId)

	} else {
		return parentRequestResult
	}
}

