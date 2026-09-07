package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing JobRequisitionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateJobRequisition - creates a new db entry
//----------------------------------------------------------------------------
func CreateJobRequisition(obj model.JobRequisition)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a JobRequisition with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a JobRequisition", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateJobRequisition", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetJobRequisition - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetJobRequisition(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.JobRequisition

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a JobRequisition with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a JobRequisition using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a JobRequisition using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetJobRequisition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllJobRequisition - returns all
//----------------------------------------------------------------------------
func GetAllJobRequisition()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.JobRequisition

	//----------------------------------------------------------------------------
	// Request the ORM to find all JobRequisition
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all JobRequisition" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all JobRequisition", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllJobRequisition", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateJobRequisition - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateJobRequisition(obj model.JobRequisition)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a JobRequisition using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a JobRequisition using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateJobRequisition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteJobRequisition - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteJobRequisition(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetJobRequisition(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a JobRequisition using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a JobRequisition using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteJobRequisition", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Department on a JobRequisition
//----------------------------------------------------------------------------
func AssignDepartmentToJobRequisition( jobRequisitionId uint64, departmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Department

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Department with a
		// matching departmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, departmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Department	to the JobRequisition
			//----------------------------------------------------------------------------
			parentObj.Department = &childObj

			//----------------------------------------------------------------------------
			// save the JobRequisition
			//----------------------------------------------------------------------------
			return UpdateJobRequisition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Department", departmentId )
			return utils.RequestResult{false, msg, "assignDepartment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Department on a JobRequisition
//----------------------------------------------------------------------------
func UnassignDepartmentFromJobRequisition(jobRequisitionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// assign an empty Department to the Department
		//----------------------------------------------------------------------------
		parentObj.Department = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Department
		//----------------------------------------------------------------------------
		parentObj.DepartmentId = nil;

		//----------------------------------------------------------------------------
		// save the JobRequisition
		//----------------------------------------------------------------------------
		return UpdateJobRequisition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a HiringManager on a JobRequisition
//----------------------------------------------------------------------------
func AssignHiringManagerToJobRequisition( jobRequisitionId uint64, hiringManagerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching hiringManagerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, hiringManagerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the HiringManager	to the JobRequisition
			//----------------------------------------------------------------------------
			parentObj.HiringManager = &childObj

			//----------------------------------------------------------------------------
			// save the JobRequisition
			//----------------------------------------------------------------------------
			return UpdateJobRequisition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "HiringManager", hiringManagerId )
			return utils.RequestResult{false, msg, "assignHiringManager", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a HiringManager on a JobRequisition
//----------------------------------------------------------------------------
func UnassignHiringManagerFromJobRequisition(jobRequisitionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the HiringManager
		//----------------------------------------------------------------------------
		parentObj.HiringManager = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the HiringManager
		//----------------------------------------------------------------------------
		parentObj.HiringManagerId = nil;

		//----------------------------------------------------------------------------
		// save the JobRequisition
		//----------------------------------------------------------------------------
		return UpdateJobRequisition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Recruiter on a JobRequisition
//----------------------------------------------------------------------------
func AssignRecruiterToJobRequisition( jobRequisitionId uint64, recruiterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching recruiterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, recruiterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Recruiter	to the JobRequisition
			//----------------------------------------------------------------------------
			parentObj.Recruiter = &childObj

			//----------------------------------------------------------------------------
			// save the JobRequisition
			//----------------------------------------------------------------------------
			return UpdateJobRequisition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Recruiter", recruiterId )
			return utils.RequestResult{false, msg, "assignRecruiter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Recruiter on a JobRequisition
//----------------------------------------------------------------------------
func UnassignRecruiterFromJobRequisition(jobRequisitionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Recruiter
		//----------------------------------------------------------------------------
		parentObj.Recruiter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Recruiter
		//----------------------------------------------------------------------------
		parentObj.RecruiterId = nil;

		//----------------------------------------------------------------------------
		// save the JobRequisition
		//----------------------------------------------------------------------------
		return UpdateJobRequisition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a JobProfile on a JobRequisition
//----------------------------------------------------------------------------
func AssignJobProfileToJobRequisition( jobRequisitionId uint64, jobProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobProfile with a
		// matching jobProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, jobProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the JobProfile	to the JobRequisition
			//----------------------------------------------------------------------------
			parentObj.JobProfile = &childObj

			//----------------------------------------------------------------------------
			// save the JobRequisition
			//----------------------------------------------------------------------------
			return UpdateJobRequisition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfile", jobProfileId )
			return utils.RequestResult{false, msg, "assignJobProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a JobProfile on a JobRequisition
//----------------------------------------------------------------------------
func UnassignJobProfileFromJobRequisition(jobRequisitionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		//----------------------------------------------------------------------------
		// assign an empty JobProfile to the JobProfile
		//----------------------------------------------------------------------------
		parentObj.JobProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the JobProfile
		//----------------------------------------------------------------------------
		parentObj.JobProfileId = nil;

		//----------------------------------------------------------------------------
		// save the JobRequisition
		//----------------------------------------------------------------------------
		return UpdateJobRequisition(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more candidatesIds as a Candidates to a JobRequisition
//----------------------------------------------------------------------------
func AddCandidatesToJobRequisition ( jobRequisitionId uint64, candidatesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( candidatesIds, ",")

		for _, candidatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Candidate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Candidate
			// with a matching candidatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , candidatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Candidates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Candidates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidates", candidatesId )
				return utils.RequestResult{false, msg, "unassignCandidates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more candidatesIds as a Candidates from a JobRequisition
//----------------------------------------------------------------------------
func RemoveCandidatesFromJobRequisition( jobRequisitionId uint64, candidatesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( candidatesIds, ",")

		for _, candidatesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Candidate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Candidate
			// with a matching candidatesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , candidatesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CandidateObj from the Candidates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Candidates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidates", candidatesId )
				return utils.RequestResult{false, msg, "removeCandidates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more interviewsIds as a Interviews to a JobRequisition
//----------------------------------------------------------------------------
func AddInterviewsToJobRequisition ( jobRequisitionId uint64, interviewsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewsIds, ",")

		for _, interviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Interview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Interview
			// with a matching interviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Interviews using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviews").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviews", interviewsId )
				return utils.RequestResult{false, msg, "unassignInterviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more interviewsIds as a Interviews from a JobRequisition
//----------------------------------------------------------------------------
func RemoveInterviewsFromJobRequisition( jobRequisitionId uint64, interviewsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( interviewsIds, ",")

		for _, interviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Interview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Interview
			// with a matching interviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , interviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InterviewObj from the Interviews array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Interviews").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Interviews", interviewsId )
				return utils.RequestResult{false, msg, "removeInterviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more offersIds as a Offers to a JobRequisition
//----------------------------------------------------------------------------
func AddOffersToJobRequisition ( jobRequisitionId uint64, offersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( offersIds, ",")

		for _, offersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Offer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Offer
			// with a matching offersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , offersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Offers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Offers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Offers", offersId )
				return utils.RequestResult{false, msg, "unassignOffers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more offersIds as a Offers from a JobRequisition
//----------------------------------------------------------------------------
func RemoveOffersFromJobRequisition( jobRequisitionId uint64, offersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobRequisition with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobRequisition(jobRequisitionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobRequisition so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobRequisition)

		// slice the ids on comma with no spaces
		ids := strings.Split( offersIds, ",")

		for _, offersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Offer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Offer
			// with a matching offersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , offersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OfferObj from the Offers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Offers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Offers", offersId )
				return utils.RequestResult{false, msg, "removeOffers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobRequisition from the gorm
		//----------------------------------------------------------------------------
		return GetJobRequisition(jobRequisitionId)

	} else {
		return parentRequestResult
	}
}

