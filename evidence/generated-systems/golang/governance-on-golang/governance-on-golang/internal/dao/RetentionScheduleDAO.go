package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RetentionScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRetentionSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateRetentionSchedule(obj model.RetentionSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RetentionSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RetentionSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRetentionSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRetentionSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRetentionSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RetentionSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RetentionSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RetentionSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RetentionSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRetentionSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRetentionSchedule - returns all
//----------------------------------------------------------------------------
func GetAllRetentionSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RetentionSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all RetentionSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RetentionSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RetentionSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRetentionSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRetentionSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRetentionSchedule(obj model.RetentionSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RetentionSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RetentionSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRetentionSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRetentionSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRetentionSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRetentionSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RetentionSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RetentionSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RetentionSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRetentionSchedule", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more repositoriesIds as a Repositories to a RetentionSchedule
//----------------------------------------------------------------------------
func AddRepositoriesToRetentionSchedule ( retentionScheduleId uint64, repositoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( repositoriesIds, ",")

		for _, repositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching repositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Repositories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Repositories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Repositories", repositoriesId )
				return utils.RequestResult{false, msg, "unassignRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more repositoriesIds as a Repositories from a RetentionSchedule
//----------------------------------------------------------------------------
func RemoveRepositoriesFromRetentionSchedule( retentionScheduleId uint64, repositoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( repositoriesIds, ",")

		for _, repositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching repositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RecordsRepositoryObj from the Repositories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Repositories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Repositories", repositoriesId )
				return utils.RequestResult{false, msg, "removeRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a RetentionSchedule
//----------------------------------------------------------------------------
func AddRecordsToRetentionSchedule ( retentionScheduleId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsIds, ",")

		for _, recordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Record_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Record_
			// with a matching recordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Records using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Records").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Records", recordsId )
				return utils.RequestResult{false, msg, "unassignRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a RetentionSchedule
//----------------------------------------------------------------------------
func RemoveRecordsFromRetentionSchedule( retentionScheduleId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsIds, ",")

		for _, recordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Record_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Record_
			// with a matching recordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Record_Obj from the Records array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Records").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Records", recordsId )
				return utils.RequestResult{false, msg, "removeRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more exceptionsIds as a Exceptions to a RetentionSchedule
//----------------------------------------------------------------------------
func AddExceptionsToRetentionSchedule ( retentionScheduleId uint64, exceptionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exception_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exception_
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Exceptions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "unassignExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more exceptionsIds as a Exceptions from a RetentionSchedule
//----------------------------------------------------------------------------
func RemoveExceptionsFromRetentionSchedule( retentionScheduleId uint64, exceptionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exception_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exception_
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Exception_Obj from the Exceptions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "removeExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dispositionReviewsIds as a DispositionReviews to a RetentionSchedule
//----------------------------------------------------------------------------
func AddDispositionReviewsToRetentionSchedule ( retentionScheduleId uint64, dispositionReviewsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( dispositionReviewsIds, ",")

		for _, dispositionReviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DispositionReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DispositionReview
			// with a matching dispositionReviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dispositionReviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DispositionReviews using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DispositionReviews").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DispositionReviews", dispositionReviewsId )
				return utils.RequestResult{false, msg, "unassignDispositionReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dispositionReviewsIds as a DispositionReviews from a RetentionSchedule
//----------------------------------------------------------------------------
func RemoveDispositionReviewsFromRetentionSchedule( retentionScheduleId uint64, dispositionReviewsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RetentionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRetentionSchedule(retentionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RetentionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RetentionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( dispositionReviewsIds, ",")

		for _, dispositionReviewsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DispositionReview

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DispositionReview
			// with a matching dispositionReviewsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dispositionReviewsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DispositionReviewObj from the DispositionReviews array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DispositionReviews").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DispositionReviews", dispositionReviewsId )
				return utils.RequestResult{false, msg, "removeDispositionReviews", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RetentionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRetentionSchedule(retentionScheduleId)

	} else {
		return parentRequestResult
	}
}

