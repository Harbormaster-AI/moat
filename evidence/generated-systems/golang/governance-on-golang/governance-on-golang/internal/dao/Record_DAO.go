package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing Record_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRecord_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateRecord_(obj model.Record_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Record_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Record_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRecord_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRecord_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRecord_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Record_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Record_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Record_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Record_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRecord_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRecord_ - returns all
//----------------------------------------------------------------------------
func GetAllRecord_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Record_

	//----------------------------------------------------------------------------
	// Request the ORM to find all Record_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Record_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Record_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRecord_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRecord_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRecord_(obj model.Record_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Record_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Record_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRecord_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRecord_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRecord_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRecord_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Record_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Record_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Record_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRecord_", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Repository on a Record_
//----------------------------------------------------------------------------
func AssignRepositoryToRecord_( record_Id uint64, repositoryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RecordsRepository

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RecordsRepository with a
		// matching repositoryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, repositoryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Repository	to the Record_
			//----------------------------------------------------------------------------
			parentObj.Repository = &childObj

			//----------------------------------------------------------------------------
			// save the Record_
			//----------------------------------------------------------------------------
			return UpdateRecord_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Repository", repositoryId )
			return utils.RequestResult{false, msg, "assignRepository", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Repository on a Record_
//----------------------------------------------------------------------------
func UnassignRepositoryFromRecord_(record_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		//----------------------------------------------------------------------------
		// assign an empty RecordsRepository to the Repository
		//----------------------------------------------------------------------------
		parentObj.Repository = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Repository
		//----------------------------------------------------------------------------
		parentObj.RepositoryId = nil;

		//----------------------------------------------------------------------------
		// save the Record_
		//----------------------------------------------------------------------------
		return UpdateRecord_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RetentionSchedule on a Record_
//----------------------------------------------------------------------------
func AssignRetentionScheduleToRecord_( record_Id uint64, retentionScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RetentionSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RetentionSchedule with a
		// matching retentionScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, retentionScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RetentionSchedule	to the Record_
			//----------------------------------------------------------------------------
			parentObj.RetentionSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the Record_
			//----------------------------------------------------------------------------
			return UpdateRecord_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RetentionSchedule", retentionScheduleId )
			return utils.RequestResult{false, msg, "assignRetentionSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RetentionSchedule on a Record_
//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromRecord_(record_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		//----------------------------------------------------------------------------
		// assign an empty RetentionSchedule to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the Record_
		//----------------------------------------------------------------------------
		return UpdateRecord_(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a Record_
//----------------------------------------------------------------------------
func AddProcessingActivitiesToRecord_ ( record_Id uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( processingActivitiesIds, ",")

		for _, processingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching processingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , processingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProcessingActivities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcessingActivities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcessingActivities", processingActivitiesId )
				return utils.RequestResult{false, msg, "unassignProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a Record_
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromRecord_( record_Id uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( processingActivitiesIds, ",")

		for _, processingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching processingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , processingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataProcessingActivityObj from the ProcessingActivities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcessingActivities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcessingActivities", processingActivitiesId )
				return utils.RequestResult{false, msg, "removeProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataCategoriesIds as a DataCategories to a Record_
//----------------------------------------------------------------------------
func AddDataCategoriesToRecord_ ( record_Id uint64, dataCategoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataCategoriesIds, ",")

		for _, dataCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataCategory
			// with a matching dataCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataCategories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataCategories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataCategories", dataCategoriesId )
				return utils.RequestResult{false, msg, "unassignDataCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataCategoriesIds as a DataCategories from a Record_
//----------------------------------------------------------------------------
func RemoveDataCategoriesFromRecord_( record_Id uint64, dataCategoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataCategoriesIds, ",")

		for _, dataCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataCategory
			// with a matching dataCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataCategoryObj from the DataCategories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataCategories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataCategories", dataCategoriesId )
				return utils.RequestResult{false, msg, "removeDataCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more legalHoldsIds as a LegalHolds to a Record_
//----------------------------------------------------------------------------
func AddLegalHoldsToRecord_ ( record_Id uint64, legalHoldsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( legalHoldsIds, ",")

		for _, legalHoldsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LegalHold

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LegalHold
			// with a matching legalHoldsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , legalHoldsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LegalHolds using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LegalHolds").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LegalHolds", legalHoldsId )
				return utils.RequestResult{false, msg, "unassignLegalHolds", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more legalHoldsIds as a LegalHolds from a Record_
//----------------------------------------------------------------------------
func RemoveLegalHoldsFromRecord_( record_Id uint64, legalHoldsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( legalHoldsIds, ",")

		for _, legalHoldsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LegalHold

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LegalHold
			// with a matching legalHoldsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , legalHoldsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LegalHoldObj from the LegalHolds array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LegalHolds").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LegalHolds", legalHoldsId )
				return utils.RequestResult{false, msg, "removeLegalHolds", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataSubjectRequestsIds as a DataSubjectRequests to a Record_
//----------------------------------------------------------------------------
func AddDataSubjectRequestsToRecord_ ( record_Id uint64, dataSubjectRequestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSubjectRequestsIds, ",")

		for _, dataSubjectRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSubjectRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSubjectRequest
			// with a matching dataSubjectRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSubjectRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataSubjectRequests using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSubjectRequests").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSubjectRequests", dataSubjectRequestsId )
				return utils.RequestResult{false, msg, "unassignDataSubjectRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataSubjectRequestsIds as a DataSubjectRequests from a Record_
//----------------------------------------------------------------------------
func RemoveDataSubjectRequestsFromRecord_( record_Id uint64, dataSubjectRequestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Record_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecord_(record_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Record_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Record_)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSubjectRequestsIds, ",")

		for _, dataSubjectRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSubjectRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSubjectRequest
			// with a matching dataSubjectRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSubjectRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSubjectRequestObj from the DataSubjectRequests array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSubjectRequests").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSubjectRequests", dataSubjectRequestsId )
				return utils.RequestResult{false, msg, "removeDataSubjectRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Record_ from the gorm
		//----------------------------------------------------------------------------
		return GetRecord_(record_Id)

	} else {
		return parentRequestResult
	}
}

