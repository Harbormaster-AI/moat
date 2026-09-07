package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing JobFamilyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateJobFamily - creates a new db entry
//----------------------------------------------------------------------------
func CreateJobFamily(obj model.JobFamily)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a JobFamily with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a JobFamily", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateJobFamily", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetJobFamily - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetJobFamily(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.JobFamily

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a JobFamily with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a JobFamily using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a JobFamily using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetJobFamily", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllJobFamily - returns all
//----------------------------------------------------------------------------
func GetAllJobFamily()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.JobFamily

	//----------------------------------------------------------------------------
	// Request the ORM to find all JobFamily
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all JobFamily" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all JobFamily", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllJobFamily", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateJobFamily - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateJobFamily(obj model.JobFamily)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a JobFamily using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a JobFamily using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateJobFamily", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteJobFamily - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteJobFamily(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the JobFamily with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetJobFamily(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.JobFamily)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a JobFamily using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a JobFamily using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteJobFamily", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a JobFamily
//----------------------------------------------------------------------------
func AssignOrganizationToJobFamily( jobFamilyId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobFamily(jobFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobFamily)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the JobFamily
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the JobFamily
			//----------------------------------------------------------------------------
			return UpdateJobFamily(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a JobFamily
//----------------------------------------------------------------------------
func UnassignOrganizationFromJobFamily(jobFamilyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobFamily(jobFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobFamily)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the JobFamily
		//----------------------------------------------------------------------------
		return UpdateJobFamily(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more jobProfilesIds as a JobProfiles to a JobFamily
//----------------------------------------------------------------------------
func AddJobProfilesToJobFamily ( jobFamilyId uint64, jobProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobFamily(jobFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobFamily)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobProfilesIds, ",")

		for _, jobProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobProfile
			// with a matching jobProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the JobProfiles using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobProfiles").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfiles", jobProfilesId )
				return utils.RequestResult{false, msg, "unassignJobProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobFamily from the gorm
		//----------------------------------------------------------------------------
		return GetJobFamily(jobFamilyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more jobProfilesIds as a JobProfiles from a JobFamily
//----------------------------------------------------------------------------
func RemoveJobProfilesFromJobFamily( jobFamilyId uint64, jobProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobFamily with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobFamily(jobFamilyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobFamily so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobFamily)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobProfilesIds, ",")

		for _, jobProfilesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobProfile

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobProfile
			// with a matching jobProfilesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobProfilesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove JobProfileObj from the JobProfiles array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobProfiles").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobProfiles", jobProfilesId )
				return utils.RequestResult{false, msg, "removeJobProfiles", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobFamily from the gorm
		//----------------------------------------------------------------------------
		return GetJobFamily(jobFamilyId)

	} else {
		return parentRequestResult
	}
}

