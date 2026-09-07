package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TrainingCourseDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTrainingCourse - creates a new db entry
//----------------------------------------------------------------------------
func CreateTrainingCourse(obj model.TrainingCourse)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TrainingCourse with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TrainingCourse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTrainingCourse", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTrainingCourse - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTrainingCourse(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TrainingCourse

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TrainingCourse with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TrainingCourse using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TrainingCourse using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTrainingCourse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTrainingCourse - returns all
//----------------------------------------------------------------------------
func GetAllTrainingCourse()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TrainingCourse

	//----------------------------------------------------------------------------
	// Request the ORM to find all TrainingCourse
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TrainingCourse" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TrainingCourse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTrainingCourse", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTrainingCourse - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTrainingCourse(obj model.TrainingCourse)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TrainingCourse using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TrainingCourse using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTrainingCourse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTrainingCourse - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTrainingCourse(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTrainingCourse(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TrainingCourse)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TrainingCourse using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TrainingCourse using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTrainingCourse", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more prerequisitesIds as a Prerequisites to a TrainingCourse
//----------------------------------------------------------------------------
func AddPrerequisitesToTrainingCourse ( trainingCourseId uint64, prerequisitesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

		// slice the ids on comma with no spaces
		ids := strings.Split( prerequisitesIds, ",")

		for _, prerequisitesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingCourse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingCourse
			// with a matching prerequisitesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , prerequisitesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Prerequisites using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Prerequisites").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Prerequisites", prerequisitesId )
				return utils.RequestResult{false, msg, "unassignPrerequisites", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more prerequisitesIds as a Prerequisites from a TrainingCourse
//----------------------------------------------------------------------------
func RemovePrerequisitesFromTrainingCourse( trainingCourseId uint64, prerequisitesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

		// slice the ids on comma with no spaces
		ids := strings.Split( prerequisitesIds, ",")

		for _, prerequisitesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingCourse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingCourse
			// with a matching prerequisitesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , prerequisitesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrainingCourseObj from the Prerequisites array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Prerequisites").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Prerequisites", prerequisitesId )
				return utils.RequestResult{false, msg, "removePrerequisites", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more enrollmentsIds as a Enrollments to a TrainingCourse
//----------------------------------------------------------------------------
func AddEnrollmentsToTrainingCourse ( trainingCourseId uint64, enrollmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

		// slice the ids on comma with no spaces
		ids := strings.Split( enrollmentsIds, ",")

		for _, enrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingEnrollment
			// with a matching enrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Enrollments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enrollments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enrollments", enrollmentsId )
				return utils.RequestResult{false, msg, "unassignEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more enrollmentsIds as a Enrollments from a TrainingCourse
//----------------------------------------------------------------------------
func RemoveEnrollmentsFromTrainingCourse( trainingCourseId uint64, enrollmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

		// slice the ids on comma with no spaces
		ids := strings.Split( enrollmentsIds, ",")

		for _, enrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TrainingEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TrainingEnrollment
			// with a matching enrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TrainingEnrollmentObj from the Enrollments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enrollments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enrollments", enrollmentsId )
				return utils.RequestResult{false, msg, "removeEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more jobProfilesIds as a JobProfiles to a TrainingCourse
//----------------------------------------------------------------------------
func AddJobProfilesToTrainingCourse ( trainingCourseId uint64, jobProfilesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

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
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more jobProfilesIds as a JobProfiles from a TrainingCourse
//----------------------------------------------------------------------------
func RemoveJobProfilesFromTrainingCourse( trainingCourseId uint64, jobProfilesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingCourse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingCourse(trainingCourseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingCourse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingCourse)

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
		// retrieve the modified TrainingCourse from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingCourse(trainingCourseId)

	} else {
		return parentRequestResult
	}
}

