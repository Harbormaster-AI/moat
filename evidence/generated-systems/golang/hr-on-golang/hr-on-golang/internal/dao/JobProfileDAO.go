package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing JobProfileDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateJobProfile - creates a new db entry
//----------------------------------------------------------------------------
func CreateJobProfile(obj model.JobProfile)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a JobProfile with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a JobProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateJobProfile", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetJobProfile - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetJobProfile(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.JobProfile

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a JobProfile with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a JobProfile using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a JobProfile using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetJobProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllJobProfile - returns all
//----------------------------------------------------------------------------
func GetAllJobProfile()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.JobProfile

	//----------------------------------------------------------------------------
	// Request the ORM to find all JobProfile
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all JobProfile" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all JobProfile", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllJobProfile", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateJobProfile - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateJobProfile(obj model.JobProfile)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a JobProfile using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a JobProfile using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateJobProfile", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteJobProfile - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteJobProfile(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetJobProfile(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.JobProfile)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a JobProfile using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a JobProfile using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteJobProfile", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a JobFamily on a JobProfile
//----------------------------------------------------------------------------
func AssignJobFamilyToJobProfile( jobProfileId uint64, jobFamilyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobFamily

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobFamily with a
		// matching jobFamilyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, jobFamilyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the JobFamily	to the JobProfile
			//----------------------------------------------------------------------------
			parentObj.JobFamily = &childObj

			//----------------------------------------------------------------------------
			// save the JobProfile
			//----------------------------------------------------------------------------
			return UpdateJobProfile(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobFamily", jobFamilyId )
			return utils.RequestResult{false, msg, "assignJobFamily", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a JobFamily on a JobProfile
//----------------------------------------------------------------------------
func UnassignJobFamilyFromJobProfile(jobProfileId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		//----------------------------------------------------------------------------
		// assign an empty JobFamily to the JobFamily
		//----------------------------------------------------------------------------
		parentObj.JobFamily = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the JobFamily
		//----------------------------------------------------------------------------
		parentObj.JobFamilyId = nil;

		//----------------------------------------------------------------------------
		// save the JobProfile
		//----------------------------------------------------------------------------
		return UpdateJobProfile(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more competenciesIds as a Competencies to a JobProfile
//----------------------------------------------------------------------------
func AddCompetenciesToJobProfile ( jobProfileId uint64, competenciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( competenciesIds, ",")

		for _, competenciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Competency

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Competency
			// with a matching competenciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competenciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Competencies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Competencies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Competencies", competenciesId )
				return utils.RequestResult{false, msg, "unassignCompetencies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more competenciesIds as a Competencies from a JobProfile
//----------------------------------------------------------------------------
func RemoveCompetenciesFromJobProfile( jobProfileId uint64, competenciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( competenciesIds, ",")

		for _, competenciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Competency

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Competency
			// with a matching competenciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , competenciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CompetencyObj from the Competencies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Competencies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Competencies", competenciesId )
				return utils.RequestResult{false, msg, "removeCompetencies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more trainingRecommendationsIds as a TrainingRecommendations to a JobProfile
//----------------------------------------------------------------------------
func AddTrainingRecommendationsToJobProfile ( jobProfileId uint64, trainingRecommendationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( trainingRecommendationsIds, ",")

		for _, trainingRecommendationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingCourse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingCourse
			// with a matching trainingRecommendationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingRecommendationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TrainingRecommendations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingRecommendations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRecommendations", trainingRecommendationsId )
				return utils.RequestResult{false, msg, "unassignTrainingRecommendations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more trainingRecommendationsIds as a TrainingRecommendations from a JobProfile
//----------------------------------------------------------------------------
func RemoveTrainingRecommendationsFromJobProfile( jobProfileId uint64, trainingRecommendationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( trainingRecommendationsIds, ",")

		for _, trainingRecommendationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingCourse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingCourse
			// with a matching trainingRecommendationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , trainingRecommendationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrainingCourseObj from the TrainingRecommendations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TrainingRecommendations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRecommendations", trainingRecommendationsId )
				return utils.RequestResult{false, msg, "removeTrainingRecommendations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more positionsIds as a Positions to a JobProfile
//----------------------------------------------------------------------------
func AddPositionsToJobProfile ( jobProfileId uint64, positionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( positionsIds, ",")

		for _, positionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching positionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , positionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Positions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Positions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Positions", positionsId )
				return utils.RequestResult{false, msg, "unassignPositions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more positionsIds as a Positions from a JobProfile
//----------------------------------------------------------------------------
func RemovePositionsFromJobProfile( jobProfileId uint64, positionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the JobProfile with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetJobProfile(jobProfileId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.JobProfile so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.JobProfile)

		// slice the ids on comma with no spaces
		ids := strings.Split( positionsIds, ",")

		for _, positionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching positionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , positionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PositionObj from the Positions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Positions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Positions", positionsId )
				return utils.RequestResult{false, msg, "removePositions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified JobProfile from the gorm
		//----------------------------------------------------------------------------
		return GetJobProfile(jobProfileId)

	} else {
		return parentRequestResult
	}
}

